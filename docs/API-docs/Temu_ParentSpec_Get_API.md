# 查询父规格列表 SDK

接口 `bg.glo.goods.parentspec.get`，用于发布时查询父规格列表。来源为用户提供的两张官方页面截图，页面更新时间为 2026-01-11 14:19:11；用户明确说明没有业务请求参数。截图未提供页面 URL、权限包、适用商家模式和业务错误码清单，不自行补造。

公共授权、时间戳与签名参见 [Temu_Developer_Guide.md](Temu_Developer_Guide.md)。当前实现使用项目 PA 网关 `https://openapi-b-partner.temu.com/openapi/router` 和 POST JSON；截图没有包含该接口完整网关或迁移信息，尚未通过真实父规格查询核实网关和权限适用性。

## 调用

```go
response, err := client.BgGloGoodsParentspecGet(ctx)
if err != nil {
    return err
}
if response.Result != nil {
    for _, item := range response.Result.ParentSpecDTOs {
        // item.ParentSpecName：父规格名称
        // item.ParentSpecID：父规格 ID
    }
}
```

方法只有 context 参数，不要求调用方构造空请求结构体，也没有新增业务 Request 类型。发送时只包含公共参数 type、app_key、access_token、timestamp、data_type、sign，不传类目 ID、站点、空 data 或 biz_content，不发送 app_secret。

SDK 方法在 `temu/category.go`，响应模型在 `temu/category/goods_parentspec_get.go`，父规格查询与类目搜索统一放入 category。

## 响应字段

`category.BgGloGoodsParentspecGetRespose` 匿名嵌入 common.Response，保留 success、errorCode、errorMsg，result 为可空的 `*category.Result`。

| JSON 路径 | Go 类型 | 说明 |
| --- | --- | --- |
| result.parentSpecDTOS | []category.ParentSpecDTO | 父规格列表，DTOS 大小写按截图保留 |
| result.parentSpecDTOS[].parentSpecName | string | 父规格名称 |
| result.parentSpecDTOS[].parentSpecId | int64 | 父规格 ID |

只按 success 判断业务状态，不假定成功错误码固定为 0。业务失败同时返回完整响应和 `*temu.BusinessError`，保留 APIType、Code、Message；HTTP、网络或解析失败返回 nil 响应和错误。取消及超时经 %w 保留。默认不自动重试。

该接口只提供父规格 ID 与名称；截图没有子规格 ID 或类目属性映射，不能仅凭此结果推断发品所需的全部规格数据。SDK 不会自动改写发品 JSON 或提交商品。

## 可运行示例

在 [examples/parentspec_get/main.go](../../examples/parentspec_get/main.go) 中直接填写 appKey、appSecret、accessToken 三个字符串，然后从项目根目录执行：

```sh
go run ./examples/parentspec_get
```

默认凭据为空，运行会在发送前退出。填写后会真实查询父规格，完整响应输出到 stdout 并保存到 `parentspec-result.json`；最终签名请求日志由公共客户端输出到 stderr。示例不读取环境变量或命令行参数。

## 验证边界

`tests/temu/parentspec_get_test.go` 使用模拟 HTTP Transport 验证仅发送公共参数、接口名、列表和名称解析、大整数 ID、输出 JSON 字段拼写、空结果、业务错误、HTTP 失败、非法响应、取消与超时，以及不重试行为。签名和请求日志复用公共客户端及已有测试。

```sh
go test ./...
go vet ./...
```

本次没有执行真实父规格查询，不将模拟结果中的名称或 ID 当作官方可用数据。
