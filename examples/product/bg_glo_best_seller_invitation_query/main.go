package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/examples/config"
	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/product"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 业务字段结构参考同目录 request.example.json；占位值需替换，不能直接发起真实业务。
func run() error {
	// 凭据统一从系统环境变量或 examples/.env 获取；运行前核对是否为写操作。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	// 填写真实业务 JSON，字段见 docs/Temu_PA_APIs/product/bg.glo.best.seller.invitation.query/parameters.json。
	requestJSON := ""
	if requestJSON == "" {
		return errors.New("请先填写 requestJSON 中的真实业务参数")
	}
	var request product.BgGloBestSellerInvitationQueryRequest
	decoder := json.NewDecoder(bytes.NewBufferString(requestJSON))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return fmt.Errorf("解析业务参数: %w", err)
	}
	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, callErr := client.BgGloBestSellerInvitationQuery(ctx, &request)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(callErr, fmt.Errorf("编码响应: %w", err))
		}
		if err := os.WriteFile("examples/product/bg_glo_best_seller_invitation_query/bg_glo_best_seller_invitation_query-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(callErr, fmt.Errorf("保存响应: %w", err))
		}
		if _, err := fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(callErr, fmt.Errorf("输出响应: %w", err))
		}
	}
	return callErr
}
