# 商品发布 SDK

`(*temu.Temu).BgGloGoodsAdd` 调用 `bg.glo.goods.add`，使用 PA 网关 `https://openapi-b-partner.temu.com/openapi/router`，通过 POST JSON 创建并提交货品信息。接口成功不表示已审核通过或已可售。

接口范围、权限、完整业务规则和错误码以 [Temu_Product_Publish_API.md](API-docs/Temu_Product_Publish_API.md) 为准；公共授权和签名见 [Temu_Developer_Guide.md](API-docs/Temu_Developer_Guide.md)。字段依据 `docs/API-json/bg.glo.goods.add.{request,response}.parameters.json`，来源为官方接口页，页面更新时间 2026-08-25 22:03:37，导出日期 2026-11-12。

## 调用方式

```go
func publish(ctx context.Context, client *temu.Temu, request *product.BgGloGoodsAddRequest) error {
    response, err := client.BgGloGoodsAdd(ctx, request)
    if err != nil {
        var businessErr *temu.BusinessError
        if errors.As(err, &businessErr) {
            // response 同时保留平台完整响应，可读取错误码和错误信息。
            return fmt.Errorf("发品失败（%d）: %w", businessErr.Code, err)
        }
        // 超时或解析失败时，平台可能已经创建商品，先核对记录。
        return err
    }
    if response.Result == nil || response.Result.ProductID == 0 {
        return errors.New("发布响应缺少商品 ID，请先核对平台记录")
    }
    // 调用方在这里保存以下数据，SDK 不自动写文件或数据库。
    // response.Result.ProductID
    // response.Result.ProductSKCList
    // response.Result.ProductSKUList：包含 SKU ID、SKC ID、extCode、skuSpecList。
    return nil
}
```

以上代码片段使用 `context`、`errors`、`fmt`，以及本模块的 `temu` 和 `temu/product` 包。客户端通过 `temu.NewClient(appKey, appSecret, accessToken)` 初始化，仅首次初始化的凭据生效；默认 HTTP 超时为 30 秒，调用方也可通过 context 设置截止时间。

方法签名为：

```go
func (t *Temu) BgGloGoodsAdd(
    ctx context.Context,
    request *product.BgGloGoodsAddRequest,
) (*product.BgGloGoodsAddRespose, error)
```

请求和响应模型位于 `temu/product/add_goods.go`，调用方法位于 `temu/product.go`。响应匿名嵌入 `common.Response`，可直接读取 `Success`、`ErrorCode`、`ErrorMessage`，业务结果使用具体类型的 `Result`。

## 请求准备与序列化

- 只填写业务字段，客户端自动在 JSON 根对象添加 type、app_key、access_token、秒级时间戳、data_type 和 sign。默认不传 version，不发送 app_secret，不包装 data 或 biz_content。
- `docs/API-json/*.request.json` 展示完整结构，包含公共参数和占位值，不是可直接用于发品的请求。全托模板也含占位数据，必须替换类目路径、属性/规格 ID、图片 URL、价格等信息，并删除不使用的可选字段。
- 发品图片使用上传接口返回的平台图片 URL；不传本地路径、Base64 或 Data URL。半托场景按文档使用分站点价格、物流及仓库字段；不在 SDK 内推断模式或自动换算金额。
- INTEGER 使用 int64，重量值按 mg、边长按 mm、净含量按实际值乘 1000 填写。分站点 supplierPrice 单位为分或美分；单字段 supplierPrice 的单位仍需按实际商家模式确认。
- 普通字段使用值类型，普通可选零值由 omitempty 省略；必填字段保留零值。普通可选切片使用 omitzero，nil 省略、非 nil 空切片发送 []。
- 可选对象和有明确零值、null 或清空语义的字段按需保留指针。例如 imgUrlList 的 nil 指针省略，指向 nil 切片发送 null，指向空切片发送 []。

SDK 拒绝 nil 请求，其他业务字段按模型序列化后原样发送，不自动填充业务默认值或调整图片、价格、规格。当前方法没有独立的商品业务校验器：调用方需要按业务文档检查必填和条件必填，类目路径、属性与规格的 ID 匹配、站点及商家模式规则仍需结合实际模板和平台校验；本方法不查询这些依赖接口。

## 响应、错误和重复提交

- success=true：返回完整响应和 nil 错误，不能假定成功时 errorCode 固定为 0。调用方检查 result，并保存商品、SKC、SKU ID 及 extCode 映射。
- success=false：同时返回完整响应和 `*temu.BusinessError`，通过 errors.As 提取 APIType、Code、Message。
- HTTP 非 2xx、网络错误或 JSON 解码失败：返回 nil 响应和错误。取消、超时等底层错误经 %w 保留，可用 errors.Is 判断。

所有情况均不自动重试。接口未声明幂等键，不能假定 extCode 去重；超时、取消、解析失败或结果保存失败后先核对是否已创建，再决定是否重发。

按当前调试要求，公共客户端在发送前将方法、URL、请求头和完整签名参数输出到 stderr，不脱敏；app_secret 不输出。SDK 方法本身不向 stdout 写响应，由调用方选择保存方式。

## 本地验证

```sh
go test ./...
go vet ./...
```

`tests/temu/product_add_test.go` 使用模拟 HTTP Transport 验证 PA 网关和接口名、根参数合并、嵌套 SKU/规格/重量体积/价格、必填零值与空数组、响应 ID 精度、编码及规格映射、业务失败、HTTP 失败、非法响应、nil 请求、传输错误和不重试。签名及发送前日志由公共客户端复用，相关行为另有图片上传测试覆盖。`tests/models/` 对照独立字段字典验证模型。

上述检查没有连接真实 Temu 店铺，未创建真实商品，尚未完成真实发品联调。
