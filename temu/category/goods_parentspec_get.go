package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsParentspecGetRespose 是查询父规格列表的完整响应。
type BgGloGoodsParentspecGetRespose struct {
	common.Response
	Result *Result `json:"result"`
}

// Result 包含父规格列表，JSON 字段名保留官方 parentSpecDTOS 拼写。
type Result struct {
	ParentSpecDTOs []ParentSpecDTO `json:"parentSpecDTOS"`
}

// ParentSpecDTO 是一个父规格，普通字段使用值类型。
type ParentSpecDTO struct {
	ParentSpecName string `json:"parentSpecName"` // 父规格名称
	ParentSpecID   int64  `json:"parentSpecId"`   // 父规格 ID
}
