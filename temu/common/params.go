package common

// PageParams 是使用 page、pageSize 字段的接口共用分页参数。
// 零值时省略，由对应接口决定默认分页行为。
type PageParams struct {
	Page     int `json:"page,omitempty"`     // 页码
	PageSize int `json:"pageSize,omitempty"` // 每页数量
}
