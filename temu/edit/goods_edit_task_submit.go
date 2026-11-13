// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goods.edit.task.submit/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsEditTaskSubmitRequest 是 bg.glo.goods.edit.task.submit 的业务请求。
type BgGloGoodsEditTaskSubmitRequest struct {
	ProductID           int64      `json:"productId"`                    // 必填；Product ID
	TaskUID             string     `json:"taskUid"`                      // 必填；Modify Task UID
	ProductPropertyReqs []Property `json:"productPropertyReqs,omitzero"` // 可选；Product Attributes
	TaskVersion         int64      `json:"taskVersion"`                  // 必填；Task Version Number
}

// BgGloGoodsEditTaskSubmitRespose 是 bg.glo.goods.edit.task.submit 的完整响应。
type BgGloGoodsEditTaskSubmitRespose struct {
	common.Response
	Result *Result `json:"result"` // result
}
