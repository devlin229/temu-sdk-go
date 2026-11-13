// 参数来源：docs/Temu_PA_APIs/product/bg.glo.product.search/parameters.json
// 页面更新时间：2025-12-23 16:02:46；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloProductSearchRequest 是 bg.glo.product.search 的业务请求。
type BgGloProductSearchRequest struct {
	ProductSKUIDList []int64 `json:"productSkuIdList,omitzero"` // 可选；货品skuId列表
	PageSize         int64   `json:"pageSize"`                  // 必填；页大小
	PageNum          int64   `json:"pageNum"`                   // 必填；页编号(从1开始)
}

// SKU 对应接口中的嵌套对象。
type SKU struct {
	SKUID int64 `json:"skuId"` // Product SKU ID
}

// SKC 对应接口中的嵌套对象。
type SKC struct {
	SKCID           int64 `json:"skcId"`           // Product SKC ID
	SelectStatus    int64 `json:"selectStatus"`    // Selection Status
	SKUList         []SKU `json:"skuList"`         // SKU List
	ApplyJitStatus  int64 `json:"applyJitStatus"`  // Appeal JIT Status (1-Application Allowed; 3-Application Not Allowed)
	SuggestCloseJit bool  `json:"suggestCloseJit"` // Whether to Suggest Closing JIT Button
}

// Data 对应接口中的嵌套对象。
type Data struct {
	SKCList   []SKC `json:"skcList"`   // SKC List
	ProductID int64 `json:"productId"` // Product ID
}

// ProductSearchResult result
type ProductSearchResult struct {
	Total    int64  `json:"total"`    // Total Count
	DataList []Data `json:"dataList"` // Data List
}

// BgGloProductSearchRespose 是 bg.glo.product.search 的完整响应。
type BgGloProductSearchRespose struct {
	common.Response
	Result *ProductSearchResult `json:"result"` // result
}
