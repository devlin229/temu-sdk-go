package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/image"
)

// BgGoodsImageUploadGlobal 调用 bg.goods.image.upload.global，上传已编码为 Base64 的图片。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.goods.image.upload.global/parameters.json。
// 网关与公共签名由客户端处理，请求字段和图片处理选项按原值发送，不自动重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGoodsImageUploadGlobal(ctx context.Context, request *image.BgGoodsImageUploadGlobalRequest) (*image.BgGoodsImageUploadGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.goods.image.upload.global: request must not be nil")
	}
	const apiType = "bg.goods.image.upload.global"
	var response image.BgGoodsImageUploadGlobalRespose
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
