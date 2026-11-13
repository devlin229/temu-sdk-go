// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.accessories.get/parameters.json
// 页面更新时间：2025-09-17 11:01:15；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloGoodsAccessoriesGetRequest 是 bg.glo.goods.accessories.get 的业务请求。
type BgGloGoodsAccessoriesGetRequest struct {
	FuzzyValue string  `json:"fuzzyValue,omitempty"` // 可选；Fuzzy Value
	PageSize   int64   `json:"pageSize,omitempty"`   // 可选；Page Size
	Page       int64   `json:"page,omitempty"`       // 可选；Page Number
	VIDList    []int64 `json:"vidList,omitzero"`     // 可选；Property Value ID List
}

// GoodsAccessoriesData 对应接口中的嵌套对象。
type GoodsAccessoriesData struct {
	VID      int64  `json:"vid"`      // Property Value ID
	UnitName string `json:"unitName"` // Unit Name
	UnitCode int64  `json:"unitCode"` // Unit Code
	Value    string `json:"value"`    // Property Value Name
}

// GoodsAccessoriesResult result
type GoodsAccessoriesResult struct {
	Data       []GoodsAccessoriesData `json:"data"`       // Result List
	TotalCount int64                  `json:"totalCount"` // Total Count
}

// BgGloGoodsAccessoriesGetRespose 是 bg.glo.goods.accessories.get 的完整响应。
type BgGloGoodsAccessoriesGetRespose struct {
	common.Response
	Result *GoodsAccessoriesResult `json:"result"` // result
}
