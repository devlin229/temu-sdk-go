# 项目说明

`temu-tool` 是使用 Go 开发的 Temu 接口对接工具，用于辅助商家执行商品管理等日常操作。具体功能以实际接入的 Temu 官方接口和业务需求为准。

当前已实现公共签名与 HTTP 客户端、图片上传、商品发布、类目搜索、父规格、类目属性查询及规格创建 SDK 方法，以及七个接口的响应模型和相应业务请求模型。图片上传已由用户反馈成功；商品发布和类目搜索已通过本地模拟测试，尚未完成真实发品或类目查询联调。商品发布方法不自动查询或校验类目模板和商家模式，调用方应按文档准备业务参数并由平台验证；调用成功不代表已审核通过或已可售。

## 项目目录

```text
temu-sdk-go/
├── AGENTS.md   # 项目介绍与开发约定
├── go.mod      # Go 模块与依赖声明
├── README.md   # 配置、运行与接口扩展说明
├── temu/       # 公共客户端、签名与公共参数/响应
│   ├── common/response.go # 各接口共用的响应状态字段
│   ├── common/params.go # 各接口共用的分页参数
│   ├── image.go # 图片上传 SDK 方法
│   ├── product.go # 商品发布 SDK 方法
│   ├── category.go # 类目搜索、父规格、属性查询及规格创建 SDK 方法
│   ├── category/
│   │   ├── goods_category_match.go # 类目搜索请求及路径响应模型
│   │   ├── goods_attrs_get.go # 类目属性模板查询请求、响应及嵌套模型
│   │   ├── goods_spec_create.go # 创建自定义规格请求及响应模型
│   │   └── goods_parentspec_get.go # 父规格查询响应模型，无业务请求参数
│   ├── product/
│   │   ├── add_goods.go      # 商品发布请求、响应及嵌套模型
│   │   └── goods_list_get.go # 商品列表请求、响应及嵌套模型
│   └── image/
│       └── goods_image_upload_global.go # 图片上传请求、响应模型
├── tests/
│   ├── temu/   # 客户端、图片上传及商品发布模拟测试
│   └── models/ # 参数字典逐层核对、JSON 编解码及可选值测试
├── examples/basic/ # 图片上传示例：读取本地图片、Base64 编码、调用 SDK 并输出 JSON
├── examples/product_add/ # 两款娃娃可选 SKU 的发品测试示例，业务数据为虚构测试值
├── examples/category_match/ # 类目搜索示例，默认搜索“娃娃”并保存响应 JSON
├── examples/attrs_get/ # 类目属性模板查询示例，填写凭据和叶子类目 ID 后保存响应 JSON
├── examples/spec_create/ # 创建规格调用示例，写操作，凭据及业务参数留空
├── examples/parentspec_get/ # 父规格查询示例，凭据字符串直接填写
├── scripts/generate_api_models.py # 参数转换辅助工具，结果输出到指定目录供人工对照
└── docs/
    ├── product-publish.md             # 商品发布使用说明与验证边界
    ├── API-json/                      # 三个接口的请求/响应模板与参数字典
    └── API-docs/                      # Temu 接口接入文档
        ├── Temu_Developer_Guide.md     # 网关、授权、公共参数、签名与通用排错
        ├── Temu_Product_Publish_API.md # 商品发布参数、业务约束、响应与错误码
        ├── Temu_Image_Upload_API.md    # 图片上传 SDK 使用、错误处理与验证边界
        ├── Temu_Category_Match_API.md  # 用户截图整理的类目搜索契约及验证边界
        ├── Temu_Goods_Attrs_Get_API.md # 类目属性模板查询契约及验证边界
        ├── Temu_Spec_Create_API.md # 截图整理的创建规格契约及验证边界
        └── Temu_ParentSpec_Get_API.md  # 用户截图整理的父规格查询契约及验证边界
```

新增代码时按实际需求拆分包，避免提前创建空目录或引入复杂架构。程序入口只负责参数解析、配置加载、依赖初始化和调用业务逻辑。公共协议放在 `temu/`，具体接口按业务域放入其子包；跨接口的商家操作流程在需要时单独组织。目录结构发生变化时同步更新本文件。

