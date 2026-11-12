package product

import "github.com/devlin229/temu-sdk-go/temu"

// GetGoodsListRequest 是获取商品列表的业务请求参数，所有字段均可选。
// 公共鉴权和签名参数由客户端添加。
type GetGoodsListRequest struct {
	temu.PageParams

	ProductName   string  `json:"productName,omitempty"`   // 商品名称
	ProductSKCIDs []int64 `json:"productSkcIds,omitempty"` // 商品 SKC ID 列表
	SKCExtCode    string  `json:"skcExtCode,omitempty"`    // 商品 SKC 外部编码
	BindSiteIDs   []int64 `json:"bindSiteIds,omitempty"`   // 经营站点 ID 列表

	// SKC 一级状态：0 未发布到站点，100 在售中，200 已下架/已终止，300 已删除。
	SKCTopStatus             *int  `json:"skcTopStatus,omitempty"`
	SKCSiteStatus            *int  `json:"skcSiteStatus,omitempty"`            // SKC 站点状态
	IsSupportPersonalization *bool `json:"isSupportPersonalization,omitempty"` // 是否支持定制商品
	MatchJitMode             *bool `json:"matchJitMode,omitempty"`             // 是否命中 JIT 模式
	QuickSellAgtSignStatus   *int  `json:"quickSellAgtSignStatus,omitempty"`   // 快销协议签署状态：0 未签署，1 已签署

	// 创建时间范围，时间单位按对应接口文档确认，不在 SDK 中自动换算。
	CreatedAtStart int64 `json:"createdAtStart,omitempty"`
	CreatedAtEnd   int64 `json:"createdAtEnd,omitempty"`

	Cat1ID  int64 `json:"cat1Id,omitempty"`  // 一级类目 ID
	Cat2ID  int64 `json:"cat2Id,omitempty"`  // 二级类目 ID
	Cat3ID  int64 `json:"cat3Id,omitempty"`  // 三级类目 ID
	Cat4ID  int64 `json:"cat4Id,omitempty"`  // 四级类目 ID
	Cat5ID  int64 `json:"cat5Id,omitempty"`  // 五级类目 ID
	Cat6ID  int64 `json:"cat6Id,omitempty"`  // 六级类目 ID
	Cat7ID  int64 `json:"cat7Id,omitempty"`  // 七级类目 ID
	Cat8ID  int64 `json:"cat8Id,omitempty"`  // 八级类目 ID
	Cat9ID  int64 `json:"cat9Id,omitempty"`  // 九级类目 ID
	Cat10ID int64 `json:"cat10Id,omitempty"` // 十级类目 ID
}

// GetGoodsListResponse 是商品列表接口的完整响应。
type GetGoodsListResponse struct {
	Result       *GetGoodsListResult `json:"result"`
	Success      bool                `json:"success"`
	ErrorCode    int64               `json:"errorCode"`
	ErrorMessage string              `json:"errorMsg"`
}

// GetGoodsListResult 包含当前页数据与总数量。
type GetGoodsListResult struct {
	Data       []GoodsListItem `json:"data"`
	TotalCount int64           `json:"totalCount"`
}

// GoodsListItem 是列表中的一条商品 SKC 信息。
type GoodsListItem struct {
	ProductProperties        []ProductProperty   `json:"productProperties"`
	ProductID                int64               `json:"productId"`
	ProductJitMode           *ProductJitMode     `json:"productJitMode"`
	ProductSKUSummaries      []ProductSKUSummary `json:"productSkuSummaries"`
	ProductName              string              `json:"productName"`
	CreatedAt                int64               `json:"createdAt"` // 时间单位以接口说明为准
	ProductSemiManaged       *ProductSemiManaged `json:"productSemiManaged"`
	IsSupportPersonalization bool                `json:"isSupportPersonalization"`
	ExtCode                  string              `json:"extCode"` // SKC 外部编码
	LeafCat                  *Category           `json:"leafCat"`
	SKCSiteStatus            int64               `json:"skcSiteStatus"`
	Categories               *ProductCategories  `json:"categories"`
	ProductSKCID             int64               `json:"productSkcId"`
	MainImageURL             string              `json:"mainImageUrl"`
	MatchSKCJitMode          bool                `json:"matchSkcJitMode"`
}

