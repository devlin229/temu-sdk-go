// 参数来源：docs/Temu_PA_APIs/image/bg.algo.image.translate.result.global/parameters.json
// 页面更新时间：2025-12-29 23:56:58；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgAlgoImageTranslateResultGlobalRequest 是 bg.algo.image.translate.result.global 的业务请求。
type BgAlgoImageTranslateResultGlobalRequest struct {
	TaskID string `json:"taskId"` // 必填；Image processing task Id
}

// BgAlgoImageTranslateResultGlobalRespose 是 bg.algo.image.translate.result.global 的完整响应。
type BgAlgoImageTranslateResultGlobalRespose struct {
	common.Response
	Result *TranslationResult `json:"result"` // result
}

// TextBox 对应 bg.algo.image.translate.result.global.textBoxList 的嵌套对象。
type TextBox struct {
	TranslatedText string  `json:"translatedText"` // Translated text content
	Polygon        []int64 `json:"polygon"`        // The four corners of the text box, as [top-left x, top-left y, top-right x, top-right y, bottom-left x, bottom-left y, bottom-right x, bottom-right y]
	NeedTranslate  bool    `json:"needTranslate"`  // Whether translation is needed
	Text           string  `json:"text"`           // Original text content
	TextColor      []int64 `json:"textColor"`      // Text color, rgb color, format is [red, green, blue]
}

// TranslationResult 对应 bg.algo.image.translate.result.global.result 的嵌套对象。
type TranslationResult struct {
	ImageBackgroundURL string    `json:"imageBackgroundUrl"` // Background image URL
	ImageURL           string    `json:"imageUrl"`           // original image url
	TextBoxList        []TextBox `json:"textBoxList"`        // All recognized text boxes
	ResultCode         int64     `json:"resultCode"`         // resultCode
	NeedTranslate      bool      `json:"needTranslate"`      // Needs translation
	ImageResultURL     string    `json:"imageResultUrl"`     // Result image URL
	HasText            bool      `json:"hasText"`            // Has text
	ResultMsg          string    `json:"resultMsg"`          // resultMsg
}