## 接口对接约定

- 开发前先阅读 `docs/API-docs/Temu_Developer_Guide.md`，再阅读对应业务接口文档；商品发布的完整嵌套字段见 `Temu_Product_Publish_API.md` 第 8 节。
- 当前文档是官方页面的开发整理版，示例仅用于说明结构，尚未经过真实接口联调。遇到文档冲突、未明确的单位或条件必填规则，应记录并核实，不自行猜测。
- 当前 PA 网关为 `https://openapi-b-partner.temu.com/openapi/router`，商品发布使用 `bg.glo.goods.add`；旧接口 `bg.goods.add` 不与该网关混用。后续接口需单独核对区域、商家模式、网关和授权权限，不统一套用迁移规则。
- 开放接口使用 `app_key`、`app_secret` 和卖家授权的 `access_token`；文档网站的登录 Cookie 不作为开放接口凭据。迁移接口时核对是否需要卖家重新授权。
- 公共参数与业务字段放在同一个 JSON 根对象，不额外包装 `data` 或 `biz_content`。时间戳使用秒级 UNIX 时间，允许误差为前后 300 秒；`version` 无明确要求时按当前文档说明省略。
- 签名排除 `sign`，只将最外层 key 按 ASCII 升序排列，拼接 `key + value`，首尾加 `app_secret`，对 UTF-8 字节计算 MD5 并转成大写十六进制。对象和数组使用紧凑 JSON，内层不重新排序；签名使用的值和序列化结果必须与实际发送内容一致。
- 不将 `app_secret` 放入请求体，不记录签名前含密钥的原文。可选字段需区分省略、`null`、空数组和零值，避免不当使用 `omitempty` 改变业务含义。
- 商品发布需校验类目路径、属性及规格 ID 的匹配关系，以及商家模式、类目和站点决定的条件必填项。示例 ID、图片地址和空属性数组不能直接用于真实发品。
- 单位按具体字段处理：重量为 mg，体积边长为 mm，净含量为实际值乘 1000；半托管使用 `siteSupplierPrices`，其金额单位为分或美分。单字段 `supplierPrice` 的单位仍需按商家模式核实，不从示例推断。金额换算避免浮点精度损失。
- 商品发布按 `success` 判断业务成功，保留 `errorCode`、`errorMsg`；成功后保存 `result` 中的商品、SKC、SKU ID 与本地编码映射。其他接口的返回结构按各自文档解析。
- 当前商品发布文档未给出幂等键，不能假定 `extCode` 自动去重；超时或系统异常后先核对是否已经创建，再决定是否重新提交。
- 新接口文档放在 `docs/API-docs/`，引用通用开发文档，并补充接口适用范围、权限、字段、约束、示例、错误码及来源版本；接口特殊规则不要改成全局默认规则。
- 三个接口的模型都是普通 SDK 源码，可直接维护；字段契约以 `docs/API-json/*.parameters.json` 为准，并由 `tests/models/` 对照字典逐层验证。`scripts/generate_api_models.py` 仅为参数转换辅助工具，必须用 `--output-dir` 指定输出目录，建议输出至临时目录供人工对照。业务校验和调用逻辑维护在独立文件中。
- 类目搜索模型按用户提供的官方页面截图维护，字段及来源记录在 `Temu_Category_Match_API.md`；尤其保留 JSON 名称 `categoryPathDTOS`。截图未提供网关及权限信息，实现暂沿用当前 PA 网关，真实接入时需核验。
- 父规格查询按用户提供的截图维护，响应 JSON 名称保留 `parentSpecDTOS`，模型放在 `temu/category/`；没有业务参数的方法只接收 context，不创建多余的空 Request 类型。

## 代码规范

