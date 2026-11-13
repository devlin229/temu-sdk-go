// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.spec.create/parameters.json
// 页面更新时间：2026-01-11 14:19:11；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsSpecCreateRequest 是 bg.glo.goods.spec.create 的业务请求。
type BgGloGoodsSpecCreateRequest struct {
	ParentSpecID int64  `json:"parentSpecId"` // 必填；父规格 ID
	SpecName     string `json:"specName"`     // 必填；子规格名称
}

// BgGloGoodsSpecCreateRespose 是创建自定义规格的完整响应。
type BgGloGoodsSpecCreateRespose struct {
	common.Response
	Result *SpecCreateResult `json:"result"`
}

// SpecCreateResult 包含创建后的子规格 ID，与同包其他接口的 Result 区分。
type SpecCreateResult struct {
	SpecID int64 `json:"specId"` // 子规格 ID
}
