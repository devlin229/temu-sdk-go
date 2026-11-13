// 参数来源：docs/Temu_PA_APIs/category/bg.glo.goods.attrs.get/parameters.json
// 页面更新时间：2026-04-29 22:27:22；导出日期：2026-11-12。
package category

import "github.com/devlin229/temu-sdk-go/temu/common"

// BgGloGoodsAttrsGetRequest 是 bg.glo.goods.attrs.get 的业务请求。
type BgGloGoodsAttrsGetRequest struct {
	CatID             int64    `json:"catId"`                       // 必填；叶子类目 ID
	ProductCreateTime int64    `json:"productCreateTime,omitempty"` // 可选；商品创建时间，毫秒时间戳
	SupportedType     int64    `json:"supportedType,omitempty"`     // 可选；渠道类型，传入时优先于 mallId；枚举及默认值未明确
	LangList          []string `json:"langList,omitzero"`           // 可选；多语言查询影响性能，一般不传；nil 省略，空切片发送 []
}

// BgGloGoodsAttrsGetRespose 是类目属性模板查询的完整响应。
type BgGloGoodsAttrsGetRespose struct {
	common.Response
	Result *AttrsResult `json:"result"`
}

// AttrsResult 包含属性模板及自定义规格限制；与父规格查询的 Result 区分。
type AttrsResult struct {
	InputMaxSpecNum      int64      `json:"inputMaxSpecNum"`      // 模板允许的最大自定义规格数量
	ChooseAllQualifySpec bool       `json:"chooseAllQualifySpec"` // 是否全选规格；字典类型为 BOOLEAN，原说明使用 0/1
	SingleSpecValueNum   int64      `json:"singleSpecValueNum"`   // 单个自定义规格值上限
	TemplateID           int64      `json:"templateId"`           // 模板 ID
	Properties           []Property `json:"properties"`           // 模板属性
}

// Property 是一个模板属性，Required 是业务必填标记。
type Property struct {
	NumberInputTitle                string                        `json:"numberInputTitle"`                // 数字输入标题
	TemplatePropertyValueParentList []TemplatePropertyValueParent `json:"templatePropertyValueParentList"` // 属性值关联关系
	Values                          []Value                       `json:"values"`                          // 模板属性值
	ValueUnit                       []string                      `json:"valueUnit"`                       // 属性值单位
	ReferenceType                   int64                         `json:"referenceType"`                   // 属性引用类型
	PID                             int64                         `json:"pid"`                             // 基础属性 ID
	TemplatePID                     int64                         `json:"templatePid"`                     // 模板属性 ID
	Required                        bool                          `json:"required"`                        // 属性是否必填
	InputMaxNum                     int64                         `json:"inputMaxNum"`                     // 最大输入数量，0 表示不允许输入
	PropertyValueType               int64                         `json:"propertyValueType"`               // 属性值类型
	MinValue                        string                        `json:"minValue"`                        // 输入最小值，保留原始字符串
	Feature                         int64                         `json:"feature"`                         // 属性特征
	ValueRule                       int64                         `json:"valueRule"`                       // 值规则
	PropertyChooseTitle             string                        `json:"propertyChooseTitle"`             // 属性选择标题
	ShowType                        int64                         `json:"showType"`                        // 展示类型
	ParentTemplatePID               int64                         `json:"parentTemplatePid"`               // 父模板属性 ID
	MainSale                        bool                          `json:"mainSale"`                        // 是否主销售属性
	ParentSpecID                    int64                         `json:"parentSpecId"`                    // 父规格 ID
	MaxValue                        string                        `json:"maxValue"`                        // 文本最大长度、数值上限或时间上限，保留原始字符串
	Lang2Name                       map[string]string             `json:"lang2Name"`                       // 动态语言代码到属性名的映射
	ChooseMaxNum                    int64                         `json:"chooseMaxNum"`                    // 最大可选数量
	ValuePrecision                  int64                         `json:"valuePrecision"`                  // 最大小数位数，0 表示不允许小数
	ShowCondition                   []ShowCondition               `json:"showCondition"`                   // 属性展示条件
	ControlType                     int64                         `json:"controlType"`                     // 控件类型
	Name                            string                        `json:"name"`                            // 属性名称
	IsSale                          bool                          `json:"isSale"`                          // 是否销售属性，区分普通属性与规格属性
	RefPID                          int64                         `json:"refPid"`                          // 引用属性 ID
}

// TemplatePropertyValueParent 表示属性值关联关系。
type TemplatePropertyValueParent struct {
	ParentVIDList []int64 `json:"parentVidList"` // 关联属性值 ID
	VIDList       []int64 `json:"vidList"`       // 属性值 ID
}

// Value 是模板属性的一个候选值。
type Value struct {
	VID           int64             `json:"vid"`           // 属性值 ID
	SpecID        int64             `json:"specId"`        // 规格 ID
	Lang2Value    map[string]string `json:"lang2Value"`    // 动态语言代码到属性值的映射
	ParentVIDList []int64           `json:"parentVidList"` // 对应父属性值 ID
	ExtendInfo    string            `json:"extendInfo"`    // 扩展信息
	Value         string            `json:"value"`         // 属性值
	Group         *Group            `json:"group"`         // 分组；缺失或 null 时为 nil
}

// Group 是属性值分组。
type Group struct {
	Name string `json:"name"` // 分组名称
	ID   int64  `json:"id"`   // 分组 ID
}

// ShowCondition 表示选中指定父属性值后才能使用该属性。
type ShowCondition struct {
	ParentRefPID int64   `json:"parentRefPid"` // 父引用属性 ID
	ParentVIDs   []int64 `json:"parentVids"`   // 允许展示该属性的父属性值 ID
}
