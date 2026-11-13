package dictionary

import (
	"errors"
	"slices"
	"unicode"
	"unicode/utf8"
)

// ProvinceName 根据 Temu 省份 ID 查询名称，该 ID 不是行政区划编码。
func (s *Snapshot) ProvinceName(id int64) (string, bool) {
	for _, p := range s.Provinces.Data.Values {
		if p.Region2ID == id {
			return p.Name, true
		}
	}
	return "", false
}

// ProvinceID 按官方完整名称查询，不猜测简称、别名或同名地区。
func (s *Snapshot) ProvinceID(name string) (int64, bool) {
	for _, p := range s.Provinces.Data.Values {
		if p.Name == name {
			return p.Region2ID, true
		}
	}
	return 0, false
}

// SizeChartClassName 接收 classId，不能传入发品的 sizeTemplateId。
func (s *Snapshot) SizeChartClassName(id int64) (string, bool) {
	for _, c := range s.SizeChartClasses.Data.Values {
		if c.ClassID == id {
			return c.ClassName, true
		}
	}
	return "", false
}

// ProductNameLimit 使用一级类目 ID，未列出的类目按当前快照的 default_limit。
func (s *Snapshot) ProductNameLimit(cat1ID int64) int {
	for _, rule := range s.NameLengthRules.Data.Overrides {
		if rule.CatID == cat1ID {
			return rule.Limit
		}
	}
	return s.NameLengthRules.Data.DefaultLimit
}

// CategoryRules 仅表示快照中已列出的规则；车型库仍受适用车型维护条件限制。
type CategoryRules struct {
	BasePlateSupported      bool
	ModelRequired           bool
	VehicleConditionApplies bool
	NetWeightRequired       bool
	TotalPiecesRequired     bool
}

func (s *Snapshot) CategoryRules(categoryID int64) CategoryRules {
	return CategoryRules{
		BasePlateSupported:      slices.Contains(s.BasePlateSupported.Data.CategoryIDs, categoryID),
		ModelRequired:           slices.Contains(s.ModelRequired.Data.CategoryIDs, categoryID),
		VehicleConditionApplies: slices.Contains(s.VehicleRequired.Data.CategoryIDs, categoryID),
		NetWeightRequired:       slices.Contains(s.NetWeightRequired.Data.CategoryIDs, categoryID),
		TotalPiecesRequired:     slices.Contains(s.TotalPiecesRequired.Data.CategoryIDs, categoryID),
	}
}

// LengthRange 表示官方未明确边界字符计数时的可能范围，相等时计数无歧义。
type LengthRange struct {
	Minimum int
	Maximum int
}

// ProductNameLength 不使用 UTF-8 字节长度。ASCII 字母、数字和标点计 1，
// 汉字、非 ASCII 标点等明确的其他字符计 2；非 ASCII 字母/数字、空白及
// ASCII 符号的归类未被原文明确，保留 1..2 范围，不擅自裁剪或规范化标题。
func ProductNameLength(name string) (LengthRange, error) {
	if !utf8.ValidString(name) {
		return LengthRange{}, errors.New("product name must be valid UTF-8")
	}
	var result LengthRange
	for _, r := range name {
		switch {
		case r < utf8.RuneSelf && (unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsPunct(r)):
			result.Minimum++
			result.Maximum++
		case unicode.Is(unicode.Han, r):
			result.Minimum += 2
			result.Maximum += 2
		case unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r < utf8.RuneSelf:
			result.Minimum++
			result.Maximum += 2
		default:
			result.Minimum += 2
			result.Maximum += 2
		}
	}
	return result, nil
}