type ProductProperty struct {
	VID              int64  `json:"vid"`
	ValueUnit        string `json:"valueUnit"`
	Language         string `json:"language"`
	PID              int64  `json:"pid"`
	TemplatePID      int64  `json:"templatePid"`
	NumberInputValue string `json:"numberInputValue"`
	PropValue        string `json:"propValue"`
	PropName         string `json:"propName"`
	ValueExtendInfo  string `json:"valueExtendInfo"`
	RefPID           int64  `json:"refPid"`
}

type ProductJitMode struct {
	SignLatestJitVersion   bool  `json:"signLatestJitVersion"`
	QuickSellAgtSignStatus int64 `json:"quickSellAgtSignStatus"` // 0 未签署，1 已签署
	MatchJitMode           bool  `json:"matchJitMode"`
}

type ProductSKUSummary struct {
	ProductSKUID          int64                  `json:"productSkuId"`
	ExtCode               string                 `json:"extCode"`
	ProductSKUWhExtAttr   *ProductSKUWhExtAttr   `json:"productSkuWhExtAttr"`
	VirtualStock          int64                  `json:"virtualStock"`
	ProductSKUSpecList    []ProductSKUSpec       `json:"productSkuSpecList"`
	ProductSKUSaleExtAttr *ProductSKUSaleExtAttr `json:"productSkuSaleExtAttr"`
}

// ProductSKUWhExtAttr 是 SKU 仓库扩展属性，重量为 mg，边长为 mm。
type ProductSKUWhExtAttr struct {
	ProductSKUWeight           *ProductSKUWeight           `json:"productSkuWeight"`
	ProductSKUWMSVolume        *ProductSKUWMSVolume        `json:"productSkuWmsVolume"`
	ProductSKUBarCodes         []ProductSKUBarCode         `json:"productSkuBarCodes"`
	ProductSKUSubSellMode      int64                       `json:"productSkuSubSellMode"`
	ProductSKUSensitiveAttr    *ProductSKUSensitiveAttr    `json:"productSkuSensitiveAttr"`
	ProductSKUFragileLabels    *ProductSKUFragileLabels    `json:"productSkuFragileLabels"`
	ProductSKUNewSensitiveAttr *ProductSKUNewSensitiveAttr `json:"productSkuNewSensitiveAttr"`
	ProductSKUNetWeight        *ProductSKUNetWeight        `json:"productSkuNetWeight"`
	ProductSKUVolumeLabel      *ProductSKUVolumeLabel      `json:"productSkuVolumeLabel"`
	ProductSKUWMSWeight        *ProductSKUWMSWeight        `json:"productSkuWmsWeight"`
	ProductSKUVolume           *ProductSKUVolume           `json:"productSkuVolume"`
	ProductSKUSensitiveLimit   *ProductSKUSensitiveLimit   `json:"productSkuSensitiveLimit"`
	ProductSKUWMSVolumeLabel   *ProductSKUVolumeLabel      `json:"productSkuWmsVolumeLabel"`
}

type ProductSKUWeight struct {
	Value int64 `json:"value"`
} // mg

type ProductSKUWMSVolume struct {
	Len                     int64 `json:"len"` // mm
	WMSCollectionSourceType int64 `json:"wmsCollectionSourceType"`
	Width                   int64 `json:"width"`  // mm
	Height                  int64 `json:"height"` // mm
}

type ProductSKUBarCode struct {
	Code     string `json:"code"`
	CodeType int64  `json:"codeType"` // 1 EAN，2 UPC，3 ISBN
}

