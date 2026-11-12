# 类目搜索 SDK

接口 `bg.glo.goods.category.match` 用于根据类目名称搜索可用于发品挂靠的类目路径。本文依据用户提供的四张官方接口页面截图整理，页面更新时间为 2026-04-29 22:27:22；未提供页面 URL、权限包、适用商家模式或业务错误码清单，不自行补造。

公共参数、授权、签名和通用错误处理见 [Temu_Developer_Guide.md](Temu_Developer_Guide.md)。截图没有显示该接口的完整网关或迁移清单；当前实现暂沿用项目 PA 网关 `https://openapi-b-partner.temu.com/openapi/router` 和 POST JSON，尚未通过真实查询确认该接口的网关及授权适用性。模拟测试只验证实现使用了该地址，不能代替平台核验。

## 调用

```go
response, err := client.BgGloGoodsCategoryMatch(ctx, &category.BgGloGoodsCategoryMatchRequest{
    SearchText: "娃娃",
    // SiteID: 实际站点ID, // 可选
})
if err != nil {
    return err
}
if response.Result != nil {
    for _, path := range response.Result.CategoryPathDTOs {
        // path.Cat1DTO ... path.Cat10DTO 保留各级类目信息。
        // 根据名称、IsLeaf、IsHidden 和 HiddenType 判断合适的路径。
    }
}
```

可直接运行的示例在 [examples/category_match/main.go](../../examples/category_match/main.go)。按既有测试习惯，在源码中填写 appKey、appSecret、accessToken 三个字符串；searchText 默认为“娃娃”，siteID 默认为 0（省略）。不读取环境变量或命令行参数。

```sh
go run ./examples/category_match
```

填写凭据后会真实发起类目查询，完整响应输出到 stdout 并保存到 `category-match-result.json`，请求日志输出到 stderr；不会自动选中某个返回路径，也不会提交商品。默认凭据为空时直接退出，不发送请求。

## 请求参数

| JSON 字段 | Go 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| searchText | string | 否 | Search Text，文档最大长度 200；SDK 不自动截断或改写 |
| siteId | int64 | 否 | Site ID |

请求类型为 `category.BgGloGoodsCategoryMatchRequest`，普通字段使用值类型及 omitempty；空字符串和零值站点省略。两个字段均可选，允许传入空结构体；nil 请求返回本地错误。截图没有说明空查询是否返回全量类目，也没有说明默认站点，不据此推断接口行为。最大长度的计量方式未注明，SDK 暂不自行增加字节或字符计数校验。

## 响应参数

响应类型为 `category.BgGloGoodsCategoryMatchRespose`，匿名嵌入 `common.Response`，根字段为 success、errorCode、errorMsg、result。result 使用 `category.CategoryMatchResult`。类目模型位于 `temu/category/goods_category_match.go`，客户端方法位于 `temu/category.go`。

`result.categoryPathDTOS` 是 `[]CategoryPathDTO`。注意 JSON 字段名按截图使用 `categoryPathDTOS`，不改成 categoryPathDTOs。每条路径包含 `cat1DTO` 至 `cat10DTO`，对应各级 `*CategoryDTO`；缺失或 null 的层级保持 nil，避免将不存在的层级伪造成类目 ID 0 的节点。

截图展开了一级和九级节点，字段如下；其余层级按同一类目节点结构表示。

| JSON 字段 | Go 类型 | 说明 |
| --- | --- | --- |
| catId | int64 | 类目 ID |
| catName | string | 类目名称 |
| parentCatId | int64 | 父类目 ID |
| catType | int64 | 类目类型，截图未给出枚举定义 |
| isLeaf | bool | 是否叶子类目 |
| catLevel | int64 | 类目层级 |
| isHidden | bool | 是否隐藏 |
| hiddenType | int64 | 0 不隐藏，1 正常隐藏，2 旧类目，3 废弃类目 |

SDK 保留平台返回的全部路径和标记，不自行过滤隐藏节点或选择第一条结果。用于发品时，由调用方选择适合商品的有效路径，将该路径各层 `catNDTO.catId` 对应填写到发品请求的 `CatNID`，无对应层级按发品文档传 0。类目路径有效不代表之前虚构的商品属性和规格 ID 也有效；它们仍须匹配实际类目模板。

## 错误与验证

success=false 时返回完整响应和 `*temu.BusinessError`，保留 APIType、Code、Message。HTTP 非 2xx、网络错误和解析失败返回 nil 响应及错误；底层取消和超时经 %w 保留。默认不自动重试，不自行定义截图未提供的业务错误码。

`tests/temu/category_match_test.go` 使用模拟 Transport 验证请求参数及省略行为、接口名、十级路径、多条结果、全部节点字段、大整数 ID、缺失/null 层级、空结果、业务错误、HTTP 错误、非法响应、nil 请求以及取消/超时。公共签名和发送前日志复用现有客户端及其测试。

```sh
go test ./...
go vet ./...
```

本地测试不调用真实接口，尚未验证真实站点、权限及网关是否适用。
