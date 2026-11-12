# temu-sdk-go

使用 Go 标准库对接 Temu 开放接口，提供公共签名、HTTP 客户端和明确类型的业务模型。

| 接口 | SDK 方法 | 当前状态 |
| --- | --- | --- |
| 图片上传 bg.goods.image.upload.global | BgGoodsImageUploadGlobal | 已实现；用户已反馈上传成功，自动化测试使用模拟网关 |
| 商品发布 bg.glo.goods.add | BgGloGoodsAdd | 已实现并通过模拟测试，尚未真实发品联调 |
| 类目搜索 bg.glo.goods.category.match | BgGloGoodsCategoryMatch | 按截图实现并通过模拟测试，尚未真实查询验证 |
| 父规格查询 bg.glo.goods.parentspec.get | BgGloGoodsParentspecGet | 无业务请求参数，已通过模拟测试，尚未真实查询验证 |
| 类目属性查询 bg.glo.goods.attrs.get | BgGloGoodsAttrsGet | 已通过参数字典和模拟测试，尚未真实查询验证 |
| 创建规格 bg.glo.goods.spec.create | BgGloGoodsSpecCreate | 按截图实现并通过模拟测试，尚未真实创建验证 |
| 商品列表 bg.glo.goods.list.get | 暂未提供调用方法 | 请求和响应模型已齐全 |

## 初始化与调用

```go
client := temu.NewClient(appKey, appSecret, accessToken)
response, err := client.BgGloGoodsAdd(ctx, request)
```

商品发布的 request 类型为 `*product.BgGloGoodsAddRequest`，返回 `*product.BgGloGoodsAddRespose`。图片上传分别使用 `*image.BgGoodsImageUploadGlobalRequest` 和 `*image.BgGoodsImageUploadGlobalRespose`。请求仅含业务字段，公共层在同一 JSON 根对象添加接口 type、授权、时间戳及签名。

客户端沿用单例行为，仅第一次 NewClient 的凭据生效，默认 HTTP 超时为 30 秒。网络调用接收 context，不自动重试。业务失败同时返回完整响应和 `*temu.BusinessError`；成功不代表商品已审核或可售，调用方保存返回的 ID 与本地编码映射。

- [商品发布调用与限制](docs/product-publish.md)
- [类目搜索 SDK 与查询示例](docs/API-docs/Temu_Category_Match_API.md)：填写源码中的凭据后运行 `go run ./examples/category_match`。
- [父规格查询 SDK 与示例](docs/API-docs/Temu_ParentSpec_Get_API.md)：填写源码中的凭据后运行 `go run ./examples/parentspec_get`。
- [双款娃娃发品测试示例](docs/product-dolls-test.md)：`go run ./examples/product_add`，优先读取已有请求 JSON，缺失时才生成；凭据留空不发送请求。
- [图片上传调用与示例](docs/API-docs/Temu_Image_Upload_API.md)
- [类目属性模板查询 SDK 与调用示例](docs/API-docs/Temu_Goods_Attrs_Get_API.md)：填写凭据和叶子类目 ID 后运行 `go run ./examples/attrs_get`。
- [创建规格 SDK 与测试示例](docs/API-docs/Temu_Spec_Create_API.md)：填写凭据、父规格 ID 和名称后运行 `go run ./examples/spec_create`，会真实创建规格。
- [通用开发规则](docs/API-docs/Temu_Developer_Guide.md)
- [商品发布完整接口文档](docs/API-docs/Temu_Product_Publish_API.md)

## 图片上传示例

在 `examples/basic/main.go` 中填写 appKey、appSecret、accessToken 和 imagePath 四个字符串后，从项目根目录执行：

```sh
go run ./examples/basic
```

示例读取本地图片，构造包含正确 MIME 前缀的 Data URL，并在上传前保存到 image-base64.txt。执行会真实上传；返回 JSON 输出到 stdout。当前请求调试日志包含完整签名参数并输出到 stderr，app_secret 不输出。可以分别保存结果和日志：

```sh
go run ./examples/basic > upload-result.json 2> request-debug.log
```

## 目录与扩展

```text
temu/
├── client.go、sign.go    # HTTP、公共参数与签名
├── image.go             # 图片上传方法
├── product.go           # 商品发布方法
├── category.go          # 类目搜索、父规格、属性查询及规格创建方法
├── common/              # 共用响应状态和分页参数
├── image/               # 图片上传请求/响应模型
├── product/             # 商品发布、商品列表模型
└── category/            # 类目搜索、父规格、属性模板及规格创建模型
examples/basic/          # 本地图片上传示例
examples/category_match/ # 类目搜索示例
examples/parentspec_get/ # 父规格查询示例
examples/attrs_get/      # 类目属性模板查询示例
examples/product_add/    # 商品发布示例
examples/spec_create/    # 创建规格示例（写操作）
tests/                   # 模拟接口与模型测试
docs/API-docs/           # 通用及业务接口文档
docs/API-json/           # 参数字典和占位结构模板
scripts/                 # 参数转换辅助工具
```

新接口按官方文档核对适用范围、权限和网关；SDK 方法放在 temu 包，业务模型按领域放在子包，并通过 common 共享公共类型，避免反向引用客户端。方法名由接口 type 转成大驼峰，请求/响应入口结构体采用“方法名 + Request/Respose”；嵌套结构体按字段含义命名。普通字段优先用值类型，必要的可选或清空语义才使用指针。

模型是可直接维护的普通源码。参数转换脚本仅用于输出至指定目录供人工对照，不是 SDK 构建或运行依赖；使用方式见 docs/API-json/README.txt。

## 验证

使用 go.mod 声明的 Go 版本：

```sh
go test ./...
go vet ./...
go build ./...
```

测试集中位于 tests/，默认使用模拟 HTTP Transport，不读取示例中的凭据，不调用真实商家接口。
