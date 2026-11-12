package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/product"
)

// BgGloGoodsAdd 调用 bg.glo.goods.add，创建并提交商品信息。
// 请求业务字段按原值发送，网关、公共参数、签名和请求日志由客户端处理，不自动重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
// 成功不表示商品已审核或可售；调用方应保存结果中的商品、SKC、SKU ID 与编码映射。
func (t *Temu) BgGloGoodsAdd(ctx context.Context, request *product.BgGloGoodsAddRequest) (*product.BgGloGoodsAddRespose, error) {
	if request == nil {
		return nil, errors.New("商品发布请求不能为空")
	}
	const apiType = "bg.glo.goods.add"
	var response product.BgGloGoodsAddRespose
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
