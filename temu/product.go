package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/product"
)

// BgGloGoodsAdd 调用 bg.glo.goods.add，创建并提交商品信息。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.add/parameters.json。
// 请求业务字段按原值发送，网关、公共参数、签名和请求日志由客户端处理，不自动重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
// 成功不表示商品已审核或可售；调用方应保存结果中的商品、SKC、SKU ID 与编码映射。
func (t *Temu) BgGloGoodsAdd(ctx context.Context, request *product.BgGloGoodsAddRequest) (*product.BgGloGoodsAddRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.add: request must not be nil")
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

// BgBtgGoodsStockWarehouseListGet 调用 bg.btg.goods.stock.warehouse.list.get，根据站点列表查询自发货模式品可绑定的发货仓信息。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.btg.goods.stock.warehouse.list.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgBtgGoodsStockWarehouseListGet(ctx context.Context, request *product.BgBtgGoodsStockWarehouseListGetRequest) (*product.BgBtgGoodsStockWarehouseListGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.btg.goods.stock.warehouse.list.get: request must not be nil")
	}
	const apiType = "bg.btg.goods.stock.warehouse.list.get"
	var response product.BgBtgGoodsStockWarehouseListGetRespose
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

// BgGloBestSellerInvitationQuery 调用 bg.glo.best.seller.invitation.query，best seller招标单查询。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.best.seller.invitation.query/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloBestSellerInvitationQuery(ctx context.Context, request *product.BgGloBestSellerInvitationQueryRequest) (*product.BgGloBestSellerInvitationQueryRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.best.seller.invitation.query: request must not be nil")
	}
	const apiType = "bg.glo.best.seller.invitation.query"
	var response product.BgGloBestSellerInvitationQueryRespose
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

// BgGloGoodsBrandGet 调用 bg.glo.goods.brand.get，用于货品发布前查询品牌。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.brand.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsBrandGet(ctx context.Context, request *product.BgGloGoodsBrandGetRequest) (*product.BgGloGoodsBrandGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.brand.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.brand.get"
	var response product.BgGloGoodsBrandGetRespose
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

// BgGloGoodsDetailGet 调用 bg.glo.goods.detail.get，查询商品详情信息。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.detail.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsDetailGet(ctx context.Context, request *product.BgGloGoodsDetailGetRequest) (*product.BgGloGoodsDetailGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.detail.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.detail.get"
	var response product.BgGloGoodsDetailGetRespose
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

// BgGloGoodsListGet 调用 bg.glo.goods.list.get，货品skc列表查询。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.list.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsListGet(ctx context.Context, request *product.BgGloGoodsListGetRequest) (*product.BgGloGoodsListGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.list.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.list.get"
	var response product.BgGloGoodsListGetRespose
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

// BgGloGoodsMigrate 调用 bg.glo.goods.migrate，半托管店铺搬运同主体下全托管店铺的货品。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.migrate/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsMigrate(ctx context.Context, request *product.BgGloGoodsMigrateRequest) (*product.BgGloGoodsMigrateRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.migrate: request must not be nil")
	}
	const apiType = "bg.glo.goods.migrate"
	var response product.BgGloGoodsMigrateRespose
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

// BgGloGoodsRemovedGet 调用 bg.glo.goods.removed.get，卖家中心已废弃列表。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.removed.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsRemovedGet(ctx context.Context, request *product.BgGloGoodsRemovedGetRequest) (*product.BgGloGoodsRemovedGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.goods.removed.get: request must not be nil")
	}
	const apiType = "bg.glo.goods.removed.get"
	var response product.BgGloGoodsRemovedGetRespose
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

// BgGloGoodsTopsellingSoldoutGet 调用 bg.glo.goods.topselling.soldout.get，批量查询爆款售罄商品。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.goods.topselling.soldout.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsTopsellingSoldoutGet(ctx context.Context) (*product.BgGloGoodsTopsellingSoldoutGetRespose, error) {
	const apiType = "bg.glo.goods.topselling.soldout.get"
	var response product.BgGloGoodsTopsellingSoldoutGetRespose
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

// BgGloLogisticsTemplateGet 调用 bg.glo.logistics.template.get，查询运费模版，用于半托管商品的API发布时关联运费模版ID。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.logistics.template.get/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloLogisticsTemplateGet(ctx context.Context, request *product.BgGloLogisticsTemplateGetRequest) (*product.BgGloLogisticsTemplateGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.logistics.template.get: request must not be nil")
	}
	const apiType = "bg.glo.logistics.template.get"
	var response product.BgGloLogisticsTemplateGetRespose
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

// BgGloProductSearch 调用 bg.glo.product.search，外部erp系统，查询货品生命周期状态。
// 字段及权限见 docs/Temu_PA_APIs/product/bg.glo.product.search/parameters.json。
// 不自动重试；业务失败返回完整响应与 *BusinessError，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloProductSearch(ctx context.Context, request *product.BgGloProductSearchRequest) (*product.BgGloProductSearchRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.product.search: request must not be nil")
	}
	const apiType = "bg.glo.product.search"
	var response product.BgGloProductSearchRespose
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
