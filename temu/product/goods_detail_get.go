// 参数来源：docs/Temu_PA_APIs/product/bg.glo.goods.detail.get/parameters.json
// 页面更新时间：2025-08-27 16:57:39；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloGoodsDetailGetRequest 是 bg.glo.goods.detail.get 的业务请求。
type BgGloGoodsDetailGetRequest struct {
	ProductID int64 `json:"productId"` // 必填；Product ID
}

// CategorySummary 是商品详情中的类目 ID 与名称。
type CategorySummary struct {
	CatID   int64  `json:"catId"`
	CatName string `json:"catName"`
}

// Categories Category
type Categories struct {
	Cat4    *CategorySummary `json:"cat4"`    // Fourth-Level Category
	Cat5    *CategorySummary `json:"cat5"`    // Fifth-Level Category
	Cat2    *CategorySummary `json:"cat2"`    // Secondary Category
	Cat3    *CategorySummary `json:"cat3"`    // Third-Level Category
	Cat1    *CategorySummary `json:"cat1"`    // First-Level Category
	Cat10   *CategorySummary `json:"cat10"`   // Tenth-Level Category
	LeafCat *CategorySummary `json:"leafCat"` // Leaf Category
	CatType int64            `json:"catType"` // Category Type (0: Unclassified, 1: Apparel)
	Cat8    *CategorySummary `json:"cat8"`    // 8th Level Category
	Cat9    *CategorySummary `json:"cat9"`    // Ninth-Level Category
	Cat6    *CategorySummary `json:"cat6"`    // Sixth-Level Category
	Cat7    *CategorySummary `json:"cat7"`    // Seventh-Level Category
}

// GoodsDetailProductOrigin Product Origin Information
type GoodsDetailProductOrigin struct {
	Region2ID        int64  `json:"region2Id"`        // Secondary Region ID
	Region1ShortName string `json:"region1ShortName"` // First-Level Region Abbreviation (Two-Character Code)
}

// ProductWhExtAttr Commodity Warehouse Logistics Supply Chain Extension Properties
type ProductWhExtAttr struct {
	ProductOrigin *GoodsDetailProductOrigin `json:"productOrigin"` // Product Origin Information
}

// GoodsDetailTextModuleDetails Text Module Details
type GoodsDetailTextModuleDetails struct {
	BackgroundColor string `json:"backgroundColor"` // Background Color
	FontFamily      int64  `json:"fontFamily"`      // Font Type
	FontSize        int64  `json:"fontSize"`        // Text Module Font Size
	Align           string `json:"align"`           // Text Alignment, left--Left Aligned；right--Right Aligned；center--Centered；justify--Justified
	FontColor       string `json:"fontColor"`       // Text Color
}

// GoodsDetailContent 对应接口中的嵌套对象。
type GoodsDetailContent struct {
	ImgURL            string                        `json:"imgUrl"`            // Image Address--General
	TextModuleDetails *GoodsDetailTextModuleDetails `json:"textModuleDetails"` // Text Module Details
	Width             int64                         `json:"width"`             // Image Width--General
	Text              string                        `json:"text"`              // Text Information--Text Module
	Height            int64                         `json:"height"`            // Image Height--General
}

// GoodsLayerDecoration 对应接口中的嵌套对象。
type GoodsLayerDecoration struct {
	FloorID     *int64               `json:"floorId"`     // Floor ID, null: Create new, otherwise Update
	Lang        string               `json:"lang"`        // Language Type
	Type        string               `json:"type"`        // Component Type
	Priority    int64                `json:"priority"`    // Floor Sorting
	ContentList []GoodsDetailContent `json:"contentList"` // Floor Content
	Key         string               `json:"key"`         // Floor Type Key
}

// GoodsDetailResult result
type GoodsDetailResult struct {
	ProductID                int64                  `json:"productId"`                // Product ID
	Categories               *Categories            `json:"categories"`               // Category
	ProductWhExtAttr         *ProductWhExtAttr      `json:"productWhExtAttr"`         // Commodity Warehouse Logistics Supply Chain Extension Properties
	GoodsLayerDecorationList []GoodsLayerDecoration `json:"goodsLayerDecorationList"` // Product Detail Decoration Information
	ProductName              string                 `json:"productName"`              // Product Name
}

// BgGloGoodsDetailGetRespose 是 bg.glo.goods.detail.get 的完整响应。
type BgGloGoodsDetailGetRespose struct {
	common.Response
	Result *GoodsDetailResult `json:"result"` // result
}
