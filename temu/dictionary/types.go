// Package dictionary 提供官方数据字典的版本快照和辅助查询，不改变接口请求行为。
// 快照不保证覆盖未来类目或枚举；PA 适用性和条件规则见各 Document 的来源说明。
package dictionary

// Source 保留官方页面、更新时间、提取日期与适用性，便于调用方核对版本。
type Source struct {
	Name            string   `json:"name"`
	SourceURL       string   `json:"source_url"`
	SourceUpdatedAt string   `json:"source_updated_at"`
	ExtractedAt     string   `json:"extracted_at"`
	Recommendation  string   `json:"recommendation"`
	Purpose         string   `json:"purpose"`
	PACompatibility string   `json:"pa_compatibility"`
	Notes           []string `json:"notes"`
}

// Document 将来源与对应的明确数据结构放在一起；不内置重复的页面原文。
type Document[T any] struct {
	Source
	Data T `json:"data"`
}

type Province struct {
	Region2ID int64  `json:"region2Id"`
	Name      string `json:"name"`
}
type Provinces struct {
	Field  string     `json:"field"`
	Values []Province `json:"values"`
}
type SizeChartClass struct {
	ClassID   int64  `json:"classId"`
	ClassName string `json:"className"`
}
type SizeChartClasses struct {
	Values []SizeChartClass `json:"values"`
}

// CategoryList 中的 ID 为快照列出的类目，未列出不表示未来也不适用。
type CategoryList struct {
	CategoryIDs []int64 `json:"category_ids"`
}
type VehicleRequired struct {
	CategoryList
	SemiManagedSiteIDs []int64 `json:"semi_managed_site_ids"`
	Condition          string  `json:"condition"`
}
type NetWeightRequired struct {
	CategoryList
	Field string `json:"field"`
	Unit  string `json:"unit"`
}
type TotalPiecesRequired struct {
	CategoryList
	Field string `json:"field"`
	Scope string `json:"scope"`
}

// NameLimit 的 CatName 可为 null，不根据类目 ID 补造名称。
type NameLimit struct {
	CatID   int64   `json:"catId"`
	CatName *string `json:"catName"`
	Limit   int     `json:"limit"`
}
type NameLengthRules struct {
	Field         string      `json:"field"`
	CategoryLevel int         `json:"category_level"`
	DefaultLimit  int         `json:"default_limit"`
	CountingRule  string      `json:"counting_rule"`
	Overrides     []NameLimit `json:"overrides"`
}
type Technology struct {
	Code int64  `json:"code"`
	Name string `json:"name"`
}

// CustomizedTechnology 的层级外层键含义未明确，不推断对应商家模式或 technologyType。
type CustomizedTechnology struct {
	FirstTypes             []Technology                  `json:"first_types"`
	SecondTypes            []Technology                  `json:"second_types"`
	Hierarchy              map[string]map[string][]int64 `json:"hierarchy"`
	UnnamedSecondTypeCodes []int64                       `json:"unnamed_second_type_codes"`
}

// SizeSpecGroup 保留原始顺序、重复 specId、缺省 specIds 和尺码子类型。
type SizeSpecGroup struct {
	SizeType           int64         `json:"sizeType"`
	GroupIDs           []int64       `json:"groupIds"`
	SpecIDs            []int64       `json:"specIds,omitzero"`
	SizeSubTypeConfigs []SizeSubType `json:"sizeSubTypeConfigs,omitzero"`
}
type SizeSubType struct {
	SizeSubType int64   `json:"sizeSubType"`
	SpecIDs     []int64 `json:"specIds"`
}
type SizeSpecGroups struct {
	Groups []SizeSpecGroup `json:"groups"`
}

// Snapshot 是独立可维护的字典集合，适合在批量操作开始时加载一次复用。
// 可在发品前修改本实例的数据；不要一边查询一边并发修改。
type Snapshot struct {
	Provinces            Document[Provinces]
	SizeChartClasses     Document[SizeChartClasses]
	BasePlateSupported   Document[CategoryList]
	ModelRequired        Document[CategoryList]
	VehicleRequired      Document[VehicleRequired]
	NetWeightRequired    Document[NetWeightRequired]
	TotalPiecesRequired  Document[TotalPiecesRequired]
	NameLengthRules      Document[NameLengthRules]
	CustomizedTechnology Document[CustomizedTechnology]
	SizeSpecGroups       Document[SizeSpecGroups]
}
