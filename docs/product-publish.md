# 商品发布使用说明

依据项目内 `API-docs/Temu_Developer_Guide.md`、`API-docs/Temu_Product_Publish_API.md`（版本 1.1，2026-11-12）实现 `bg.glo.goods.add`，适用于文档所述 PA 网关及授权体系。使用前核对当前店铺模式及货品 API 权限。

## 输入与命令

参考 `examples/product-add.json` 准备只含业务字段的 JSON。示例中的类目、属性、规格 ID、图片 URL 和价格仅为占位数据，通过基础校验也不代表能通过平台校验。不要直接提交示例。

```sh
go run . -validate-product product.json
go run . -publish-product product.json > product-result.json
```

校验命令不需要凭据，也不发请求。发布命令读取环境变量 `TEMU_APP_KEY`、`TEMU_APP_SECRET`、`TEMU_ACCESS_TOKEN`，网关和超时参考 `.env.example`；不会读取文档登录 Cookie。命令互斥，默认显示帮助。输入限制为 8 MiB；未知字段、错误类型、额外 JSON 对象及公共鉴权字段会被拒绝。

输出为包含 `productId`、`productSkcList`、`productSkuList` 的 JSON；SKU 结果保留 `extCode` 和 `skuSpecList`，便于保存本地编码映射。当前没有数据库，调用方负责保存结果。业务失败输出错误码与错误信息，进程返回非零状态。接口成功不表示已审核通过或可售。

## Go 调用

在本项目模块内复用已初始化的 `*temu.Client`：

```go
req, err := product.DecodeAddRequest(reader)
if err != nil {
    return err
}
result, err := product.NewService(client).Add(ctx, req)
if err != nil {
    var apiErr *product.APIError
    if errors.As(err, &apiErr) {
        // 按 apiErr.Code 和 apiErr.Message 处理业务失败。
    }
    return err
}
// 保存 result.ProductID、result.ProductSKCList 和 result.ProductSKUList。
```

请求模型覆盖字段字典的 57 个对象结构、304 条字段及元素定义，ID、金额等 INTEGER 使用 `int64`，不经浮点数中转。模型由 `scripts/generate_product_models.py` 生成，维护该脚本需要 Python 3，程序运行只需要 Go。

## 校验与序列化边界

- 本地检查标题、类目路径、素材 URL、属性/规格数组、SKC/SKU、币种、预览图 URL、重量、体积和必要敏感属性对象，以及部分价格、库存、物流、产地与合规字段。
- 半托管信息存在时必须提供绑定站点及 SKU 分站点价格，拒绝同时发送已废弃的单字段 `supplierPrice`；站点价格按调用方录入的分/美分原样发送。店铺支持币种、站点范围等由平台判定。
- 可选标量使用指针，显式 `0`、`false`、空字符串不会被自动省略。可选列表一般使用切片指针，nil 指针省略，指向空切片发送 `[]`。多语言图片 `imgUrlList` 按文档保留空列表与 null 的差异：空切片发送 `[]`，nil 切片发送 `null`。
- 必填类目层级即使为零也会发送。允许空编码或空单位的字段不强制非空。可选净含量对象 `{}` 等清空语义会保留。
- 本地校验不完整复刻平台规则：不会查询属性/规格模板或核验 ID 与名称匹配，不检查远程图片像素、大小和平台域名白名单，也不推断服饰类目、定制资格、站点语言或店铺特殊条件。其他嵌套对象的模式相关必填项仍需调用方依文档补齐并由平台确认。
- `supplierPrice` 单字段单位仍按原文保留疑点，不自动换算；体积按 `len >= width >= height > 0` 检查，字段文案冲突仍需真实联调核实。

## 错误与重复提交

HTTP 和业务失败分别保留为 `*temu.HTTPError`、`*product.APIError`，网络及上下文错误通过 `%w` 保留。只按 `success` 判断业务状态，不假定成功错误码固定为零。缺失状态或成功响应缺少有效商品 ID 时返回待核对错误。

所有失败都不会自动重试，`extCode` 不视作幂等键。超时、取消、响应解析失败或结果保存失败可能发生在平台已创建商品之后，必须先核对平台记录。写入成功后审核与销售状态查询不在本次接口范围内。

## 验证记录

本次以模拟 Transport / 响应验证公共参数合并、发布请求、结果 ID 与规格映射、业务错误、异常响应、输入校验、可选值语义和半托管价格；命令行离线校验不依赖凭据。未连接真实 Temu 店铺，未创建真实商品。
