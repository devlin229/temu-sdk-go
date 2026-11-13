package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/jit"
)

// BgGloJitmodeActivate 调用 bg.glo.jitmode.activate，提交指定货品及 SKC 的 JIT 操作。
// 字段及权限见 docs/Temu_PA_APIs/jit/bg.glo.jitmode.activate/parameters.json。
// 文档未定义开关字段，不推断调用可用于关闭 JIT；不自动重试。
// 业务失败返回完整响应与 *BusinessError；传输或解析失败返回 nil 响应。
func (t *Temu) BgGloJitmodeActivate(ctx context.Context, request *jit.BgGloJitmodeActivateRequest) (*jit.BgGloJitmodeActivateRespose, error) {
	if request == nil {
		return nil, errors.New("JIT 操作请求不能为空")
	}
	const apiType = "bg.glo.jitmode.activate"
	var response jit.BgGloJitmodeActivateRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{
			APIType: apiType,
			Code:    response.ErrorCode,
			Message: response.ErrorMessage,
		}
	}
	return &response, nil
}
