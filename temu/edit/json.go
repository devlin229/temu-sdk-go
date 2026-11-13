package edit

import (
	"bytes"
	"encoding/json"
)

// decodeNullableFields 保留字段是否显式传 null，同时保持示例请求的严格字段校验。
// 标准库直接解码到指针会把 null 和省略都置为 nil，导致再次发送时语义丢失。
func decodeNullableFields(data []byte, target any) (map[string]json.RawMessage, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

func isNull(raw json.RawMessage) bool {
	return bytes.Equal(bytes.TrimSpace(raw), []byte("null"))
}

// UnmarshalJSON 保留视频列表的省略、null、空数组及非空数组状态。
func (r *BgGloGoodsEditPicturesSubmitRequest) UnmarshalJSON(data []byte) error {
	type plain BgGloGoodsEditPicturesSubmitRequest
	decoded := plain(*r)
	fields, err := decodeNullableFields(data, &decoded)
	if err != nil {
		return err
	}
	if isNull(fields["productCarouseVideoReqList"]) {
		var videos []Video
		decoded.ProductCarouseVideoReqList = &videos
	}
	if isNull(fields["productDetailVideoReqList"]) {
		var videos []Video
		decoded.ProductDetailVideoReqList = &videos
	}
	*r = BgGloGoodsEditPicturesSubmitRequest(decoded)
	return nil
}

// UnmarshalJSON 保留多语言图片列表的显式 null（不修改）和 []（删除）。
func (r *ImageI18nRequest) UnmarshalJSON(data []byte) error {
	type plain ImageI18nRequest
	decoded := plain(*r)
	fields, err := decodeNullableFields(data, &decoded)
	if err != nil {
		return err
	}
	if isNull(fields["imgUrlList"]) {
		var images []string
		decoded.ImgURLList = &images
	}
	*r = ImageI18nRequest(decoded)
	return nil
}

// UnmarshalJSON 保留 floorId: null 的新建语义，避免被 omitempty 省略。
func (r *DecorationRequest) UnmarshalJSON(data []byte) error {
	type plain DecorationRequest
	decoded := plain(*r)
	fields, err := decodeNullableFields(data, &decoded)
	if err != nil {
		return err
	}
	if isNull(fields["floorId"]) {
		var id *int64
		decoded.FloorID = &id
	}
	*r = DecorationRequest(decoded)
	return nil
}
