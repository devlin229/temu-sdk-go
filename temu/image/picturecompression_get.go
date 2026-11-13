// 参数来源：docs/Temu_PA_APIs/image/bg.glo.picturecompression.get/parameters.json
// 页面更新时间：2025-06-26 20:25:16；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloPicturecompressionGetRequest 是 bg.glo.picturecompression.get 的业务请求。
type BgGloPicturecompressionGetRequest struct {
	URLs []string `json:"urls,omitzero"` // 可选；图片链接数组，数量建议不超过30个
}

// BgGloPicturecompressionGetRespose 是 bg.glo.picturecompression.get 的完整响应。
type BgGloPicturecompressionGetRespose struct {
	common.Response
	Result *PictureCompressionResult `json:"result"` // 结果
}

// CompressedPicture 对应 bg.glo.picturecompression.get.results 的嵌套对象。
type CompressedPicture struct {
	Size      int64  `json:"size"`      // 结果图片存储空间占用，单位 Byte。例如 724955，就是图片有 700 KB 左右。
	OriginURL string `json:"originUrl"` // 原始图片链接。
	Width     int64  `json:"width"`     // 结果图片宽度，单位：像素。例如 1500。
	ResultURL string `json:"resultUrl"` // 结果图片链接。公网可直接访问，永久有效，但建议调用方自己转存一份。
	Height    int64  `json:"height"`    // 结果图片高度，单位：像素。例如 2000。
}

// PictureCompressionResult 对应 bg.glo.picturecompression.get.result 的嵌套对象。
type PictureCompressionResult struct {
	Results []CompressedPicture `json:"results"` // 结果数组
}
