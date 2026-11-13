package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/examples/config"
	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/image"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 将同目录 request.example.json 复制为 request.json，填写真实业务参数后运行。
func run() error {
	// 凭据统一从系统环境变量或 examples/.env 获取；运行将调用图片接口，不自动重试或轮询。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	var request image.BgCompliancepictureGetGlobalRequest
	if err = config.LoadRequest("examples/image/bg_compliancepicture_get_global/request.json", &request); err != nil {
		return err
	}
	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, callErr := client.BgCompliancepictureGetGlobal(ctx, &request)
	if response != nil {
		var raw []byte
		raw, err = json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(callErr, fmt.Errorf("编码响应: %w", err))
		}
		if err = os.WriteFile("examples/image/bg_compliancepicture_get_global/bg_compliancepicture_get_global-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(callErr, fmt.Errorf("保存响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(callErr, fmt.Errorf("输出响应: %w", err))
		}
	}
	return callErr
}
