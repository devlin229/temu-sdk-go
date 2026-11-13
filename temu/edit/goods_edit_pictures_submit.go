// 参数来源：docs/Temu_PA_APIs/edit/bg.glo.goods.edit.pictures.submit/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package edit

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsEditPicturesSubmitRequest 是 bg.glo.goods.edit.pictures.submit 的业务请求。
type BgGloGoodsEditPicturesSubmitRequest struct {
	ProductCarouseVideoReqList *[]Video            `json:"productCarouseVideoReqList,omitempty"` // 可选；Main Image Video, empty list implies deletion, null implies no change；nil 省略，指向 nil 切片发送 null，空切片发送 []
	SKCList                    []SKCRequest        `json:"skcList,omitzero"`                     // 可选；SKC Information
	ProductID                  int64               `json:"productId"`                            // 必填；Product ID
	ProductDetailVideoReqList  *[]Video            `json:"productDetailVideoReqList,omitempty"`  // 可选；Detail Video, empty list implies deletion, null implies no change；nil 省略，指向 nil 切片发送 null，空切片发送 []
	GoodsLayerDecorationReqs   []DecorationRequest `json:"goodsLayerDecorationReqs,omitzero"`    // 可选；Product Details Decoration
	CarouselImageURLs          []string            `json:"carouselImageUrls,omitzero"`           // 可选；Commodity Carousel, not required for apparel category, will be aggregated from SKC
	CarouselImageI18nReqs      []ImageI18nRequest  `json:"carouselImageI18nReqs,omitzero"`       // 可选；Commodity Carousel Multi-Language Info Request, not required for apparel category, will be aggregated from SKC
	MaterialImgURL             string              `json:"materialImgUrl,omitempty"`             // 可选；Material Image
	MaterialMultiLanguages     []string            `json:"materialMultiLanguages,omitzero"`      // 可选；Image Multilingual List
}

// BgGloGoodsEditPicturesSubmitRespose 是 bg.glo.goods.edit.pictures.submit 的完整响应。
type BgGloGoodsEditPicturesSubmitRespose struct {
	common.Response
	Result *Result `json:"result"` // result
}
