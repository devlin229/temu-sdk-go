package temu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const gatewayPA = "https://openapi-b-partner.temu.com/openapi/router"
const gatewayCN = "https://openapi.kuajingmaihuo.com/openapi/router"

type Temu struct {
	AppKey      string
	AppSecret   string
	AccessToken string
	httpClient  *http.Client
}

var (
	instance *Temu
	once     sync.Once
)

// NewClient 返回进程内唯一客户端，只有第一次调用传入的凭据生效。
// 初始化后不要修改凭据字段；切换账号需要重启进程。
func NewClient(appKey, appSecret, accessToken string) *Temu {
	once.Do(func() {
		instance = &Temu{
			AppKey:      appKey,
			AppSecret:   appSecret,
			AccessToken: accessToken,
			httpClient:  &http.Client{Timeout: 30 * time.Second},
		}
	})
	return instance
}

// requestInto 复用签名与 HTTP 流程，直接解码到各接口的响应类型。
func (t *Temu) requestInto(ctx context.Context, gateway, apiType string, params, result any) error {
	raw, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("encode business parameters: %w", err)
	}
	var body map[string]json.RawMessage
	if err = json.Unmarshal(raw, &body); err != nil {
		return fmt.Errorf("business parameters must be a JSON object: %w", err)
	}
	if body == nil {
		body = make(map[string]json.RawMessage)
	}
	// 只操作新解码的对象，不修改调用方的参数。
	for key, value := range map[string]string{
		"app_key": t.AppKey, "access_token": t.AccessToken,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		"type":      apiType, "data_type": "JSON",
	} {
		body[key], _ = json.Marshal(value)
	}
	delete(body, "app_secret")
	delete(body, "sign")
	signature, err := sign(body, t.AppSecret)
	if err != nil {
		return err
	}
	body["sign"], _ = json.Marshal(signature)
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("encode request parameters: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, gateway, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("call %s: %w", apiType, err)
	}
	defer resp.Body.Close()
	// 先保留原始正文，HTTP 错误或模型解析失败时调用方仍能获取平台实际返回。
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &ResponseError{
			APIType: apiType, StatusCode: resp.StatusCode, Body: responseBody,
			Err: fmt.Errorf("read response body: %w", err),
		}
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &ResponseError{APIType: apiType, StatusCode: resp.StatusCode, Body: responseBody}
	}
	decoder := json.NewDecoder(bytes.NewReader(responseBody))
	decoder.UseNumber()
	if err = decoder.Decode(result); err != nil {
		return &ResponseError{
			APIType: apiType, StatusCode: resp.StatusCode, Body: responseBody,
			Err: fmt.Errorf("decode response: %w", err),
		}
	}
	return nil
}
