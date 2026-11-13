// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goods.add.property/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsAddPropertyRequest 是 bg.glo.goods.add.property 的业务请求。
type BgGloGoodsAddPropertyRequest struct {
	ProductProperties []Property `json:"productProperties"` // 必填；Product Attributes
	ProductID         int64      `json:"productId"`         // 必填；Product ID
}

// BgGloGoodsAddPropertyRespose 是 bg.glo.goods.add.property 的完整响应。
type BgGloGoodsAddPropertyRespose struct {
	common.Response
	Result *PropertyResult `json:"result"` // result
}
