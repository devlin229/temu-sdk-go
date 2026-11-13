// 参数来源：docs/Temu_PA_APIs/product/bg.glo.best.seller.invitation.query/parameters.json
// 页面更新时间：2025-09-16 23:04:22；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// BgGloBestSellerInvitationQueryRequest 是 bg.glo.best.seller.invitation.query 的业务请求。
type BgGloBestSellerInvitationQueryRequest struct {
	PageNo    int64   `json:"pageNo"`             // 必填；范围1-100
	PageSize  int64   `json:"pageSize"`           // 必填；范围1-100
	SiteID    int64   `json:"siteId,omitempty"`   // 可选；site id
	CatIDList []int64 `json:"catIdList,omitzero"` // 可选；cat id list
}

// BestSellerInvitationItem 对应接口中的嵌套对象。
type BestSellerInvitationItem struct {
	CatList        []string `json:"catList"`        // Category name
	InvitationName string   `json:"invitationName"` // invitation name
	SiteID         int64    `json:"siteId"`         // Site ID
	SiteName       string   `json:"siteName"`       // Site Name
	InvitationID   int64    `json:"invitationId"`   // invitation ID
	EndTime        int64    `json:"endTime"`        // Deadline
	Type           int64    `json:"type"`           // type 1-mallTarget 2-full2Semi 3-catTarget
	ImageList      []string `json:"imageList"`      // Image
	CatIDList      []int64  `json:"catIdList"`      // Category ID
}

// BestSellerInvitationResult Specific information
type BestSellerInvitationResult struct {
	Total int64                      `json:"total"` // total
	List  []BestSellerInvitationItem `json:"list"`  // list
}

// BgGloBestSellerInvitationQueryRespose 是 bg.glo.best.seller.invitation.query 的完整响应。
type BgGloBestSellerInvitationQueryRespose struct {
	common.Response
	Result *BestSellerInvitationResult `json:"result"` // Specific information
}
