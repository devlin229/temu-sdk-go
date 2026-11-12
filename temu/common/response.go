package common

// Response 是各接口共用的响应状态，匿名嵌入后字段仍位于 JSON 根对象。
type Response struct {
	Success      bool   `json:"success"`   // 是否成功
	ErrorCode    int64  `json:"errorCode"` // 平台错误码
	ErrorMessage string `json:"errorMsg"`  // 平台错误信息
}
