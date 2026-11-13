// 参数来源：docs/Temu_PA_APIs/image/bg.compliancepicture.get.global/parameters.json
// 页面更新时间：2025-12-29 23:57:47；导出日期：2026-11-12。
package image

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgCompliancepictureGetGlobalRequest 是 bg.compliancepicture.get.global 的业务请求。
type BgCompliancepictureGetGlobalRequest struct {
	URLList []string `json:"urlList,omitzero"` // 可选；List of image URLs
}

// BgCompliancepictureGetGlobalRespose 是 bg.compliancepicture.get.global 的完整响应。
type BgCompliancepictureGetGlobalRespose struct {
	common.Response
	Result *CompliancePictureResult `json:"result"` // result
}

// CompliancePicture 对应 bg.compliancepicture.get.global.list 的嵌套对象。
type CompliancePicture struct {
	HasSpot  string `json:"hasSpot"`  // Return result true or false
	ImageURL string `json:"imageUrl"` // Image URL
}

// CompliancePictureResult 对应 bg.compliancepicture.get.global.result 的嵌套对象。
type CompliancePictureResult struct {
	List []CompliancePicture `json:"list"` // Return result
}
