package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/image"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// 本地测试时直接填写凭据和图片路径。
	appKey := "47bb4bb7769e12d9f7aa93cf029fe529"
	appSecret := "ac0a3e952eaaa5b19c0e615c2ef497f50afa6e49"
	accessToken := "ayzqiujmpkuyst6jpwschlfc6kgzaxhhd9aghi2hdujvia1grfkwvrsv"
	imagePath := "/Users/j/workspace/bjd/c4f/O1CN01YJs6gG1dXGOBB1Sln_!!4611686018427384145-0-item_pic.jpg_.webp"

	content, err := os.ReadFile(imagePath)
	if err != nil {
		return fmt.Errorf("读取图片: %w", err)
	}
	// 按接口示例使用 Data URL；根据文件内容识别类型，避免把 JPEG 标成 PNG。
	encoded := "data:" + http.DetectContentType(content) + ";base64," + base64.StdEncoding.EncodeToString(content)
	request := &image.BgGoodsImageUploadGlobalRequest{Image: encoded}
	// 外包装图片可设置 ImageBizType = bizType，其中 bizType := int64(1)。
	// 图片处理选项按需设置 request.Options = &image.Options{...}。

	client := temu.NewClient(appKey, appSecret, accessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, uploadErr := client.BgGoodsImageUploadGlobal(ctx, request)

	// 业务失败也保留平台的完整 JSON 响应；错误信息和退出状态单独处理。
	if response != nil {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err = encoder.Encode(response); err != nil {
			return errors.Join(uploadErr, fmt.Errorf("输出响应: %w", err))
		}
	}
	if uploadErr != nil {
		return fmt.Errorf("上传图片: %w", uploadErr)
	}
	return nil
}
