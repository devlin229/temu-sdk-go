// 参数来源：docs/Temu_PA_APIs/product/bg.glo.goods.removed.get/parameters.json
// 页面更新时间：2026-04-19 18:18:33；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloGoodsRemovedGetRequest 是 bg.glo.goods.removed.get 的业务请求。
type BgGloGoodsRemovedGetRequest struct {
	ProductIDs []int64 `json:"productIds,omitzero"` // 可选；Product ID List
	PageSize   int64   `json:"pageSize,omitempty"`  // 可选；Page Size
	Page       int64   `json:"page,omitempty"`      // 可选；Page Number
}

// ProductSKUSummaries 对应接口中的嵌套对象。
type ProductSKUSummaries struct {
	ProductSKUID int64  `json:"productSkuId"` // Product SKU ID
	ExtCode      string `json:"extCode"`      // Sku Number
}

// GoodsRemovedData 对应接口中的嵌套对象。
type GoodsRemovedData struct {
	RemoveStatusUpdatedAtTs int64                 `json:"removeStatusUpdatedAtTs"` // Remove Status Updated At Timestamp ms
	ExtCode                 string                `json:"extCode"`                 // Product SKC External Code
	ProductID               int64                 `json:"productId"`               // Product Id
	ProductSKUSummaries     []ProductSKUSummaries `json:"productSkuSummaries"`     // SKU Summary Information List
	ProductSKCID            int64                 `json:"productSkcId"`            // Product SKC ID
	MainImageURL            string                `json:"mainImageUrl"`            // SKC Main Image
	ProductName             string                `json:"productName"`             // Product Name
}

// GoodsRemovedResult result
type GoodsRemovedResult struct {
	Data       []GoodsRemovedData `json:"data"`       // Result List
	TotalCount int64              `json:"totalCount"` // Total Count
}

// BgGloGoodsRemovedGetRespose 是 bg.glo.goods.removed.get 的完整响应。
type BgGloGoodsRemovedGetRespose struct {
	common.Response
	Result *GoodsRemovedResult `json:"result"` // result
}
