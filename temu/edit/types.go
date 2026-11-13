// 编辑域共用的嵌套模型，字段契约见各接口 parameters.json。
package edit

// Property 编辑接口的嵌套对象。
type Property struct {
	VID              int64  `json:"vid"`                        // 必填；Basic Attribute Value ID, pass 0 if none
	ValueUnit        string `json:"valueUnit"`                  // 必填；Unit of Attribute Value, empty string if not available
	PID              int64  `json:"pid"`                        // 必填；Attribute ID
	TemplatePID      int64  `json:"templatePid"`                // 必填；Template Attribute ID
	NumberInputValue string `json:"numberInputValue,omitempty"` // 可选；Numerical Input
	PropValue        string `json:"propValue"`                  // 必填；Basic Property Value
	PropName         string `json:"propName"`                   // 必填；Reference Property Name
	RefPID           int64  `json:"refPid"`                     // 必填；Reference Property ID
}

// PropertyResult 编辑接口的嵌套对象。
type PropertyResult struct {
	GoodsCommitID    int64 `json:"goodsCommitId"`    // Product Draft ID
	MatchEditInAudit bool  `json:"matchEditInAudit"` // Whether to Hit Modify for Review
}

// Video 编辑接口的嵌套对象。
type Video struct {
	VID      string `json:"vid"`      // 必填；Video VID
	CoverURL string `json:"coverUrl"` // 必填；Video Cover Image (B-side stores the first frame image)
	VideoURL string `json:"videoUrl"` // 必填；Video URL
	Width    int64  `json:"width"`    // 必填；Video Width
	Height   int64  `json:"height"`   // 必填；Video Height
}

// SKCRequest 编辑接口的嵌套对象。
type SKCRequest struct {
	SKCID                           int64              `json:"skcId"`                                    // 必填；Product skcId
	SKUCommonReqList                []SKURequest       `json:"skuCommonReqList,omitzero"`                // 可选；SKU Information
	PreviewImgURLs                  []string           `json:"previewImgUrls"`                           // 必填；Carousel List, not required for non-clothing categories
	ProductSKCCarouselImageI18nReqs []ImageI18nRequest `json:"productSkcCarouselImageI18nReqs,omitzero"` // 可选；SKC Carousel Multi-Language Information Request
	ColorImageURL                   string             `json:"colorImageUrl,omitempty"`                  // 可选；SKC Color Block Diagram
}

// SKURequest 编辑接口的嵌套对象。
type SKURequest struct {
	ProductSKUThumbURLI18nReqs []ImageI18nRequest `json:"productSkuThumbUrlI18nReqs,omitzero"` // 可选；SKU Preview Image Multilingual Information Request
	ThumbURL                   string             `json:"thumbUrl"`                            // 必填；Preview Image
	SKUID                      int64              `json:"skuId"`                               // 必填；skuId
}

// ImageI18nRequest 编辑接口的嵌套对象。
type ImageI18nRequest struct {
	ImgURLList *[]string `json:"imgUrlList,omitempty"` // 可选；Image List, empty list implies deletion, null implies no change；nil 省略，指向 nil 切片发送 null，空切片发送 []
	Language   string    `json:"language"`             // 必填；Language
}

// DecorationRequest 编辑接口的嵌套对象。
type DecorationRequest struct {
	FloorID     **int64   `json:"floorId,omitempty"` // 可选；Floor ID, null: Create new, otherwise Update；双层指针区分省略、null 和具体 ID
	Type        string    `json:"type"`              // 必填；Component Type, image - image, text - text
	Priority    int64     `json:"priority"`          // 必填；Floor Sorting
	Lang        string    `json:"lang"`              // 必填；Language Type
	ContentList []Content `json:"contentList"`       // 必填；Floor Content
	Key         string    `json:"key"`               // 必填；Floor Type Key, currently defaults to 'DecImage'
}

// Content 编辑接口的嵌套对象。
type Content struct {
	ImgURL            string      `json:"imgUrl,omitempty"`            // 可选；Image Address--General
	TextModuleDetails *TextModule `json:"textModuleDetails,omitempty"` // 可选；Text Module Details
	Width             int64       `json:"width,omitempty"`             // 可选；Image Width--General
	Text              string      `json:"text,omitempty"`              // 可选；Text Information--Text Module
	Height            int64       `json:"height,omitempty"`            // 可选；Image Height--General
}

// TextModule 编辑接口的嵌套对象。
type TextModule struct {
	BackgroundColor string `json:"backgroundColor"`      // 必填；Background Color
	FontFamily      *int64 `json:"fontFamily,omitempty"` // 可选；Font Type
	FontSize        int64  `json:"fontSize"`             // 必填；Text Module Font Size
	Align           string `json:"align"`                // 必填；Text Alignment, left--Left Aligned；right--Right Aligned；center--Centered；justify--Justified
	FontColor       string `json:"fontColor"`            // 必填；Text Color
}

// Result 编辑接口的嵌套对象。
type Result struct {
}

