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
	// 直接填写字符串，不读取环境变量或命令行参数。
	appKey := "47bb4bb7769e12d9f7aa93cf029fe529"
	appSecret := "ac0a3e952eaaa5b19c0e615c2ef497f50afa6e49"
	accessToken := "ayzqiujmpkuyst6jpwschlfc6kgzaxhhd9aghi2hdujvia1grfkwvrsv"
	siteID := int64(0) // 可选，不按站点筛选时保持 0，SDK 会省略。

	client := temu.NewClient(appKey, appSecret, accessToken)
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
		if err = os.WriteFile("category-match-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(queryErr, fmt.Errorf("保存类目响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(queryErr, fmt.Errorf("输出类目响应: %w", err))
		}
	}
	return queryErr
}
