// 参数来源：docs/Temu_PA_APIs/image/bg.glo.colorimageurl.get/parameters.json
// 页面更新时间：2025-06-26 20:25:16；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloColorimageurlGetRequest 是 bg.glo.colorimageurl.get 的业务请求。
type BgGloColorimageurlGetRequest struct {
	RGBA     *RGBA       `json:"rgba,omitempty"`     // 可选；可选项，颜色值。
	Coor     *Coordinate `json:"coor,omitempty"`     // 可选；可选项，商品服饰的中心点坐标，如果不传会自动计算。如果图片尺寸为 400 × 500，服装中心点在 200×300， 那 x 就传 200， y 就传 300。
	ImageURL string      `json:"imageUrl,omitempty"` // 可选；图片链接。
}

// BgGloColorimageurlGetRespose 是 bg.glo.colorimageurl.get 的完整响应。
type BgGloColorimageurlGetRespose struct {
	common.Response
	Result *ColorImageResult `json:"result"` // 结果对象
}

// RGBA 对应 bg.glo.colorimageurl.get.rgba 的嵌套对象。
type RGBA struct {
	A *int64 `json:"a,omitempty"` // 可选；可选项，透明度，冗余参数，默认赋值为1。；nil 省略，指向 0 时发送 0
	R *int64 `json:"r,omitempty"` // 可选；可选项，red，红色色值，范围为0到255。；nil 省略，指向 0 时发送 0
	B *int64 `json:"b,omitempty"` // 可选；可选项，blue，蓝色色值，范围为0到255。；nil 省略，指向 0 时发送 0
	G *int64 `json:"g,omitempty"` // 可选；可选项，green，绿色色值，范围为0到255。；nil 省略，指向 0 时发送 0
}

// Coordinate 对应 bg.glo.colorimageurl.get.coor 的嵌套对象。
type Coordinate struct {
	X *int64 `json:"x,omitempty"` // 可选；可选项，商品服饰中心点横坐标。例如传 200。；nil 省略，指向 0 时发送 0
	Y *int64 `json:"y,omitempty"` // 可选；可选项，商品服饰中心点纵坐标。例如传 300。；nil 省略，指向 0 时发送 0
}

// RGBAResult 对应 bg.glo.colorimageurl.get.rgbaRec 的嵌套对象。
type RGBAResult struct {
	A int64 `json:"a"` // 透明度，默认赋值为1。
	R int64 `json:"r"` // red，红色色值，范围为0到255。
	B int64 `json:"b"` // blue，蓝色色值，范围为0到255。
	G int64 `json:"g"` // green，绿色色值，范围为0到255。
}

// ColorImageResult 对应 bg.glo.colorimageurl.get.result 的嵌套对象。
type ColorImageResult struct {
	RGBARec    *RGBAResult `json:"rgbaRec"`    // 结果对象的颜色信息
	Confidence string      `json:"confidence"` // 置信度，[0, 1]，为 1 时置信度最高。
	ResultURL  string      `json:"resultUrl"`  // 结果图片 URL，公网可直接访问，永久有效，但建议调用方自己转存一份。
}
