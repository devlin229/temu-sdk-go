package temu

import "fmt"

// BusinessError 表示接口返回 success=false，保留平台错误码与错误信息。
type BusinessError struct {
	APIType string
	Code    int64
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("Temu %s business error (%d): %s", e.APIType, e.Code, e.Message)
}

// ResponseError 表示已收到 HTTP 响应，但状态异常、正文读取失败或无法解析为接口模型。
// Body 保留平台返回的原始字节（读取失败时为已收到的部分），不补造业务响应。
// 可通过 errors.As 获取本类型，或通过 errors.Is/As 检查底层读取和解析错误。
type ResponseError struct {
	APIType    string
	StatusCode int
	Body       []byte
	Err        error
}

func (e *ResponseError) Error() string {
	message := fmt.Sprintf("Temu %s returned HTTP %d", e.APIType, e.StatusCode)
	if e.Err != nil {
		message += ": " + e.Err.Error()
	}
	if len(e.Body) > 0 {
		message += "\n" + string(e.Body)
	}
	return message
}

// Unwrap 保留读取、解析失败的原始原因；单纯 HTTP 状态错误时返回 nil。
func (e *ResponseError) Unwrap() error {
	return e.Err
}
