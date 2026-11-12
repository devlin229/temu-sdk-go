package temu

import "github.com/devlin229/temu-sdk-go/temu/common"

// Response 是公共客户端的通用响应，业务接口使用各自的 Result 类型。
type Response struct {
	common.Response
	Result any `json:"result"`
}
