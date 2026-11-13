# temu-sdk-go

使用 Go 标准库对接 Temu 开放接口，提供公共签名、HTTP 客户端和明确类型的业务模型。

| 接口 | SDK 方法 | 当前状态 |
| --- | --- | --- |
| 图片上传 bg.goods.image.upload.global | BgGoodsImageUploadGlobal | 已实现；用户已反馈上传成功，自动化测试使用模拟网关 |
| 商品发布 bg.glo.goods.add | BgGloGoodsAdd | 已实现并通过模拟测试，尚未真实发品联调 |
| 类目搜索 bg.glo.goods.category.match | BgGloGoodsCategoryMatch | 按 PA 字典核对并通过模拟测试，尚未真实查询验证 |
| 父规格查询 bg.glo.goods.parentspec.get | BgGloGoodsParentspecGet | 无业务请求参数，已通过模拟测试，尚未真实查询验证 |
| 类目属性查询 bg.glo.goods.attrs.get | BgGloGoodsAttrsGet | 已通过参数字典和模拟测试，尚未真实查询验证 |
| 创建规格 bg.glo.goods.spec.create | BgGloGoodsSpecCreate | 按 PA 字典核对并通过模拟测试，尚未真实创建验证 |
| 商品列表 bg.glo.goods.list.get | BgGloGoodsListGet | 已按 PA 字典接入并通过模拟测试 |

`docs/Temu_PA_APIs/category` 和 `product` 中的全部 22 个接口已接入，分别提供模型和调用方法；完整清单、来源、权限、字段与示例见 [PA 类目与货品接口](docs/API-docs/PA_Category_Product_APIs.md)。全部调用示例按 `examples/category`、`examples/product`、`examples/image` 分组，目录及命令见 [示例说明](examples/README.md)。示例统一从环境变量或 examples/.env 加载凭据，测试使用模拟网关。

## 初始化与调用

```go
client := temu.NewClient(appKey, appSecret, accessToken)
response, err := client.BgGloGoodsAdd(ctx, request)
```

商品发布的 request 类型为 `*product.BgGloGoodsAddRequest`，返回 `*product.BgGloGoodsAddRespose`。图片上传分别使用 `*image.BgGoodsImageUploadGlobalRequest` 和 `*image.BgGoodsImageUploadGlobalRespose`。请求仅含业务字段，公共层在同一 JSON 根对象添加接口 type、授权、时间戳及签名。

客户端沿用单例行为，仅第一次 NewClient 的凭据生效，默认 HTTP 超时为 30 秒。网络调用接收 context，不自动重试。业务失败同时返回完整响应和 `*temu.BusinessError`；成功不代表商品已审核或可售，调用方保存返回的 ID 与本地编码映射。

- [商品发布调用与限制](docs/product-publish.md)
- [类目搜索 SDK 与查询示例](docs/API-docs/Temu_Category_Match_API.md)：配置环境变量或 examples/.env 后运行 `go run ./examples/category/bg_glo_goods_category_match`。
- [父规格查询 SDK 与示例](docs/API-docs/Temu_ParentSpec_Get_API.md)：配置环境变量或 examples/.env 后运行 `go run ./examples/category/bg_glo_goods_parentspec_get`。
- [双款娃娃发品测试示例](docs/product-dolls-test.md)：`go run ./examples/product/bg_glo_goods_add`，优先读取已有请求 JSON，缺失时才生成；缺少凭据时不发送请求。
- [图片上传调用与示例](docs/API-docs/Temu_Image_Upload_API.md)
- [类目属性模板查询 SDK 与调用示例](docs/API-docs/Temu_Goods_Attrs_Get_API.md)：配置凭据并填写叶子类目 ID 后运行 `go run ./examples/category/bg_glo_goods_attrs_get`。
- [创建规格 SDK 与测试示例](docs/API-docs/Temu_Spec_Create_API.md)：配置凭据并填写父规格 ID 和名称后运行 `go run ./examples/category/bg_glo_goods_spec_create`，会真实创建规格。
- [通用开发规则](docs/API-docs/Temu_Developer_Guide.md)
- [商品发布完整接口文档](docs/API-docs/Temu_Product_Publish_API.md)

