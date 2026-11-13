// 参数来源：docs/Temu_PA_APIs/category/bg.goods.attribute.mapping.global/parameters.json
// 页面更新时间：2026-03-20 17:32:36；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// GoodsProp 对应接口中的嵌套对象。
type GoodsProp struct {
	Values   []string `json:"values,omitzero"`    // 可选；Product English Attribute Values
	PropName string   `json:"propName,omitempty"` // 可选；Product English Attribute Name
}

// BgGoodsAttributeMappingGlobalRequest 是 bg.goods.attribute.mapping.global 的业务请求。
type BgGoodsAttributeMappingGlobalRequest struct {
	CatID        string      `json:"catId"`                  // 必填；Leaf Category ID
	CatName      string      `json:"catName,omitempty"`      // 可选；Leaf Category Name
	GoodsProp    []GoodsProp `json:"goodsProp,omitzero"`     // 可选；Product Attributes
	GoodsName    string      `json:"goodsName"`              // 必填；Product Name
	MainImageURL string      `json:"mainImageUrl,omitempty"` // 可选；Product Main Image
}

// AttributeMappingResult 对应接口中的嵌套对象。
type AttributeMappingResult struct {
	VID         int64  `json:"vid"`         // Value ID
	RefPIDName  string `json:"refPidName"`  // Attribute Name
	PID         int64  `json:"pid"`         // Base property ID
	PIDName     string `json:"pidName"`     // Base property name
	TemplatePID int64  `json:"templatePid"` // Template property ID
	Value       string `json:"value"`       // Value
	RefPID      int64  `json:"refPid"`      // Attribute ID
}

// BgGoodsAttributeMappingGlobalRespose 是 bg.goods.attribute.mapping.global 的完整响应。
type BgGoodsAttributeMappingGlobalRespose struct {
	common.Response
	Result []AttributeMappingResult `json:"result"` // result
}
