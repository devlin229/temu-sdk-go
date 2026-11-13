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
	// 凭据统一加载；运行后会真实创建规格，不自动重试。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	parentSpecID := int64(1001) // 填写真实父规格 ID。
	specName := "灰蓝上衣款"         // 填写要创建的子规格名称。

	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, queryErr := client.BgGloGoodsSpecCreate(ctx, &category.BgGloGoodsSpecCreateRequest{
		ParentSpecID: parentSpecID, SpecName: specName,
	})
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(queryErr, fmt.Errorf("编码创建规格响应: %w", err))
		}
		if err = os.WriteFile("examples/category/bg_glo_goods_spec_create/spec-create-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(queryErr, fmt.Errorf("保存创建规格响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(queryErr, fmt.Errorf("输出创建规格响应: %w", err))
		}
	}
	return queryErr
}
