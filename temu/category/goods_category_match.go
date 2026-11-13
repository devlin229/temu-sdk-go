// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.category.match/parameters.json
// 页面更新时间：2026-04-29 22:27:22；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsCategoryMatchRequest 是 bg.glo.goods.category.match 的业务请求。
type BgGloGoodsCategoryMatchRequest struct {
	SearchText string `json:"searchText,omitempty"` // 可选；搜索文本，文档最大长度 200
	SiteID     int64  `json:"siteId,omitempty"`     // 可选；正数站点 ID，0 时省略
}

// BgGloGoodsCategoryMatchRespose 是类目搜索的完整响应。
type BgGloGoodsCategoryMatchRespose struct {
	common.Response
	Result *CategoryMatchResult `json:"result"`
}

// CategoryMatchResult 包含匹配的类目路径。
type CategoryMatchResult struct {
	CategoryPathDTOs []CategoryPathDTO `json:"categoryPathDTOS"`
}

// CategoryPathDTO 是从一级到十级的类目路径；缺失或 null 的层级为 nil。
type CategoryPathDTO struct {
	Cat1DTO  *CategoryDTO `json:"cat1DTO"`
	Cat2DTO  *CategoryDTO `json:"cat2DTO"`
	Cat3DTO  *CategoryDTO `json:"cat3DTO"`
	Cat4DTO  *CategoryDTO `json:"cat4DTO"`
	Cat5DTO  *CategoryDTO `json:"cat5DTO"`
	Cat6DTO  *CategoryDTO `json:"cat6DTO"`
	Cat7DTO  *CategoryDTO `json:"cat7DTO"`
	Cat8DTO  *CategoryDTO `json:"cat8DTO"`
	Cat9DTO  *CategoryDTO `json:"cat9DTO"`
	Cat10DTO *CategoryDTO `json:"cat10DTO"`
}

// CategoryDTO 是类目节点，保留平台返回的可见性与叶子类目标记。
type CategoryDTO struct {
	CatID       int64  `json:"catId"`       // 类目 ID
	CatName     string `json:"catName"`     // 类目名称
	ParentCatID int64  `json:"parentCatId"` // 父类目 ID
	CatType     int64  `json:"catType"`     // 类目类型
	IsLeaf      bool   `json:"isLeaf"`      // 是否叶子类目
	CatLevel    int64  `json:"catLevel"`    // 类目层级
	IsHidden    bool   `json:"isHidden"`    // 是否隐藏
	HiddenType  int64  `json:"hiddenType"`  // 0 不隐藏，1 正常隐藏，2 旧类目，3 废弃类目
}
