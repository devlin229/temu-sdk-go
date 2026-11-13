// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goods.edit.sensitive.attr/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsEditSensitiveAttrRequest 是 bg.glo.goods.edit.sensitive.attr 的业务请求。
type BgGloGoodsEditSensitiveAttrRequest struct {
	ProductID  int64                 `json:"productId"`  // 必填；Product ID
	SKUReqList []SensitiveSKURequest `json:"skuReqList"` // 必填；SKU Sensitive Product Attribute Request List
}

// BgGloGoodsEditSensitiveAttrRespose 是 bg.glo.goods.edit.sensitive.attr 的完整响应。
type BgGloGoodsEditSensitiveAttrRespose struct {
	common.Response
	Result *Result `json:"result"` // result
}
