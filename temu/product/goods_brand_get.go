// 参数来源：docs/Temu_PA_APIs/product/bg.glo.goods.brand.get/parameters.json
// 页面更新时间：2026-03-24 18:14:33；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloGoodsBrandGetRequest 是 bg.glo.goods.brand.get 的业务请求。
type BgGloGoodsBrandGetRequest struct {
	VID       int64  `json:"vid,omitempty"`       // 可选；Search Attribute Value ID
	BrandName string `json:"brandName,omitempty"` // 可选；Search Brand Name
	PageSize  int64  `json:"pageSize"`            // 必填；Page Size
	Page      int64  `json:"page"`                // 必填；Page Number
}

// PageItems 对应接口中的嵌套对象。
type PageItems struct {
	VID           int64  `json:"vid"`           // Attribute Value ID
	BrandID       int64  `json:"brandId"`       // Brand ID
	BrandNameEn   string `json:"brandNameEn"`   // Brand English Name
	PID           int64  `json:"pid"`           // Basic Attribute Value ID
	IpBizType     int64  `json:"ipBizType"`     // ip biz Type, If ipBizType=1, ipCodes is required in the product listing interface (bg.glo.goods.add).
	RegSerialCode string `json:"regSerialCode"` // Registration Serial Number
}

// GoodsBrandResult result
type GoodsBrandResult struct {
	Total     int64       `json:"total"`     // Total Count
	PageItems []PageItems `json:"pageItems"` // Current Page Results
}

// BgGloGoodsBrandGetRespose 是 bg.glo.goods.brand.get 的完整响应。
type BgGloGoodsBrandGetRespose struct {
	common.Response
	Result *GoodsBrandResult `json:"result"` // result
}
