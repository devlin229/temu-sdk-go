package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/edit"
)

// BgGloGoodsAddProperty 调用 bg.glo.goods.add.property，支持开平新增货品上的属性信息。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.add.property/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsAddProperty(ctx context.Context, request *edit.BgGloGoodsAddPropertyRequest) (*edit.BgGloGoodsAddPropertyRespose, error) {
	if request == nil {
		return nil, errors.New("新增货品属性请求不能为空")
	}
	const apiType = "bg.glo.goods.add.property"
	var response edit.BgGloGoodsAddPropertyRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodsEditPicturesSubmit 调用 bg.glo.goods.edit.pictures.submit，修改商品素材。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.edit.pictures.submit/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsEditPicturesSubmit(ctx context.Context, request *edit.BgGloGoodsEditPicturesSubmitRequest) (*edit.BgGloGoodsEditPicturesSubmitRespose, error) {
	if request == nil {
		return nil, errors.New("修改商品素材请求不能为空")
	}
	const apiType = "bg.glo.goods.edit.pictures.submit"
	var response edit.BgGloGoodsEditPicturesSubmitRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodsEditProperty 调用 bg.glo.goods.edit.property，支持开平编辑货品上的属性信息。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.edit.property/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsEditProperty(ctx context.Context, request *edit.BgGloGoodsEditPropertyRequest) (*edit.BgGloGoodsEditPropertyRespose, error) {
	if request == nil {
		return nil, errors.New("编辑货品属性请求不能为空")
	}
	const apiType = "bg.glo.goods.edit.property"
	var response edit.BgGloGoodsEditPropertyRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodsEditSensitiveAttr 调用 bg.glo.goods.edit.sensitive.attr，商家编辑货品敏感品属性。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.edit.sensitive.attr/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsEditSensitiveAttr(ctx context.Context, request *edit.BgGloGoodsEditSensitiveAttrRequest) (*edit.BgGloGoodsEditSensitiveAttrRespose, error) {
	if request == nil {
		return nil, errors.New("编辑货品敏感品属性请求不能为空")
	}
	const apiType = "bg.glo.goods.edit.sensitive.attr"
	var response edit.BgGloGoodsEditSensitiveAttrRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodsEditTaskApply 调用 bg.glo.goods.edit.task.apply，支持开平发起货品修改单用于编辑货品信息。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.edit.task.apply/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsEditTaskApply(ctx context.Context, request *edit.BgGloGoodsEditTaskApplyRequest) (*edit.BgGloGoodsEditTaskApplyRespose, error) {
	if request == nil {
		return nil, errors.New("发起货品修改单请求不能为空")
	}
	const apiType = "bg.glo.goods.edit.task.apply"
	var response edit.BgGloGoodsEditTaskApplyRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodsEditTaskSubmit 调用 bg.glo.goods.edit.task.submit，支持开平提交货品修改单，用于编辑货品信息。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.edit.task.submit/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsEditTaskSubmit(ctx context.Context, request *edit.BgGloGoodsEditTaskSubmitRequest) (*edit.BgGloGoodsEditTaskSubmitRespose, error) {
	if request == nil {
		return nil, errors.New("提交货品修改单请求不能为空")
	}
	const apiType = "bg.glo.goods.edit.task.submit"
	var response edit.BgGloGoodsEditTaskSubmitRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodsUpdate 调用 bg.glo.goods.update，这个接口作为后续的通用货品更新接口，用于“无编辑限制规则的字段”的更新编辑，本次只用于产地编辑。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goods.update/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodsUpdate(ctx context.Context, request *edit.BgGloGoodsUpdateRequest) (*edit.BgGloGoodsUpdateRespose, error) {
	if request == nil {
		return nil, errors.New("货品更新接口请求不能为空")
	}
	const apiType = "bg.glo.goods.update"
	var response edit.BgGloGoodsUpdateRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloGoodslogisticsTemplateEdit 调用 bg.glo.goodslogistics.template.edit，编辑商品运费模板。
// 字段及权限见 docs/Temu_PA_APIs/edit/bg.glo.goodslogistics.template.edit/parameters.json。
// 不自动重试；业务失败保留完整响应，传输或解析失败返回 nil 响应。
func (t *Temu) BgGloGoodslogisticsTemplateEdit(ctx context.Context, request *edit.BgGloGoodslogisticsTemplateEditRequest) (*edit.BgGloGoodslogisticsTemplateEditRespose, error) {
	if request == nil {
		return nil, errors.New("编辑商品运费模板请求不能为空")
	}
	const apiType = "bg.glo.goodslogistics.template.edit"
	var response edit.BgGloGoodslogisticsTemplateEditRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}
