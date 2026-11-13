package main

import (
	"bytes"
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
	// 本地测试直接填写，凭据保持为空；运行前核对接口是否为写操作。
	appKey := ""
	appSecret := ""
	accessToken := ""
	if appKey == "" || appSecret == "" || accessToken == "" {
		return errors.New("请先填写 appKey、appSecret、accessToken")
	}
	// 填写真实业务 JSON，字段见 docs/Temu_PA_APIs/category/bg.glo.oem.page.list/parameters.json。
	requestJSON := ""
	if requestJSON == "" {
		return errors.New("请先填写 requestJSON 中的真实业务参数")
	}
	var request category.BgGloOemPageListRequest
	decoder := json.NewDecoder(bytes.NewBufferString(requestJSON))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return fmt.Errorf("解析业务参数: %w", err)
	}
	client := temu.NewClient(appKey, appSecret, accessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, callErr := client.BgGloOemPageList(ctx, &request)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(callErr, fmt.Errorf("编码响应: %w", err))
		}
		if err := os.WriteFile("examples/category/bg_glo_oem_page_list/bg_glo_oem_page_list-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(callErr, fmt.Errorf("保存响应: %w", err))
		}
		if _, err := fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(callErr, fmt.Errorf("输出响应: %w", err))
		}
	}
	return callErr
}
