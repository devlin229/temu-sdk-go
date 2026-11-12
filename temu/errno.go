package temu

import "fmt"

// BusinessError 表示接口返回 success=false，保留平台错误码与错误信息。
type BusinessError struct {
	APIType string
	Code    int64
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("Temu %s 业务失败（%d）: %s", e.APIType, e.Code, e.Message)
}
