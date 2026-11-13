// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.parentspec.get/parameters.json
// 页面更新时间：2026-01-11 14:19:11；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsParentspecGetRespose 是查询父规格列表的完整响应。
type BgGloGoodsParentspecGetRespose struct {
	common.Response
	Result *ParentSpecResult `json:"result"`
}

// ParentSpecResult 包含父规格列表，JSON 字段名保留官方 parentSpecDTOS 拼写。
type ParentSpecResult struct {
	ParentSpecDTOs []ParentSpecDTO `json:"parentSpecDTOS"`
}

// ParentSpecDTO 是一个父规格，普通字段使用值类型。
type ParentSpecDTO struct {
	ParentSpecName string `json:"parentSpecName"` // 父规格名称
	ParentSpecID   int64  `json:"parentSpecId"`   // 父规格 ID
}

// Result 保留原有类型名称，新的调用可使用 ParentSpecResult。
// Deprecated: 使用 ParentSpecResult。
type Result = ParentSpecResult