- 在未经允许的情况下严禁修改任何代码
- 遵循 Go 通用惯例，使用 `gofmt` 格式化代码，优先使用标准库，按需引入第三方依赖。
- 包名使用简短的小写单词；导出标识符使用大驼峰，内部标识符使用小驼峰；文件名使用小写，必要时以下划线分隔。
- 保持函数职责清晰，避免不必要的抽象和重复代码；注释优先使用中文，说明业务规则、约束或实现原因。
- 显式处理错误，补充必要的操作上下文，包装错误时使用 `%w`；不要忽略错误或用 `panic` 处理常规业务失败。
- 网络调用接收 `context.Context`，设置合理的超时，及时关闭响应体，并分别处理 HTTP 错误和 Temu 业务错误。
- 请求与响应优先使用明确的结构体及 `json` 标签；字段类型、必填条件和签名规则以官方文档为准，不猜测接口契约。
- SDK 方法名由接口 `type` 按点分词转为大驼峰，例如 `bg.goods.image.upload.global` 对应 `BgGoodsImageUploadGlobal`；请求、响应入口结构体统一命名为“方法名 + Request”和“方法名 + Respose”（按约定使用 Respose 拼写）。
- 嵌套结构体按字段含义命名，例如图片上传的 `Options`、`Result`，不机械追加完整方法名前缀；仅在同包确实有名称冲突时添加必要的业务限定。
- 非必要不使用指针：普通字符串、数值和数组优先用值类型或切片；只为可选对象、可空响应、有清空语义的数组，以及必须区分未设置和零值的字段保留指针。图片上传基础字段使用值类型，imageBizType、选项开关和 sizeMode 不加 omitempty，保留有效的 0/false。
- 各接口响应匿名嵌入 `common.Response`，复用 `Success`、`ErrorCode`、`ErrorMessage`，并单独定义具体类型的 `Result`。公共状态放在 `temu/common/`，避免业务模型与客户端循环依赖；JSON 字段仍平铺在响应根对象。
- 鉴权、签名和公共请求逻辑集中封装；商家业务逻辑与底层 HTTP 调用分开。
- Cookie、Token、App Secret 等凭据通过环境变量或本地配置提供，不硬编码、不提交到仓库、不写入日志；日志中的敏感字段必须脱敏。
- 本地测试例外：按用户明确要求，示例直接使用 `appKey`、`appSecret`、`accessToken` 等字符串，不读取环境变量或命令行参数；仓库中的凭据示例值保持为空。`examples/product_add/` 的两款娃娃业务参数按用户授权使用虚构测试值，不将这些值描述为真实类目、规格或已核实的商品参数。
- 按用户明确要求，当前命令行调试例外：发送前将请求方法、URL、请求头和已签名请求参数原样输出到 stderr，不脱敏（含 access_token、app_key、sign 及图片签名链接）。app_secret 不属于请求参数，仍不输出；结果 JSON 独立输出到 stdout。
- 对新增、修改、删除商品等写入操作，不进行未经文档确认可安全重试的自动重试，避免重复执行。
- 为签名、参数转换和错误处理等关键逻辑编写测试；测试优先使用模拟服务，避免默认调用真实商家接口。
- 测试文件统一放在 `tests/` 下按模块分目录，测试数据放在 `tests/testdata/`，不散落到生产代码目录。通过公开入口验证行为，不为测试单独公开内部实现函数。
- 提交变更前执行 `go test ./...` 和 `go vet ./...`。若受环境或外部条件限制无法完成，明确说明原因，不将模拟测试描述为真实接口验证。
- 新增接口时，在 `docs/` 中记录文档来源、配置方式、调用示例和已知限制；示例只使用占位凭据。

类目属性查询 `bg.glo.goods.attrs.get` 按同名前缀 JSON 参数字典维护，模型放在 `temu/category/goods_attrs_get.go`；多语言字段使用动态 `map[string]string`，`productCreateTime` 为毫秒时间戳，`chooseAllQualifySpec` 按 BOOLEAN 保留 bool。已通过模型字典与模拟测试，尚未真实查询联调。

创建规格 `bg.glo.goods.spec.create` 按用户截图维护（页面更新 2026-01-11 14:19:11），请求含必填 `parentSpecId`、`specName`，响应为 `result.specId`。暂沿用 PA 网关，网关和权限需真实接入时核验；写操作不自动重试，不假设名称可去重。
