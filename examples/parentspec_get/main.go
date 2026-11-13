package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/temu"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// 本地测试直接填写字符串，无业务请求参数。
	appKey := "47bb4bb7769e12d9f7aa93cf029fe529"
	appSecret := "ac0a3e952eaaa5b19c0e615c2ef497f50afa6e49"
	accessToken := "ayzqiujmpkuyst6jpwschlfc6kgzaxhhd9aghi2hdujvia1grfkwvrsv"
	client := temu.NewClient(appKey, appSecret, accessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, queryErr := client.BgGloGoodsParentspecGet(ctx)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(queryErr, fmt.Errorf("编码父规格响应: %w", err))
		}
		if err = os.WriteFile("parentspec-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(queryErr, fmt.Errorf("保存父规格响应: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(queryErr, fmt.Errorf("输出父规格响应: %w", err))
		}
	}
	return queryErr
}
