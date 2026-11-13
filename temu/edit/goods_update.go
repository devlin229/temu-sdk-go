// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goods.update/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsUpdateRequest 是 bg.glo.goods.update 的业务请求。
type BgGloGoodsUpdateRequest struct {
	ProductWhExtAttrReq *WarehouseExtAttrRequest `json:"productWhExtAttrReq,omitempty"` // 可选；Commodity Warehouse and Supply Chain Extension Properties Request
	SupplierID          int64                    `json:"supplierId"`                    // 必填；Supplier ID
	ProductID           int64                    `json:"productId"`                     // 必填；Product ID
}

// BgGloGoodsUpdateRespose 是 bg.glo.goods.update 的完整响应。
type BgGloGoodsUpdateRespose struct {
	common.Response
	Result *Result `json:"result"` // result
}
