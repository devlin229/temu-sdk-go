// 参数来源：docs/Temu_PA_APIs/image/bg.algo.image.translate.global/parameters.json
// 页面更新时间：2025-12-29 23:57:30；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgAlgoImageTranslateGlobalRequest 是 bg.algo.image.translate.global 的业务请求。
type BgAlgoImageTranslateGlobalRequest struct {
	TargetLang      string `json:"targetLang"`      // 必填；Target language:['es', 'de', 'fr', 'nl', 'it', 'ja', 'ko', 'en','ar']
	ImageURL        string `json:"imageUrl"`        // 必填；Original image URL
	CustomTaskID    int64  `json:"customTaskId"`    // 必填；Merchant custom ID; used for de-duplication in multiple external requests
	IsContainDetail bool   `json:"isContainDetail"` // 必填；Whether to return recognition details
	Language        string `json:"language"`        // 必填；Source language ['zh','en'
	Scene           string `json:"scene,omitempty"` // 可选；Scene name
}

// BgAlgoImageTranslateGlobalRespose 是 bg.algo.image.translate.global 的完整响应。
type BgAlgoImageTranslateGlobalRespose struct {
	common.Response
	Result *TranslateTaskResult `json:"result"` // result
}

// TranslateTaskResult 对应 bg.algo.image.translate.global.result 的嵌套对象。
type TranslateTaskResult struct {
	ResultCode int64  `json:"resultCode"` // Return result
	TaskID     string `json:"taskId"`     // Image processing task Id
	ResultMsg  string `json:"resultMsg"`  // Return result
}
