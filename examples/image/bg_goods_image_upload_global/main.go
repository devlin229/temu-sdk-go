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

	"github.com/devlin229/temu-sdk-go/examples/config"
	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/image"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// 业务字段结构参考同目录 request.example.json；占位值需替换，不能直接发起真实业务。
func run() error {
	// 凭据统一加载，图片路径在下方填写。
	credentials, err := config.LoadCredentials()
	if err != nil {
		return err
	}
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

	client := temu.NewClient(credentials.AppKey, credentials.AppSecret, credentials.AccessToken)
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
