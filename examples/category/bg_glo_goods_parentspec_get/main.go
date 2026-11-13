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
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// 凭据统一加载；无业务请求参数。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, queryErr := client.BgGloGoodsParentspecGet(ctx)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(queryErr, fmt.Errorf("编码父规格响应: %w", err))
		}
		if err = os.WriteFile("examples/category/bg_glo_goods_parentspec_get/parentspec-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(queryErr, fmt.Errorf("保存父规格响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(queryErr, fmt.Errorf("输出父规格响应: %w", err))
		}
	}
	return queryErr
}
