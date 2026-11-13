// 参数来源：docs/Temu_PA_APIs/image/bg.goods.texttopicture.add.global/parameters.json
// 页面更新时间：2025-06-23 13:55:39；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGoodsTexttopictureAddGlobalRequest 是 bg.goods.texttopicture.add.global 的业务请求。
type BgGoodsTexttopictureAddGlobalRequest struct {
	BackColor string `json:"backColor"` // 必填；背景颜色,必须以#开头，后面的数字是十六进制，不指定透明通道的话总共7个字符，指定透明通道的话总共9个字符
	Text      string `json:"text"`      // 必填；文本
	Align     string `json:"align"`     // 必填；对齐方式,left,center,right
	FontColor string `json:"fontColor"` // 必填；字体颜色,必须以#开头，后面的数字是十六进制，不指定透明通道的话总共7个字符，指定透明通道的话总共9个字符
	Font      string `json:"font"`      // 必填；字体,Source Han Sans CN Heavy/Bold/Medium/Regular/Light/Extralight, Source Han Serif Heavy/Bold
}

// BgGoodsTexttopictureAddGlobalRespose 是 bg.goods.texttopicture.add.global 的完整响应。
type BgGoodsTexttopictureAddGlobalRespose struct {
	common.Response
	Result *TextToPictureResult `json:"result"` // result
}

// TextToPictureResult 对应 bg.goods.texttopicture.add.global.result 的嵌套对象。
type TextToPictureResult struct {
	ImageURL string `json:"imageUrl"` // -
}
