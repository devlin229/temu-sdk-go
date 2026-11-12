package temu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

// request 的 gateway 和 apiType 由各业务方法指定，调用方不需要区分 CN/PA。
func (t *Temu) request(ctx context.Context, gateway, apiType string, params any) (*Response, error) {
	raw, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("编码业务参数: %w", err)
	}
	var body map[string]json.RawMessage
	if err = json.Unmarshal(raw, &body); err != nil {
		return nil, fmt.Errorf("业务参数必须是对象: %w", err)
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
		return nil, err
	}
	body["sign"], _ = json.Marshal(signature)
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("编码请求参数: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, gateway, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("构造请求: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("调用 %s: %w", apiType, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("调用 %s 返回 HTTP %d", apiType, resp.StatusCode)
	}
	var result Response
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()
	if err = decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("解析 %s 响应: %w", apiType, err)
	}
	return &result, nil
}
