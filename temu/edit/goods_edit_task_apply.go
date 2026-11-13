// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goods.edit.task.apply/parameters.json
// 页面更新时间：2026-03-06 22:08:06；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsEditTaskApplyRequest 是 bg.glo.goods.edit.task.apply 的业务请求。
type BgGloGoodsEditTaskApplyRequest struct {
	BizIDType        int64      `json:"bizIdType"`        // 必填；Business ID Type (1: Product ID)
	ProductEditItems []EditItem `json:"productEditItems"` // 必填；Edit Item List
	BizIDs           []int64    `json:"bizIds"`           // 必填；Business ID List
}

// BgGloGoodsEditTaskApplyRespose 是 bg.glo.goods.edit.task.apply 的完整响应。
type BgGloGoodsEditTaskApplyRespose struct {
	common.Response
	Result *EditTaskResult `json:"result"` // result
}
