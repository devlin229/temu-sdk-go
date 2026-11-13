package temu

import (
	"context"
	"errors"

	"github.com/devlin229/temu-sdk-go/temu/image"
)

// BgGoodsImageUploadGlobal 调用 bg.goods.image.upload.global，上传已编码为 Base64 的图片。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.goods.image.upload.global/parameters.json。
// 网关与公共签名由客户端处理，请求字段和图片处理选项按原值发送，不自动重试。
// 业务失败时同时返回完整响应与 *BusinessError；传输或解析失败时响应为 nil。
func (t *Temu) BgGoodsImageUploadGlobal(ctx context.Context, request *image.BgGoodsImageUploadGlobalRequest) (*image.BgGoodsImageUploadGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.goods.image.upload.global: request must not be nil")
	}
	const apiType = "bg.goods.image.upload.global"
	var response image.BgGoodsImageUploadGlobalRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{
			APIType: apiType,
			Code:    response.ErrorCode,
			Message: response.ErrorMessage,
		}
	}
	return &response, nil
}

// BgAlgoImageTranslateGlobal 调用 bg.algo.image.translate.global，商品图片翻译。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.algo.image.translate.global/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
func (t *Temu) BgAlgoImageTranslateGlobal(ctx context.Context, request *image.BgAlgoImageTranslateGlobalRequest) (*image.BgAlgoImageTranslateGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.algo.image.translate.global: request must not be nil")
	}
	const apiType = "bg.algo.image.translate.global"
	var response image.BgAlgoImageTranslateGlobalRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgAlgoImageTranslateResultGlobal 调用 bg.algo.image.translate.result.global，商品图片翻译接口查询。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.algo.image.translate.result.global/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
func (t *Temu) BgAlgoImageTranslateResultGlobal(ctx context.Context, request *image.BgAlgoImageTranslateResultGlobalRequest) (*image.BgAlgoImageTranslateResultGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.algo.image.translate.result.global: request must not be nil")
	}
	const apiType = "bg.algo.image.translate.result.global"
	var response image.BgAlgoImageTranslateResultGlobalRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgCompliancepictureGetGlobal 调用 bg.compliancepicture.get.global，批量识别牛皮癣图片。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.compliancepicture.get.global/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
func (t *Temu) BgCompliancepictureGetGlobal(ctx context.Context, request *image.BgCompliancepictureGetGlobalRequest) (*image.BgCompliancepictureGetGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.compliancepicture.get.global: request must not be nil")
	}
	const apiType = "bg.compliancepicture.get.global"
	var response image.BgCompliancepictureGetGlobalRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloColorimageurlGet 调用 bg.glo.colorimageurl.get，色块图获取。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.glo.colorimageurl.get/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
func (t *Temu) BgGloColorimageurlGet(ctx context.Context, request *image.BgGloColorimageurlGetRequest) (*image.BgGloColorimageurlGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.colorimageurl.get: request must not be nil")
	}
	const apiType = "bg.glo.colorimageurl.get"
	var response image.BgGloColorimageurlGetRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGloFancyImageCm2in 调用 bg.glo.fancy.image.cm2in，图片中cm转inch。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.glo.fancy.image.cm2in/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
// 响应仅含 code/errMsg/resultUrl，字典未定义成功码；不推断业务成败，调用方核对响应。
// 返回 nil error 仅表示 HTTP 与模型解析成功，不代表图片转换业务成功。
func (t *Temu) BgGloFancyImageCm2in(ctx context.Context, request *image.BgGloFancyImageCm2inRequest) (*image.BgGloFancyImageCm2inRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.fancy.image.cm2in: request must not be nil")
	}
	const apiType = "bg.glo.fancy.image.cm2in"
	var response image.BgGloFancyImageCm2inRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BgGloPicturecompressionGet 调用 bg.glo.picturecompression.get，高清图片压缩处理。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.glo.picturecompression.get/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
func (t *Temu) BgGloPicturecompressionGet(ctx context.Context, request *image.BgGloPicturecompressionGetRequest) (*image.BgGloPicturecompressionGetRespose, error) {
	if request == nil {
		return nil, errors.New("bg.glo.picturecompression.get: request must not be nil")
	}
	const apiType = "bg.glo.picturecompression.get"
	var response image.BgGloPicturecompressionGetRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}

// BgGoodsTexttopictureAddGlobal 调用 bg.goods.texttopicture.add.global，文字转图片-global。
// 字段及权限见 docs/Temu_PA_APIs/image/bg.goods.texttopicture.add.global/parameters.json。
// 不自动重试或轮询；平台响应原样保留，SDK 不打印请求或鉴权数据。
func (t *Temu) BgGoodsTexttopictureAddGlobal(ctx context.Context, request *image.BgGoodsTexttopictureAddGlobalRequest) (*image.BgGoodsTexttopictureAddGlobalRespose, error) {
	if request == nil {
		return nil, errors.New("bg.goods.texttopicture.add.global: request must not be nil")
	}
	const apiType = "bg.goods.texttopicture.add.global"
	var response image.BgGoodsTexttopictureAddGlobalRespose
	if err := t.requestInto(ctx, gatewayPA, apiType, request, &response); err != nil {
		return nil, err
	}
	if !response.Success {
		return &response, &BusinessError{APIType: apiType, Code: response.ErrorCode, Message: response.ErrorMessage}
	}
	return &response, nil
}
