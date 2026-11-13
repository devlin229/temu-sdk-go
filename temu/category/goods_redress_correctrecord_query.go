// 参数来源：docs/Temu_PA_APIs/category/bg.goods.redress.correctrecord.query/parameters.json
// 页面更新时间：2025-12-19 22:01:40；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGoodsRedressCorrectrecordQueryRequest 是 bg.goods.redress.correctrecord.query 的业务请求。
type BgGoodsRedressCorrectrecordQueryRequest struct {
	QueryNeedCorrect *bool   `json:"queryNeedCorrect,omitempty"` // 可选；Only query the records that need correct
	ProductIDList    []int64 `json:"productIdList,omitzero"`     // 可选；Product ID
	PageSize         int64   `json:"pageSize"`                   // 必填；Page Size, range [1, 20]
	PageNum          int64   `json:"pageNum"`                    // 必填；Page Number, range [1, 500]
}

// CorrectionCategory 是类目纠正记录中的路径节点。
type CorrectionCategory struct {
	CatID       int64  `json:"catId"`
	CatName     string `json:"catName"`
	ParentCatID int64  `json:"parentCatId"`
	IsLeaf      bool   `json:"isLeaf"`
	CatLevel    int64  `json:"catLevel"`
}

// RedressCorrectrecordItem 对应接口中的嵌套对象。
type RedressCorrectrecordItem struct {
	CurrentCategory          []CorrectionCategory   `json:"currentCategory"`          // Product current category path
	ProductID                int64                  `json:"productId"`                // Product ID
	CorrectRecordID          int64                  `json:"correctRecordId"`          // Correct record ID
	CorrectedCategory        []CorrectionCategory   `json:"correctedCategory"`        // Categories that have been corrected
	BeforeCorrectionCategory []CorrectionCategory   `json:"beforeCorrectionCategory"` // Category path before this rectification
	OptionalCategories       [][]CorrectionCategory `json:"optionalCategories"`       // Recommended categories list
	CorrectDeadline          int64                  `json:"correctDeadline"`          // Auto-correction deadline (ms)
	Status                   int64                  `json:"status"`                   // Optional value description:[1:correction record is waiting for feedback;2:correction record is processing, cannot feedback;3:correction record is finished;4:correction record is failed;]
}

// RedressCorrectrecordResult result
type RedressCorrectrecordResult struct {
	Total    int64                      `json:"total"`    // Total
	PageSize int64                      `json:"pageSize"` // Page Size
	List     []RedressCorrectrecordItem `json:"list"`     // Correct record list
	PageNum  int64                      `json:"pageNum"`  // Page Number
}

// BgGoodsRedressCorrectrecordQueryRespose 是 bg.goods.redress.correctrecord.query 的完整响应。
type BgGoodsRedressCorrectrecordQueryRespose struct {
	common.Response
	Result *RedressCorrectrecordResult `json:"result"` // result
}