## 图片上传示例

先配置 `TEMU_APP_KEY`、`TEMU_APP_SECRET`、`TEMU_ACCESS_TOKEN` 环境变量，或复制 `examples/.env.example` 为 `examples/.env` 后填写。然后在 `examples/image/bg_goods_image_upload_global/main.go` 中填写 imagePath，从项目根目录执行：

```sh
go run ./examples/image/bg_goods_image_upload_global
```

示例读取本地图片，构造包含正确 MIME 前缀的 Data URL，并在上传前保存到 image-base64.txt。执行会真实上传；返回 JSON 输出到 stdout。当前请求调试日志包含完整签名参数并输出到 stderr，app_secret 不输出。可以分别保存结果和日志：

```sh
go run ./examples/image/bg_goods_image_upload_global > upload-result.json 2> request-debug.log
```

## 目录与扩展

```text
temu/
├── client.go、sign.go    # HTTP、公共参数与签名
├── category.go          # PA 类目域方法
├── product.go           # PA 货品域方法
├── image.go             # 图片上传方法
├── common/              # 共用响应状态和分页参数
├── category/            # 类目域请求、响应及嵌套模型
├── product/             # 货品域模型，商品发布文件为 goods_add.go
└── image/               # 图片上传模型
examples/
├── category/            # 11 个类目域调用示例
├── product/             # 11 个货品域调用示例
├── image/               # 图片上传调用示例
├── config/              # 公共凭据加载
├── .env.example         # 环境变量与 .env 空模板
└── README.md            # 完整示例清单与运行方式
tests/                   # 模拟接口、契约和兼容性测试
docs/API-docs/           # 通用及业务接口文档
docs/Temu_PA_APIs/       # PA 参数字典和结构模板
scripts/                 # 参数转换辅助工具
```

新接口按官方文档核对适用范围、权限和网关；SDK 方法放在 temu 包，业务模型按领域放在子包，并通过 common 共享公共类型，避免反向引用客户端。方法名由接口 type 转成大驼峰，请求/响应入口结构体采用“方法名 + Request/Respose”；嵌套结构体按字段含义命名。普通字段优先用值类型，必要的可选或清空语义才使用指针。

模型是可直接维护的普通源码。参数转换脚本仅用于输出至指定目录供人工对照，不是 SDK 构建或运行依赖；原始资料说明见 docs/Temu_PA_APIs/README.txt。

## 验证

接口正常返回但 `success=false` 时，SDK 同时返回具体接口响应和 `*temu.BusinessError`，调用方应保留响应中的平台错误码、信息和逐项结果。已收到 HTTP 响应但状态非 2xx、正文读取失败或模型解析失败时，返回 `*temu.ResponseError`，包含 `APIType`、`StatusCode` 和未经改写的 `Body`；读取失败时 Body 为已收到的部分。直接打印错误会包含接口正文。请求编码等本地错误、未收到响应的网络错误仍返回普通 error。

```go
response, err := client.BgGloGoodsCatsGet(ctx, request)
if response != nil {
    // 包括 success=false 的平台业务响应。
    // 按具体响应模型处理 response。
}
var responseErr *temu.ResponseError
if errors.As(err, &responseErr) {
    fmt.Println(responseErr.StatusCode)
    fmt.Println(string(responseErr.Body)) // JSON、文本或 HTML 均保留原始正文
}
```

示例需导入标准库 errors、fmt。ResponseError 支持 Unwrap，可通过 errors.Is/As 检查原始读取或解析原因；平台 HTTP 错误不被补造成业务成功响应，所有错误路径均不自动重试。

使用 go.mod 声明的 Go 版本：

```sh
go test ./...
go vet ./...
go build ./...
```

测试集中位于 tests/，默认使用模拟 HTTP Transport，不读取示例中的凭据，不调用真实商家接口。

旧 SDK 已统一来源注释、结果类型名称和文件名；原有 `category.Result`、`image.Result`、`product.Result`、`product.BgGloGoodsListGetResult` 保留为类型别名。商品列表的可选 ID 数组统一采用 omitzero：nil 省略，显式空切片发送 []。
