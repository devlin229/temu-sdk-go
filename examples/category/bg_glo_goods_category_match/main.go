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

// 业务字段结构参考同目录 request.example.json；占位值需替换，不能直接发起真实业务。
func run() error {
	// 凭据统一从系统环境变量或 examples/.env 获取。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	siteID := int64(0) // 可选；正数站点 ID，0 时不发送。

	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, queryErr := client.BgGloGoodsCategoryMatch(ctx, &category.BgGloGoodsCategoryMatchRequest{
		SearchText: "bjd",
		SiteID:     siteID,
	})
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(queryErr, fmt.Errorf("编码类目响应: %w", err))
		}
		if err = os.WriteFile("examples/category/bg_glo_goods_category_match/category-match-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(queryErr, fmt.Errorf("保存类目响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(queryErr, fmt.Errorf("输出类目响应: %w", err))
		}
	}
	return queryErr
}
