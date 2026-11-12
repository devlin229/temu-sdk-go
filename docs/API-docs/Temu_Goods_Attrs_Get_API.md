# 类目属性模板查询 bg.glo.goods.attrs.get

版本：1.0 ｜ 整理日期：2026-11-12

## 来源与适用范围

字段以仓库内 `docs/API-json/bg.glo.goods.attrs.get.{request,response}.parameters.json` 为契约，对应 `.json` 模板仅展示结构，不是真实接口数据或业务默认值。

字典记录的官方页面为 `https://agentpartner.temu.com/document?cataId=875198836203&docId=934974974297`，页面更新时间 2026-04-29 22:27:22，导出日期 2026-11-12。本次根据本地 JSON 实现，未在线复核来源页面。

接口查询叶子类目的属性模板、候选值和规格限制。字典指定 PA 网关 `https://openapi-b-partner.temu.com/openapi/router`；具体商家模式、区域、站点范围及权限包未明确，真实接入时需核验应用权限和卖家授权。公共授权、秒级时间戳、签名见 [通用开发文档](Temu_Developer_Guide.md)。

## 请求与调用

方法为 `client.BgGloGoodsAttrsGet(ctx, request)`，模型位于 `temu/category/goods_attrs_get.go`；公共参数由客户端添加到 JSON 根对象。

| JSON 字段 | Go 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| catId | int64 | 是 | 叶子类目 ID |
| productCreateTime | int64 | 否 | 商品创建时间，毫秒时间戳；零值省略 |
| supportedType | int64 | 否 | 渠道类型，传入时优先于 mallId；枚举及默认值未明确；零值省略 |
| langList | []string | 否 | 多语言列表，影响性能，一般不传；nil 省略，非 nil 空切片发送 [] |

`mallId` 不在请求字典中，不额外添加。可选数值若需显式传入 0，应先核实其含义与适用条件，再调整字段表示。`version` 按字典说明省略。

```go
// 导入 context、fmt、temu 和 temu/category；凭据及叶子类目 ID 由调用方提供。
func queryAttrs(ctx context.Context, appKey, appSecret, accessToken string, leafCatID int64) error {
    client := temu.NewClient(appKey, appSecret, accessToken)
    response, err := client.BgGloGoodsAttrsGet(ctx, &category.BgGloGoodsAttrsGetRequest{
        CatID: leafCatID,
    })
    if err != nil {
        return fmt.Errorf("查询类目属性模板: %w", err)
    }
    if response.Result != nil {
        fmt.Println(response.Result.TemplateID, len(response.Result.Properties))
    }
    return nil
}
```

凭据通过调用方环境变量或本地配置提供，不提交真实凭据。客户端沿用现有单例配置和请求调试日志行为，见 README。

## 响应与约束

响应为 `category.BgGloGoodsAttrsGetRespose`，匿名嵌入 `common.Response`，保留根层 `success`、`errorCode`、`errorMsg`。`Result` 为 `*category.AttrsResult`，缺失或 null 为 nil。

`result` 包含 `inputMaxSpecNum`、`chooseAllQualifySpec`、`singleSpecValueNum`、`templateId` 和 `properties`。属性候选值、父属性关联、分组、展示条件的完整字段见参数字典和模型注释。

- `lang2Name`、`lang2Value` 使用 `map[string]string`，键为动态语言代码，不能固定为模板中的 `LANGUAGE_CODE`。
- `properties[].required` 是商品属性是否必填的业务字段，区别于字典元数据 `required`。
- `chooseAllQualifySpec` 按字典 BOOLEAN 使用 bool；原说明使用 0/1，不据此改成整数。
- `minValue`、`maxValue` 保留字符串；整数和 ID 使用 int64，避免浮点精度损失。
- `inputMaxNum=0` 表示不允许输入，`valuePrecision=0` 表示不允许小数；调用方应处理父属性关联及展示条件。
- 查询不会自动填写属性、校验商品发布请求或提交商品。

## 错误与验证边界

业务失败同时返回完整响应与 `*temu.BusinessError`，保留接口 type、错误码与错误信息。HTTP、网络或解析失败返回 nil 响应及错误；nil 请求在发送前拒绝，不自动重试。字典未提供接口专属错误码表，不自行补造。

模拟网关覆盖公共与业务参数、毫秒时间戳、大整数、多语言映射、null 结果和错误处理；`tests/models/` 逐层核对参数字典及模板编解码。尚未真实查询联调，模拟通过不代表网关权限和业务参数已获平台认可。

## 本地调用测试示例

在 `examples/attrs_get/main.go` 中直接填写 `appKey`、`appSecret`、`accessToken` 和真实叶子类目 `catID` 后执行：

```sh
go run ./examples/attrs_get
```

可选参数 `productCreateTime` 为毫秒时间戳，`supportedType` 的枚举需先核实，`langList` 默认不传。凭据为空或 catID 未填写时退出，不发送请求。填写完成后运行会真实查询；完整响应保存至 `attrs-get-result.json` 并输出到 stdout，业务失败时也保留响应并返回错误。请求调试日志沿用客户端行为，输出到 stderr；示例不会自动重试。

自动化用例位于 `tests/temu/attrs_get_test.go`，覆盖请求、可选值、空结果、分组与父属性关联、大整数精度和失败处理；模型逐层核对见 `tests/models/models_test.go`。自动化测试不使用示例凭据或调用真实网关。
