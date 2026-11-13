// 参数来源：docs/Temu_PA_APIs/product/bg.glo.goods.topselling.soldout.get/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BindSites 对应接口中的嵌套对象。
type BindSites struct {
	SiteID   int64  `json:"siteId"`   // Site ID
	SiteName string `json:"siteName"` // Site Name
}

// SellOutProducts 对应接口中的嵌套对象。
type SellOutProducts struct {
	SellOutProductID string      `json:"sellOutProductId"` // Sold Out Product ID
	ProductPicture   string      `json:"productPicture"`   // Sold Out Product Main Image
	Categories       *Categories `json:"categories"`       // Sold-out Product Categories
	BindSites        []BindSites `json:"bindSites"`        // Sold-out Product Binding Site List
	ProductName      string      `json:"productName"`      // Sold Out Product Name
}

// GoodsTopsellingSoldoutResult result
type GoodsTopsellingSoldoutResult struct {
	SellOutProducts []SellOutProducts `json:"sellOutProducts"` // Sold Out Product List
}

// BgGloGoodsTopsellingSoldoutGetRespose 是 bg.glo.goods.topselling.soldout.get 的完整响应。
type BgGloGoodsTopsellingSoldoutGetRespose struct {
	common.Response
	Result *GoodsTopsellingSoldoutResult `json:"result"` // result
}
