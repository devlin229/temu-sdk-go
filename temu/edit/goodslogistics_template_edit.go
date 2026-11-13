// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goodslogistics.template.edit/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodslogisticsTemplateEditRequest 是 bg.glo.goodslogistics.template.edit 的业务请求。
type BgGloGoodslogisticsTemplateEditRequest struct {
	ProductShipment ShipmentRequest `json:"productShipment"` // 必填；Commodity Delivery Information
	SupplierID      int64           `json:"supplierId"`      // 必填；Supplier ID
	ProductID       int64           `json:"productId"`       // 必填；Product ID
}

// BgGloGoodslogisticsTemplateEditRespose 是 bg.glo.goodslogistics.template.edit 的完整响应。
type BgGloGoodslogisticsTemplateEditRespose struct {
	common.Response
	Result *Result `json:"result"` // result
}
