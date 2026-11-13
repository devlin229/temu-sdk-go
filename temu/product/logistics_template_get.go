// 参数来源：docs/Temu_PA_APIs/product/bg.glo.logistics.template.get/parameters.json
// 页面更新时间：2025-08-27 16:57:43；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloLogisticsTemplateGetRequest 是 bg.glo.logistics.template.get 的业务请求。
type BgGloLogisticsTemplateGetRequest struct {
	SiteIDs []int64 `json:"siteIds,omitzero"` // 可选；Site List
}

// FreightTemplates 对应接口中的嵌套对象。
type FreightTemplates struct {
	FreightTemplateID string `json:"freightTemplateId"` // Shipping Template ID
	TemplateName      string `json:"templateName"`      // Shipping Template Name
}

// LogisticsTemplateResult result
type LogisticsTemplateResult struct {
	FreightTemplates []FreightTemplates `json:"freightTemplates"` // Freight Template Summary Information List
}

// BgGloLogisticsTemplateGetRespose 是 bg.glo.logistics.template.get 的完整响应。
type BgGloLogisticsTemplateGetRespose struct {
	common.Response
	Result *LogisticsTemplateResult `json:"result"` // result
}
