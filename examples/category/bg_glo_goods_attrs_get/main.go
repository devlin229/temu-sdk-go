package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/category"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// 本地测试直接填写字符串及真实叶子类目 ID，不读取环境变量或命令行参数。
	appKey := "47bb4bb7769e12d9f7aa93cf029fe529"
	appSecret := "ac0a3e952eaaa5b19c0e615c2ef497f50afa6e49"
	accessToken := "ayzqiujmpkuyst6jpwschlfc6kgzaxhhd9aghi2hdujvia1grfkwvrsv"
	catID := int64(25589)
	productCreateTime := int64(0) // 可选；毫秒时间戳，0 时省略。
	supportedType := int64(0)     // 可选；渠道枚举需核实，0 时省略。
	var langList []string         // 可选；一般不传，多语言查询影响性能。

	if catID <= 0 {
		return errors.New("请先在 examples/category/bg_glo_goods_attrs_get/main.go 中填写真实叶子类目 catID")
	}
	client := temu.NewClient(appKey, appSecret, accessToken)
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
