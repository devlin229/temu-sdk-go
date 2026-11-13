// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.cats.get/parameters.json
// 页面更新时间：2026-05-06 18:04:01；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloGoodsCatsGetRequest 是 bg.glo.goods.cats.get 的业务请求。
type BgGloGoodsCatsGetRequest struct {
	ParentCatID int64 `json:"parentCatId,omitempty"` // 可选；Parent Category ID, not passed when querying 1st level list
}

// GoodsCatsResult result
type GoodsCatsResult struct {
	CategoryDTOList []CategoryDTO `json:"categoryDTOList"` // Category Subnode List
}

// BgGloGoodsCatsGetRespose 是 bg.glo.goods.cats.get 的完整响应。
type BgGloGoodsCatsGetRespose struct {
	common.Response
	Result *GoodsCatsResult `json:"result"` // result
}
