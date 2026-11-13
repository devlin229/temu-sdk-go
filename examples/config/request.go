package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LoadRequest 从对应接口的本地请求文件读取业务参数，不自动创建或覆盖文件。
// 只接受一个 JSON 对象，严格检查未知字段；缺文件时提示复制同目录模板。
func LoadRequest(path string, request any) error {
	raw, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		template := filepath.Join(filepath.Dir(path), "request.example.json")
		return fmt.Errorf("读取请求文件 %s: %w；请先复制 %s 为 %s，并填写真实业务参数", path, err, template, path)
	}
	if err != nil {
		return fmt.Errorf("读取请求文件 %s: %w", path, err)
	}
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 || trimmed[0] != '{' {
		return fmt.Errorf("请求文件 %s 必须包含一个 JSON 对象，不能是空文件、null 或数组", path)
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(request); err != nil {
		return fmt.Errorf("解析请求文件 %s: %w", path, err)
	}
	var extra any
	if err = decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("请求文件 %s 只能包含一个 JSON 对象，末尾不能包含其他内容", path)
	}
	return nil
}
