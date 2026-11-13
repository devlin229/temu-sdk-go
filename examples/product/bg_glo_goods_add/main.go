package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/devlin229/temu-sdk-go/temu"
	"github.com/devlin229/temu-sdk-go/temu/product"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	// 直接填写字符串；填写后运行此示例将真实调用商品发布接口。
	appKey := "47bb4bb7769e12d9f7aa93cf029fe529"
	appSecret := "ac0a3e952eaaa5b19c0e615c2ef497f50afa6e49"
	accessToken := "ayzqiujmpkuyst6jpwschlfc6kgzaxhhd9aghi2hdujvia1grfkwvrsv"

	request, err := loadRequest("examples/product/bg_glo_goods_add/product-request.json")
	if err != nil {
		return err
	}
	client := temu.NewClient(appKey, appSecret, accessToken)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, publishErr := client.BgGloGoodsAdd(ctx, request)
	if response != nil {
		raw, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return errors.Join(publishErr, fmt.Errorf("编码发品响应: %w", err))
		}
		// 即使业务失败也保存完整响应，成功时其中包含商品及 SKC/SKU 编码映射。
		if err = os.WriteFile("examples/product/bg_glo_goods_add/product-result.json", append(raw, '\n'), 0o600); err != nil {
			return errors.Join(publishErr, fmt.Errorf("保存发品响应失败，请先核对平台记录: %w", err))
		}
		if _, err = fmt.Fprintln(os.Stdout, string(raw)); err != nil {
			return errors.Join(publishErr, fmt.Errorf("输出发品响应: %w", err))
		}
	}
	if publishErr != nil {
		return fmt.Errorf("测试发品失败，不自动重试: %w", publishErr)
	}
	return nil
}

// loadRequest 优先使用用户编辑的请求，仅在文件不存在时生成模板。
func loadRequest(path string) (*product.BgGloGoodsAddRequest, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		request := makeRequest()
		raw, err := json.MarshalIndent(request, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("编码测试商品: %w", err)
		}
		// 只创建新文件，避免覆盖读取后恰好由用户创建的文件。
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
		if err != nil {
			return nil, fmt.Errorf("创建测试商品请求: %w", err)
		}
		_, writeErr := file.Write(append(raw, '\n'))
		if err := errors.Join(writeErr, file.Close()); err != nil {
			return nil, fmt.Errorf("保存测试商品请求: %w", err)
		}
		return request, nil
	}
	if err != nil {
		return nil, fmt.Errorf("读取商品请求: %w", err)
	}
	var request *product.BgGloGoodsAddRequest
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		return nil, fmt.Errorf("解析商品请求: %w", err)
	}
	if request == nil {
		return nil, errors.New("商品请求必须是 JSON 对象，不能为 null")
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, errors.New("商品请求文件只能包含一个 JSON 对象")
	}
	return request, nil
}

// 类目路径及颜色父规格取自查询结果；子规格和属性 ID、价格、尺寸重量及产地仍为测试数据。
// 本例按全托结构使用 supplierPrice；类目是否适合实际商品、金额单位和其他平台规则仍需核验。
func makeRequest() *product.BgGloGoodsAddRequest {
	const imageURL = "https://img.kwcdn.com/product/open/51b943f3f6014cf6a2e002c0078421dc-goods.jpeg"
	const parentSpecID int64 = 1001 // parentspec-result.json：颜色
	const parentSpecName = "颜色"
	request := &product.BgGloGoodsAddRequest{
		ProductName: "测试商品 娃娃 灰蓝上衣款/酒红上衣款 两款可选",
		ProductI18nReqs: []product.ProductI18nRequest{{
			Language: "en", ProductName: "Test Doll Listing - Gray-Blue or Burgundy Outfit",
		}},
		Cat1ID: 25439, Cat2ID: 25567, Cat3ID: 25589, // 玩具与游戏 → 娃娃及配件 → 娃娃
		MaterialImgURL: imageURL, CarouselImageURLs: []string{imageURL},
		ProductPropertyReqs: []product.ProductPropertyRequest{{
			PID: 40001, TemplatePID: 40002, RefPID: 40003, VID: 40004,
			PropName: "测试材质", PropValue: "虚构测试材质，非实物参数",
		}},
		ProductWhExtAttrReq: &product.ProductWhExtAttrRequest{
			ProductOrigin: product.ProductOrigin{Region1ShortName: "CN"},
		},
	}
	variants := []struct {
		code   string
		name   string
		specID int64
		vid    int64
		price  int64
		weight int64
	}{
		{"GRAY-BLUE", "灰蓝上衣款", 20001, 50011, 7780, 180000},
		{"BURGUNDY", "酒红上衣款", 20002, 50012, 7980, 190000},
	}
	for _, variant := range variants {
		request.ProductSpecPropertyReqs = append(request.ProductSpecPropertyReqs, product.ProductSpecPropertyRequest{
			PID: 50001, TemplatePID: 50002, RefPID: 50003, VID: variant.vid,
			SpecID: variant.specID, ParentSpecID: parentSpecID,
			PropName: parentSpecName, PropValue: variant.name,
		})
		nonSensitive := int64(0)
		// 每个颜色款作为一个 SKC，其下一个 SKU；不是两只一起出售的套装。
		request.ProductSKCReqs = append(request.ProductSKCReqs, product.ProductSKCRequest{
			ExtCode:        "TEST-DOLL-SKC-" + variant.code,
			PreviewImgURLs: []string{imageURL},
			// 按用户本次调试要求：保留一条主销规格对象，四个字段均为默认值。
			MainProductSKUSpecReqs: []product.MainProductSKUSpecRequest{{}},
			ProductSKUReqs: []product.ProductSKURequest{{
				ExtCode:      "TEST-DOLL-SKU-" + variant.code,
				CurrencyType: "CNY", SupplierPrice: variant.price,
				ThumbURL: imageURL,
				ProductSKUSpecReqs: []product.ProductSKUSpecRequest{{
					SpecID: variant.specID, ParentSpecID: parentSpecID,
					SpecName: variant.name, ParentSpecName: parentSpecName,
				}},
				ProductSKUWhExtAttrReq: product.ProductSKUWhExtAttrRequest{
					ProductSKUWeightReq:         product.ProductSKUWeightRequest{Value: variant.weight},
					ProductSKUVolumeReq:         product.ProductSKUVolumeRequest{Len: 200, Width: 120, Height: 100},
					ProductSKUSensitiveLimitReq: product.ProductSKUSensitiveLimitRequest{},
					ProductSKUSensitiveAttrReq:  product.ProductSKUSensitiveAttrRequest{IsSensitive: &nonSensitive},
				},
			}},
		})
	}
	return request
}
