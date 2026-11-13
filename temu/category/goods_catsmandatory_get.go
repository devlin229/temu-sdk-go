// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.catsmandatory.get/parameters.json
// 页面更新时间：2025-08-27 16:57:43；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// ProductPropertyRequest 对应接口中的嵌套对象。
type ProductPropertyRequest struct {
	VID              int64  `json:"vid"`                        // 必填；Basic Attribute Value ID, pass 0 if none
	ValueUnit        string `json:"valueUnit"`                  // 必填；Unit of Attribute Value, empty string if not available
	PID              int64  `json:"pid"`                        // 必填；Attribute ID
	TemplatePID      int64  `json:"templatePid"`                // 必填；Template Attribute ID
	NumberInputValue string `json:"numberInputValue,omitempty"` // 可选；Numerical Input
	PropValue        string `json:"propValue"`                  // 必填；Basic Property Value
	PropName         string `json:"propName"`                   // 必填；Reference Property Name
	RefPID           int64  `json:"refPid"`                     // 必填；Reference Property ID
}

// BgGloGoodsCatsmandatoryGetRequest 是 bg.glo.goods.catsmandatory.get 的业务请求。
type BgGloGoodsCatsmandatoryGetRequest struct {
	ProductPropertyReqs []ProductPropertyRequest `json:"productPropertyReqs,omitzero"` // 可选；Product Attributes
	BindSiteIDs         []int64                  `json:"bindSiteIds,omitzero"`         // 可选；Bound Site List
	ConfigItems         []int64                  `json:"configItems"`                  // 必填；Category Configuration List
	LeafCatID           int64                    `json:"leafCatId"`                    // 必填；Leaf Category ID
}

// GoodsCatsmandatoryResult result
type GoodsCatsmandatoryResult struct {
	NeedGuideFile bool `json:"needGuideFile"` // Whether the manual is required
}

// BgGloGoodsCatsmandatoryGetRespose 是 bg.glo.goods.catsmandatory.get 的完整响应。
type BgGloGoodsCatsmandatoryGetRespose struct {
	common.Response
	Result *GoodsCatsmandatoryResult `json:"result"` // result
}