// SensitiveSKURequest 编辑接口的嵌套对象。
type SensitiveSKURequest struct {
	ProductSKUSensitiveLimitReq SensitiveLimitRequest `json:"productSkuSensitiveLimitReq"` // 必填；Product SKU Sensitive Attribute Restriction Request (Edit Scenario, Empty Object if No Restriction)
	ProductSKUID                int64                 `json:"productSkuId"`                // 必填；Product SKU ID
	ProductSKUSensitiveAttrReq  SensitiveAttrRequest  `json:"productSkuSensitiveAttrReq"`  // 必填；Product SKU Sensitive Attribute Request
}

// SensitiveLimitRequest 编辑接口的嵌套对象。
type SensitiveLimitRequest struct {
	MaxBatteryCapacityHp *int64         `json:"maxBatteryCapacityHp,omitempty"` // 可选；Maximum Battery Capacity (mWh)
	MaxBatteryCapacity   *int64         `json:"maxBatteryCapacity,omitempty"`   // 可选；Maximum Battery Capacity (Wh) (Prioritizes maxBatteryCapacityHp)
	MaxLiquidCapacity    *int64         `json:"maxLiquidCapacity,omitempty"`    // 可选；Maximum Liquid Capacity (mL) (Prefer using maxLiquidCapacityHp)
	MaxLiquidCapacityHp  *int64         `json:"maxLiquidCapacityHp,omitempty"`  // 可选；Maximum Liquid Capacity (μL)
	MaxKnifeLength       *int64         `json:"maxKnifeLength,omitempty"`       // 可选；Maximum Tool Length (mm) (Prefer using maxKnifeLengthHp)
	MaxKnifeLengthHp     *int64         `json:"maxKnifeLengthHp,omitempty"`     // 可选；Maximum Tool Length (μm)
	KnifeTipAngle        *KnifeTipAngle `json:"knifeTipAngle,omitempty"`        // 可选；Blade Angle
}

// KnifeTipAngle 编辑接口的嵌套对象。
type KnifeTipAngle struct {
	Degrees int64 `json:"degrees"` // 必填；Degree
}

// SensitiveAttrRequest 编辑接口的嵌套对象。
type SensitiveAttrRequest struct {
	SensitiveTypes []int64 `json:"sensitiveTypes,omitzero"` // 可选；Sensitive Type PURE_ELECTRIC(1, "Pure Electric"), INTERNAL_ELECTRIC(2, "Internal Electric"), MAGNETISM(3, "Magnetism"), LIQUID(4, "Liquid"), POWDER(5, "Powder"), PASTE(6, "Paste"), CUTTER(7, "Tool")
	IsSensitive    *int64  `json:"isSensitive,omitempty"`   // 可选；Whether Sensitive Attribute, 0: Non-Sensitive, 1: Sensitive
	SensitiveList  []int64 `json:"sensitiveList,omitzero"`  // 可选；Sensitive Type PURE_ELECTRIC(110001,"Pure Electric"), INTERNAL_ELECTRIC(120001, "Internal Electric"), MAGNETISM(130001, "Magnetism"), LIQUID(140001, "Liquid"), POWDER(150001, "Powder"), PASTE(160001, "Ointment"), CUTTER(170001, "Tool")
}

// EditItem 编辑接口的嵌套对象。
type EditItem struct {
	EditAdvice string `json:"editAdvice"` // 必填；Edit Advice
	EditItem   int64  `json:"editItem"`   // 必填；Edit Item (3: Product Property)
}

// EditTaskResult 编辑接口的嵌套对象。
type EditTaskResult struct {
	AddFailedList           []FailedItem      `json:"addFailedList"`           // List of Failed Additions
	HandleFailedList        []FailedItem      `json:"handleFailedList"`        // Failed Processing List
	SuccessBizID2TaskUIDMap map[string]string `json:"successBizId2TaskUidMap"` // Successful id -> Task uid
}

// FailedItem 编辑接口的嵌套对象。
type FailedItem struct {
	FailedBizID  int64  `json:"failedBizId"`  // Failed Business ID
	FailedReason string `json:"failedReason"` // Failure Reason
}

// WarehouseExtAttrRequest 编辑接口的嵌套对象。
type WarehouseExtAttrRequest struct {
	ProductOrigin ProductOrigin `json:"productOrigin"` // 必填；Product Origin
}

// ProductOrigin 编辑接口的嵌套对象。
type ProductOrigin struct {
	Region2ID        int64  `json:"region2Id,omitempty"` // 可选；Secondary Region ID
	Region1ShortName string `json:"region1ShortName"`    // 必填；First-Level Region Abbreviation (Two-Character Code)
}

// ShipmentRequest 编辑接口的嵌套对象。
type ShipmentRequest struct {
	FreightTemplateID   string `json:"freightTemplateId"`   // 必填；Shipping Template ID
	ShipmentLimitSecond int64  `json:"shipmentLimitSecond"` // 必填；Promised Delivery Time (unit: s)
}
