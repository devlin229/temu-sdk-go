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
	// 凭据统一加载；在下方填写真实叶子类目 ID。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
	catID := int64(25589)
	productCreateTime := int64(0) // 可选；毫秒时间戳，0 时省略。
	supportedType := int64(0)     // 可选；渠道枚举需核实，0 时省略。
	var langList []string         // 可选；一般不传，多语言查询影响性能。

	if catID <= 0 {
		return errors.New("请先在 examples/category/bg_glo_goods_attrs_get/main.go 中填写真实叶子类目 catID")
	}
	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, queryErr := client.BgGloGoodsAttrsGet(ctx, &category.BgGloGoodsAttrsGetRequest{
		CatID: catID, ProductCreateTime: productCreateTime,
		SupportedType: supportedType, LangList: langList,
	})
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(queryErr, fmt.Errorf("编码类目属性响应: %w", err))
		}
		if err = os.WriteFile("examples/category/bg_glo_goods_attrs_get/attrs-get-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(queryErr, fmt.Errorf("保存类目属性响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(queryErr, fmt.Errorf("输出类目属性响应: %w", err))
		}
	}
	return queryErr
}
