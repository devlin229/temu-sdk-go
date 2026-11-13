package product

import (
	"errors"
	"fmt"

	"github.com/devlin229/temu-sdk-go/temu/dictionary"
)

// DictionaryWarning 是快照预检查建议，不代表平台拒绝或完整业务校验结果。
// Message 使用英文；SourceURL 指向该规则的官方页面。
type DictionaryWarning struct {
	Code      string `json:"code"`
	Field     string `json:"field"`
	Message   string `json:"message"`
	SourceURL string `json:"sourceUrl"`
}

// CheckDictionaries 显式执行本地预检查，不修改参数、不发请求、不拦截 SDK 调用。
// 使用最深的非零类目作为叶子类目；不能据此验证类目路径或模板匹配。
// 当前条件必填检查面向全托请求；带 ProductSemiManagedReq 时跳过，避免跨模式套规则。
// 未命中快照不代表没有约束，平台始终是最终依据。
func (r *BgGloGoodsAddRequest) CheckDictionaries(d *dictionary.Snapshot) ([]DictionaryWarning, error) {
	if r == nil {
		return nil, errors.New("bg.glo.goods.add: request must not be nil")
	}
	if d == nil {
		return nil, errors.New("dictionary snapshot must not be nil")
	}
	warnings := make([]DictionaryWarning, 0)
	add := func(code, field, message, url string) {
		warnings = append(warnings, DictionaryWarning{Code: code, Field: field, Message: message, SourceURL: url})
	}
	length, err := dictionary.ProductNameLength(r.ProductName)
	if err != nil {
		add("product_name_encoding", "productName", err.Error(), d.NameLengthRules.SourceURL)
	} else {
		limit := d.ProductNameLimit(r.Cat1ID)
		if length.Minimum > limit {
			add("product_name_too_long", "productName", fmt.Sprintf("product name weighted length is at least %d; snapshot limit is %d", length.Minimum, limit), d.NameLengthRules.SourceURL)
		} else if length.Maximum > limit {
			add("product_name_length_review", "productName", fmt.Sprintf("product name weighted length may be %d..%d; snapshot limit is %d; verify ambiguous character weights", length.Minimum, length.Maximum, limit), d.NameLengthRules.SourceURL)
		}
	}
	if r.ProductWhExtAttrReq != nil {
		origin := r.ProductWhExtAttrReq.ProductOrigin
		if origin.Region1ShortName == "CN" && origin.Region2ID != 0 {
			if _, ok := d.ProvinceName(origin.Region2ID); !ok {
				add("province_not_in_snapshot", "productWhExtAttrReq.productOrigin.region2Id", "province ID is not listed in the snapshot; verify with the platform", d.Provinces.SourceURL)
			}
		}
	}
	if r.ProductSemiManagedReq != nil {
		return warnings, nil
	}
	var leaf int64
	for _, id := range []int64{r.Cat1ID, r.Cat2ID, r.Cat3ID, r.Cat4ID, r.Cat5ID, r.Cat6ID, r.Cat7ID, r.Cat8ID, r.Cat9ID, r.Cat10ID} {
		if id > 0 {
			leaf = id
		}
	}
	rules := d.CategoryRules(leaf)
	if rules.ModelRequired && len(r.GoodsModelReqs) == 0 {
		add("model_required", "goodsModelReqs", "model information is required for this category in the snapshot", d.ModelRequired.SourceURL)
	}
	if rules.VehicleConditionApplies && len(r.VehicleLibraryRelationReqList) == 0 {
		add("vehicle_condition_review", "vehicleLibraryRelationReqList", "verify vehicle fitment: compatible vehicles are required for detail-specific fitment, but not for universal fitment", d.VehicleRequired.SourceURL)
	}
	for skcIndex, skc := range r.ProductSKCReqs {
		for skuIndex, sku := range skc.ProductSKUReqs {
			prefix := fmt.Sprintf("productSkcReqs[%d].productSkuReqs[%d]", skcIndex, skuIndex)
			net := sku.ProductSKUWhExtAttrReq.ProductSKUNetWeightReq
			if rules.NetWeightRequired && (net == nil || net.Value == nil || *net.Value <= 0) {
				add("net_weight_required", prefix+".productSkuWhExtAttrReq.productSkuNetWeightReq.value", "a positive SKU net weight in mg is required for this category in the snapshot", d.NetWeightRequired.SourceURL)
			}
			pack := sku.ProductSKUMultiPackReq
			if rules.TotalPiecesRequired && (pack == nil || pack.NumberOfPiecesNew <= 0) {
				add("total_pieces_required", prefix+".productSkuMultiPackReq.numberOfPiecesNew", "a positive total piece count is required for this full-managed category in the snapshot", d.TotalPiecesRequired.SourceURL)
			}
		}
	}
	return warnings, nil
}
