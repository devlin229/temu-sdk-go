package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/category"
)

// BgGloGoodsSpecCreate 调用 bg.glo.goods.spec.create，创建商品发布使用的自定义规格。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.spec.create/parameters.json。
// 使用当前 PA 字典指定的网关，真实接入时需核验卖家授权权限。
// 不自动重试；业务失败返回完整响应及 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsSpecCreate(ctx context.Context, request *category.BgGloGoodsSpecCreateRequest) (*category.BgGloGoodsSpecCreateRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.spec.create: request must not be nil")
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
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.attrs.get/parameters.json。
// 业务参数及条件由平台验证，不自动选择属性或重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGloGoodsAttrsGet(ctx context.Context, request *category.BgGloGoodsAttrsGetRequest) (*category.BgGloGoodsAttrsGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.attrs.get: request must not be nil")
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
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.category.match/parameters.json。
// 两个业务参数均可选；请求通过当前 PA 网关发送，不自动选择类目或提交商品。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGloGoodsCategoryMatch(ctx context.Context, request *category.BgGloGoodsCategoryMatchRequest) (*category.BgGloGoodsCategoryMatchRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.category.match: request must not be nil")
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
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.parentspec.get/parameters.json。
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

// BgGloGoodsAccessoriesGet 调用 bg.glo.goods.accessories.get，货品包装清单类型查询。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.accessories.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsAccessoriesGet(ctx context.Context, request *category.BgGloGoodsAccessoriesGetRequest) (*category.BgGloGoodsAccessoriesGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.accessories.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.accessories.get"
	var response category.BgGloGoodsAccessoriesGetRespose
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

// BgGloGoodsCatsGet 调用 bg.glo.goods.cats.get，按层级查询类目。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.cats.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsCatsGet(ctx context.Context, request *category.BgGloGoodsCatsGetRequest) (*category.BgGloGoodsCatsGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.cats.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.cats.get"
	var response category.BgGloGoodsCatsGetRespose
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

// BgGloGoodsCatsmandatoryGet 调用 bg.glo.goods.catsmandatory.get，类目必填信息查询接口。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.goods.catsmandatory.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsCatsmandatoryGet(ctx context.Context, request *category.BgGloGoodsCatsmandatoryGetRequest) (*category.BgGloGoodsCatsmandatoryGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.catsmandatory.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.catsmandatory.get"
	var response category.BgGloGoodsCatsmandatoryGetRespose
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

// BgGloOemPageList 调用 bg.glo.oem.page.list，提供检索 OE/OEM 数据。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.glo.oem.page.list/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloOemPageList(ctx context.Context, request *category.BgGloOemPageListRequest) (*category.BgGloOemPageListRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.oem.page.list: request must not be nil")
	}
	const apiType = "bg.glo.oem.page.list"
	var response category.BgGloOemPageListRespose
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

// BgGoodsAttributeMappingGlobal 调用 bg.goods.attribute.mapping.global，内外属性映射。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.goods.attribute.mapping.global/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGoodsAttributeMappingGlobal(ctx context.Context, request *category.BgGoodsAttributeMappingGlobalRequest) (*category.BgGoodsAttributeMappingGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.goods.attribute.mapping.global: request must not be nil")
	}
	const apiType = "bg.goods.attribute.mapping.global"
	var response category.BgGoodsAttributeMappingGlobalRespose
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

// BgGoodsRedressCorrectrecordQuery 调用 bg.goods.redress.correctrecord.query，提供给商家端，用于查询当前店铺的类目整改记录。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.goods.redress.correctrecord.query/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGoodsRedressCorrectrecordQuery(ctx context.Context, request *category.BgGoodsRedressCorrectrecordQueryRequest) (*category.BgGoodsRedressCorrectrecordQueryRespose, error) {
	if request == nil {
		return nil, errors.New("bg.goods.redress.correctrecord.query: request must not be nil")
	}
	const apiType = "bg.goods.redress.correctrecord.query"
	var response category.BgGoodsRedressCorrectrecordQueryRespose
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

// BgGoodsRedressOptionalcategoryCorrect 调用 bg.goods.redress.optionalcategory.correct，提供给商家端，用于主动选择商品的推荐类目。
// 字段及权限见 docs/Temu_PA_APIs/category/bg.goods.redress.optionalcategory.correct/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGoodsRedressOptionalcategoryCorrect(ctx context.Context, request *category.BgGoodsRedressOptionalcategoryCorrectRequest) (*category.BgGoodsRedressOptionalcategoryCorrectRespose, error) {
	if request == nil {
		return nil, errors.New("bg.goods.redress.optionalcategory.correct: request must not be nil")
	}
	const apiType = "bg.goods.redress.optionalcategory.correct"
	var response category.BgGoodsRedressOptionalcategoryCorrectRespose
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
