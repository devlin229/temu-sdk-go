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
	"github.com/devlin229/temu-sdk-go/temu/category"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 将同目录 request.example.json 复制为 request.json，填写真实业务参数后运行。
func run() error {
	// 凭据统一从系统环境变量或 examples/.env 获取；运行前核对是否为写操作。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	var request category.BgGoodsRedressOptionalcategoryCorrectRequest
	if err := config.LoadRequest("examples/category/bg_goods_redress_optionalcategory_correct/request.json", &request); err != nil {
		return err
	}
	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, callErr := client.BgGoodsRedressOptionalcategoryCorrect(ctx, &request)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(callErr, fmt.Errorf("编码响应: %w", err))
		}
		if err := os.WriteFile("examples/category/bg_goods_redress_optionalcategory_correct/bg_goods_redress_optionalcategory_correct-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(callErr, fmt.Errorf("保存响应: %w", err))
		}
		if _, err := fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(callErr, fmt.Errorf("输出响应: %w", err))
		}
	}
	return callErr
}
