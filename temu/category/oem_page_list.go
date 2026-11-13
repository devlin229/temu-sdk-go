// 参数来源：docs/Temu_PA_APIs/category/bg.glo.oem.page.list/parameters.json
// 页面更新时间：2026-08-14 20:03:27；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloOemPageListRequest 是 bg.glo.oem.page.list 的业务请求。
type BgGloOemPageListRequest struct {
	PageSize int64  `json:"pageSize"`        // 必填；Page Size
	OEMID    string `json:"oemId,omitempty"` // 可选；OE/OEM ID
	Page     int64  `json:"page"`            // 必填；Page Number
}

// PageItems 对应接口中的嵌套对象。
type PageItems struct {
	OEMID string `json:"oemId"` // OE/OEM ID
}

// OEMPageResult result
type OEMPageResult struct {
	Total     int64       `json:"total"`     // Total Count
	PageItems []PageItems `json:"pageItems"` // Current Page Results
}

// BgGloOemPageListRespose 是 bg.glo.oem.page.list 的完整响应。
type BgGloOemPageListRespose struct {
	common.Response
	Result *OEMPageResult `json:"result"` // result
}
