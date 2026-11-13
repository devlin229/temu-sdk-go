// 参数来源：docs/Temu_PA_APIs/image/bg.goods.image.upload.global/parameters.json
// 页面更新时间：2025-06-22 20:00:52；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// 请求只含业务参数；公共鉴权和签名参数由客户端添加。
// 正数参数使用值类型；需区分显式 0 的可选整数与可选对象使用指针。
// 可选切片使用 omitzero：nil 省略，非 nil 空切片发送 []。
// 必填标记仅表示字段定义；条件必填及业务有效性仍须按商家模式、类目和站点核对。

// BgGoodsImageUploadGlobalRequest 是 bg.goods.image.upload.global 的业务请求。
type BgGoodsImageUploadGlobalRequest struct {
	Image        string   `json:"image,omitempty"`        // 可选；支持格式有：jpg/jpeg、png等图片格式，注意入参图片必须转码为base64编码
	ImageBizType *int64   `json:"imageBizType,omitempty"` // 可选；枚举值：0、1，入参1返回的url用以货品发布时的外包装使用
	Options      *Options `json:"options,omitempty"`      // 可选；-
}

// Options 对应 bg.goods.image.upload.global.options 的对象字段。
type Options struct {
	CateID             int64  `json:"cateId,omitempty"`   // 可选；叶子类目ID，按不同类型进行裁剪，当doIntelligenceCrop=true生效
	DoIntelligenceCrop bool   `json:"doIntelligenceCrop"` // 可选；是否AI智能裁剪，true-根据sizeMode返回一组智能裁剪图（1张原图+3张裁剪图）
	Boost              bool   `json:"boost"`              // 可选；是否AI清晰度提升
	SizeMode           *int64 `json:"sizeMode,omitempty"` // 可选；返回尺寸大小，0-原图大小，1-800*800（1:1），2-1350*1800（3:4）
}

// BgGoodsImageUploadGlobalRespose 是 bg.goods.image.upload.global 的完整响应。
type BgGoodsImageUploadGlobalRespose struct {
	common.Response
	Result *ImageUploadResult `json:"result"` // result
}

// ImageUploadResult result
type ImageUploadResult struct {
	URLs     []string `json:"urls"`     // -
	ImageURL string   `json:"imageUrl"` // -
	URL      string   `json:"url"`      // -
}

// Result 保留原有类型名称，新的调用可使用 ImageUploadResult。
// Deprecated: 使用 ImageUploadResult。
type Result = ImageUploadResult
