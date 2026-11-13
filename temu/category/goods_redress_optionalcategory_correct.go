// 参数来源：docs/Temu_PA_APIs/category/bg.goods.redress.optionalcategory.correct/parameters.json
// 页面更新时间：2025-10-30 16:36:39；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// SelectedCategory 对应接口中的嵌套对象。
type SelectedCategory struct {
	ProductID          int64 `json:"productId"`          // 必填；Product ID
	CorrectRecordID    int64 `json:"correctRecordId"`    // 必填；Correct record Id
	SelectedCategoryID int64 `json:"selectedCategoryId"` // 必填；Selected leaf category ID
}

// BgGoodsRedressOptionalcategoryCorrectRequest 是 bg.goods.redress.optionalcategory.correct 的业务请求。
type BgGoodsRedressOptionalcategoryCorrectRequest struct {
	List []SelectedCategory `json:"list"` // 必填；selected category list
}

// CorrectionOutcome 对应接口中的嵌套对象。
type CorrectionOutcome struct {
	Reason    string `json:"reason"`    // The reason for the operation result
	ProductID int64  `json:"productId"` // Product ID
	IsSuccess bool   `json:"isSuccess"` // Is this operation successful
}

// RedressOptionalcategoryCorrectResult result
type RedressOptionalcategoryCorrectResult struct {
	Total        int64               `json:"total"`        // The quantity of items in the batch operation
	SuccessCount int64               `json:"successCount"` // Count of successful operations
	List         []CorrectionOutcome `json:"list"`         // Result details list
}

// BgGoodsRedressOptionalcategoryCorrectRespose 是 bg.goods.redress.optionalcategory.correct 的完整响应。
type BgGoodsRedressOptionalcategoryCorrectRespose struct {
	common.Response
	Result *RedressOptionalcategoryCorrectResult `json:"result"` // result
}
