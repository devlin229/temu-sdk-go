package common

// PageParams 是使用 page、pageSize 字段的接口共用分页参数。
// 页码和数量为正数，零值表示未填写并省略，由接口决定默认分页行为。
type PageParams struct {
	Page     int `json:"page,omitempty"`     // 页码
	PageSize int `json:"pageSize,omitempty"` // 每页数量
}
