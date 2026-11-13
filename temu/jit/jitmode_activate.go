// 参数来源：docs/Temu_PA_APIs/jit/bg.glo.jitmode.activate/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package jit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloJitmodeActivateRequest 是 bg.glo.jitmode.activate 的业务请求。
type BgGloJitmodeActivateRequest struct {
	ProductID    int64 `json:"productId"`    // 必填；货品 ID
	ProductSKCID int64 `json:"productSkcId"` // 必填；货品 SKC ID
}

// BgGloJitmodeActivateRespose 是 JIT 操作的完整响应。
type BgGloJitmodeActivateRespose struct {
	common.Response
	Result *Result `json:"result"`
}

// Result 对应字典中未定义子字段的 OBJECT，不补造开关或操作状态。
// result 缺失或 null 时为 nil，空对象 {} 对应非 nil 指针。
type Result struct{}
