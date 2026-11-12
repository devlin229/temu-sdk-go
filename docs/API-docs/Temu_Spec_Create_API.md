# 创建规格 bg.glo.goods.spec.create

版本：1.0 ｜ 整理日期：2026-11-12

## 来源、用途与适用范围

来源为用户提供的两张官方页面截图：接口简介、请求及返回参数说明。页面更新时间为 **2026-01-11 14:19:11**。截图未提供官方页面 URL、网关、HTTP 方法、商家模式、区域、站点范围和权限包信息，尚未在线核验。

接口用于货品发布时创建自定义规格。实现沿用项目 POST JSON 和当前 PA 网关 `https://openapi-b-partner.temu.com/openapi/router`；真实接入前核实其适用范围和卖家授权权限。授权、公共参数与签名参考 [通用开发文档](Temu_Developer_Guide.md)，不将文档网站 Cookie 作为开放接口凭据。

## 字段契约

业务字段与公共参数放在同一个 JSON 根对象。

| 请求字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| parentSpecId | INTEGER / int64 | 是 | 父规格 ID |
| specName | STRING / string | 是 | 子规格名称 |

请求模型为 `category.BgGloGoodsSpecCreateRequest`，必填字段无 `omitempty`。截图未说明 ID 范围、名称长度、重复名称处理或类目限制，不自行推断；业务有效性由平台验证。

| 响应路径 | 类型 | 说明 |
| --- | --- | --- |
| result | OBJECT / *SpecCreateResult | 结果对象，缺失或 null 时为 nil |
| result.specId | INTEGER / int64 | 子规格 ID |
| success | BOOLEAN / bool | 业务状态 |
| errorCode | INTEGER / int64 | 错误码 |
| errorMsg | STRING / string | 错误信息 |

响应模型为 `category.BgGloGoodsSpecCreateRespose`，匿名嵌入 `common.Response`，公共状态字段保留在根层，整数 ID 不经浮点转换。

结构示例仅展示格式，不代表真实数据或默认值：

```json
{"parentSpecId":0,"specName":""}
```

```json
{"success":true,"errorCode":0,"errorMsg":"","result":{"specId":0}}
```

## SDK 与本地调用测试

```go
response, err := client.BgGloGoodsSpecCreate(ctx, &category.BgGloGoodsSpecCreateRequest{
    ParentSpecID: parentSpecID,
    SpecName: specName,
})
```

调用方应处理 err，并在成功且 Result 非 nil 时保存 `response.Result.SpecID`，供商品规格参数使用。创建规格成功不代表已完成发品。

本地示例为 `examples/spec_create/main.go`。直接填写源码中的 `appKey`、`appSecret`、`accessToken`、`parentSpecID`、`specName` 后执行：

```sh
go run ./examples/spec_create
```

仓库示例凭据和业务参数留空，未填写时不发送请求。填写后运行会真实创建规格；响应保存到 `spec-create-result.json` 并输出到 stdout，业务失败也保留响应。请求调试日志沿用客户端行为输出到 stderr，app_secret 不输出。

## 错误、重试与验证边界

nil 请求在发送前拒绝；业务失败同时返回完整响应及 `*temu.BusinessError`；HTTP、网络或解析失败返回 nil 响应及错误。截图未提供专属错误码表，保留平台 errorCode、errorMsg，不补造枚举。

该接口是写操作，不自动重试。截图未声明幂等键或重复名称可去重；超时或系统异常后需核对是否已经创建，不能据名称假设安全重试。

自动化测试使用模拟网关，覆盖根层请求、公共参数、中文名称、context、大整数 ID、空结果、业务及传输错误、nil 请求和不重试；模型测试检查 JSON 名称、必填零值与公共字段平铺。尚未真实创建规格联调，网关和权限适用性仍待核实。
