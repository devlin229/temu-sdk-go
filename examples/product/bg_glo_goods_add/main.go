package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/examples/config"
	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/product"
)

// requestExample 按用户反馈已成功上架的请求维护，作为缺文件时生成本地请求的唯一来源。
//
//go:embed request.example.json
var requestExample []byte

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 同目录 request.example.json 来自用户反馈成功上架的单款请求，复用前核对当前商品和店铺参数。
func run() error {
	// 凭据统一加载；运行此示例将真实调用商品发布接口。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}

	request, err := loadRequest("examples/product/bg_glo_goods_add/product-request.json")
	if err != nil {
		return err
	}
	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, publishErr := client.BgGloGoodsAdd(ctx, request)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(publishErr, fmt.Errorf("编码发品响应: %w", err))
		}
		// 即使业务失败也保存完整响应，成功时其中包含商品及 SKC/SKU 编码映射。
		if err = os.WriteFile("examples/product/bg_glo_goods_add/product-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(publishErr, fmt.Errorf("保存发品响应失败，请先核对平台记录: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(publishErr, fmt.Errorf("输出发品响应: %w", err))
		}
	}
	if publishErr != nil {
		return fmt.Errorf("测试发品失败，不自动重试: %w", publishErr)
	}
	return nil
}

// loadRequest 优先使用本地请求；缺文件时从嵌入的成功案例模板生成，不覆盖已有文件。
func loadRequest(path string) (*product.BgGloGoodsAddRequest, error) {
	raw, err := os.ReadFile(path)
	missing := errors.Is(err, os.ErrNotExist)
	if missing {
		raw = requestExample
	} else if err != nil {
		return nil, fmt.Errorf("读取商品请求: %w", err)
	}
	var request *product.BgGloGoodsAddRequest
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return nil, fmt.Errorf("解析商品请求: %w", err)
	}
	if request == nil {
		return nil, errors.New("商品请求必须是 JSON 对象，不能为 null")
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, errors.New("商品请求文件只能包含一个 JSON 对象")
	}
	if missing {
		// 使用独占创建，避免覆盖读取后恰好由用户创建的请求文件。
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
		if err != nil {
			return nil, fmt.Errorf("创建测试商品请求: %w", err)
		}
		_, writeErr := file.Write(raw)
		if err := errors.Join(writeErr, file.Close()); err != nil {
			return nil, fmt.Errorf("保存测试商品请求: %w", err)
		}
	}
	return request, nil
}
