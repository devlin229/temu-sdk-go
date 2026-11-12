package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/category"
)

// BgGloGoodsSpecCreate 调用 bg.glo.goods.spec.create，创建商品发布使用的自定义规格。
// 暂沿用当前 PA 网关，截图未提供网关与权限信息，真实接入时需核验。
// 不自动重试；业务失败返回完整响应及 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsSpecCreate(ctx context.Context, request *category.BgGloGoodsSpecCreateRequest) (*category.BgGloGoodsSpecCreateRespose, error) {
	if request == nil {
		return nil, errors.New("创建规格请求不能为空")
	}
	const apiType = "bg.glo.goods.spec.create"
	var response category.BgGloGoodsSpecCreateRespose
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

// BgGloGoodsAttrsGet 调用 bg.glo.goods.attrs.get，查询叶子类目的属性模板。
// 业务参数及条件由平台验证，不自动选择属性或重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGloGoodsAttrsGet(ctx context.Context, request *category.BgGloGoodsAttrsGetRequest) (*category.BgGloGoodsAttrsGetRespose, error) {
	if request == nil {
		return nil, errors.New("类目属性查询请求不能为空")
	}
	const apiType = "bg.glo.goods.attrs.get"
	var response category.BgGloGoodsAttrsGetRespose
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

// BgGloGoodsCategoryMatch 调用 bg.glo.goods.category.match，按名称搜索类目路径。
// 两个业务参数均可选；请求通过当前 PA 网关发送，不自动选择类目或提交商品。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGloGoodsCategoryMatch(ctx context.Context, request *category.BgGloGoodsCategoryMatchRequest) (*category.BgGloGoodsCategoryMatchRespose, error) {
	if request == nil {
		return nil, errors.New("类目搜索请求不能为空")
	}
	const apiType = "bg.glo.goods.category.match"
	var response category.BgGloGoodsCategoryMatchRespose
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

// BgGloGoodsParentspecGet 调用 bg.glo.goods.parentspec.get，查询父规格列表。
// 无业务请求参数，客户端只添加公共授权、时间戳和签名参数，不自动重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGloGoodsParentspecGet(ctx context.Context) (*category.BgGloGoodsParentspecGetRespose, error) {
	const apiType = "bg.glo.goods.parentspec.get"
	var response category.BgGloGoodsParentspecGetRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, struct{}{}, &response); err != nil {
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
