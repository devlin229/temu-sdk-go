// 参数来源：docs/Temu_PA_APIs/image/bg.glo.fancy.image.cm2in/parameters.json
// 页面更新时间：2025-06-26 20:25:16；导出日期：2026-11-12。
package image

// BgGloFancyImageCm2inRequest 是 bg.glo.fancy.image.cm2in 的业务请求。
type BgGloFancyImageCm2inRequest struct {
	ImageURL string `json:"imageUrl"` // 必填；图片链接
}

// BgGloFancyImageCm2inRespose 是 bg.glo.fancy.image.cm2in 的完整响应。
type BgGloFancyImageCm2inRespose struct {
	Code      int64  `json:"code"`      // 错误码
	ErrMsg    string `json:"errMsg"`    // 错误信息
	ResultURL string `json:"resultUrl"` // 图片链接，无需加签，永久有效期，但建议自己转存一份
}