type ProductSKUSensitiveAttr struct {
	SensitiveTypes []int64 `json:"sensitiveTypes"` // 1 纯电，2 内电，3 磁性，4 液体，5 粉末，6 膏体，7 刀具
	IsSensitive    int64   `json:"isSensitive"`    // 0 非敏感，1 敏感
}

type ProductSKUFragileLabels struct {
	IsFragile bool `json:"isFragile"`
}

type ProductSKUNewSensitiveAttr struct {
	Force2NormalTypes []int64 `json:"force2NormalTypes"`
	SensitiveList     []int64 `json:"sensitiveList"`
	IsForce2Normal    bool    `json:"isForce2Normal"`
}

type ProductSKUNetWeight struct {
	InputUnit  string `json:"inputUnit"`
	InputValue string `json:"inputValue"`
	Value      int64  `json:"value"` // mg
}

type ProductSKUVolumeLabel struct {
	IsSideOverLength bool `json:"isSideOverLength"`
	IsVolumeOverSize bool `json:"isVolumeOverSize"`
}

type ProductSKUWMSWeight struct {
	WMSCollectionSourceType int64 `json:"wmsCollectionSourceType"`
	Value                   int64 `json:"value"` // mg
}

type ProductSKUVolume struct {
	Len    int64 `json:"len"`    // mm
	Width  int64 `json:"width"`  // mm
	Height int64 `json:"height"` // mm
}

type ProductSKUSensitiveLimit struct {
	MaxBatteryCapacityHp int64 `json:"maxBatteryCapacityHp"` // mWh
	MaxBatteryCapacity   int64 `json:"maxBatteryCapacity"`   // Wh
	MaxLiquidCapacity    int64 `json:"maxLiquidCapacity"`    // mL
	MaxLiquidCapacityHp  int64 `json:"maxLiquidCapacityHp"`  // μL
	MaxKnifeLength       int64 `json:"maxKnifeLength"`       // mm
	MaxKnifeLengthHp     int64 `json:"maxKnifeLengthHp"`     // μm
}

type ProductSKUSpec struct {
	SpecID         int64  `json:"specId"`
	ParentSpecName string `json:"parentSpecName"`
	ParentSpecID   int64  `json:"parentSpecId"`
	SpecName       string `json:"specName"`
}

type ProductSKUSaleExtAttr struct {
	ProductSKUShippingMode       int64 `json:"productSkuShippingMode"`
	ProductSKUIndividuallyPacked int64 `json:"productSkuIndividuallyPacked"`
}

type ProductSemiManaged struct {
	ProductShipment *ProductShipment `json:"productShipment"`
	BindSites       []BindSite       `json:"bindSites"`
}

type ProductShipment struct {
	FreightTemplateID      string  `json:"freightTemplateId"`
	ShipmentLimitSecond    int64   `json:"shipmentLimitSecond"` // 秒
	WarehouseRegionID1List []int64 `json:"warehouseRegionId1List"`
	LongTransport          bool    `json:"longTransport"` // 文档说明 null 视作 false
}

type BindSite struct {
	SiteID   int64  `json:"siteId"`
	SiteName string `json:"siteName"`
}

// Category 复用于叶子类目及各级类目。
type Category struct {
	CatID   int64  `json:"catId"`
	CatName string `json:"catName"`
}

type ProductCategories struct {
	Cat1    *Category `json:"cat1"`
	Cat2    *Category `json:"cat2"`
	Cat3    *Category `json:"cat3"`
	Cat4    *Category `json:"cat4"`
	Cat5    *Category `json:"cat5"`
	Cat6    *Category `json:"cat6"`
	Cat7    *Category `json:"cat7"`
	Cat8    *Category `json:"cat8"`
	Cat9    *Category `json:"cat9"`
	Cat10   *Category `json:"cat10"`
	LeafCat *Category `json:"leafCat"`
}
