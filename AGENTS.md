# 项目说明

`temu-tool` 是使用 Go 开发的 Temu 接口对接工具，用于辅助商家执行商品管理等日常操作。具体功能以实际接入的 Temu 官方接口和业务需求为准。

当前已实现公共签名与 HTTP 客户端、图片上传、JIT 操作、edit 目录全部 8 个编辑接口，以及 docs/Temu_PA_APIs/category 和 product 中全部 22 个接口的 SDK 方法与明确模型。图片上传已由用户反馈成功；PA 类目及货品接口通过本地契约与模拟测试；用户已反馈单款娃娃商品发布成功，其余接口不因此视为完成真实联调。商品发布方法不自动查询或校验类目模板和商家模式，成功不代表审核通过或可售。

## 项目目录

```text
temu-sdk-go/
├── AGENTS.md   # 项目介绍与开发约定
├── go.mod      # Go 模块与依赖声明
├── README.md   # 配置、运行与接口扩展说明
├── temu/       # 公共客户端、签名与公共参数/响应
│   ├── common/response.go # 各接口共用的响应状态字段
│   ├── common/params.go # 各接口共用的分页参数
│   ├── edit.go # PA 编辑域 8 个 SDK 方法
│   ├── edit/ # 编辑请求/响应与共享嵌套模型，保留素材 null 和清空语义
│   ├── jit.go # JIT 操作 SDK 方法
│   ├── jit/jitmode_activate.go # JIT 操作请求及响应模型
│   ├── image.go # 图片上传 SDK 方法
│   ├── product.go # PA 货品域 11 个 SDK 方法
│   ├── category.go # PA 类目域 11 个 SDK 方法
│   ├── category/
│   │   ├── goods_category_match.go # 类目搜索请求及路径响应模型
│   │   ├── goods_attrs_get.go # 类目属性模板查询请求、响应及嵌套模型
│   │   ├── goods_spec_create.go # 创建自定义规格请求及响应模型
│   │   ├── goods_parentspec_get.go # 父规格查询响应模型，无业务请求参数
│   │   └── 其他模型文件 # 子类目、必填信息、包装清单、OEM、属性映射及类目纠正
│   ├── product/
│   │   ├── goods_add.go      # 商品发布请求、响应及嵌套模型
│   │   ├── goods_list_get.go # 商品列表请求、响应及嵌套模型
│   │   └── 其他模型文件 # 品牌、详情、搬运、仓库、招标单、废弃/售罄、运费及生命周期
│   └── image/
│       └── goods_image_upload_global.go # 图片上传请求、响应模型
├── tests/
│   ├── temu/   # 客户端、图片上传及商品发布模拟测试
│   └── models/ # 参数字典逐层核对、JSON 编解码及可选值测试
├── examples/
│   ├── category/ # 11 个类目域示例，子目录为完整接口名替换点为下划线
│   ├── product/  # 11 个货品域示例，保留商品发布数据准备流程
│   ├── edit/     # 8 个编辑域示例，读取同目录 request.json
│   ├── image/    # 图片上传示例
│   ├── jit/      # JIT 操作示例，凭据统一加载并读取 request.json
│   ├── config/   # 公共凭据加载：环境变量优先、examples/.env 回退
│   ├── go.mod    # 独立示例模块，主 SDK 下载时排除整个 examples
│   ├── tests/config/ # 示例公共配置与请求读取测试
│   ├── .env.example # 可提交的空配置模板
│   └── README.md # 示例索引与执行方式
├── scripts/generate_api_models.py # 参数转换辅助工具，结果输出到指定目录供人工对照
└── docs/
    ├── product-publish.md             # 商品发布使用说明与验证边界
    ├── Temu_PA_APIs/                  # PA 请求/响应模板与参数字典
    └── API-docs/                      # Temu 接口接入文档
        ├── PA_Category_Product_APIs.md # 22 个接口的接入清单与验证边界
        ├── Temu_Edit_APIs.md # 8 个编辑接口索引与调用边界
        ├── pa/edit/、pa/jit/ # 编辑及 JIT 逐接口来源、权限、字段与示例
        ├── pa/category/、pa/product/   # 逐接口的来源、权限、字段、错误码及示例
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
- 接口模型都是普通 SDK 源码，可直接维护；字段契约以 `docs/Temu_PA_APIs/<领域>/<type>/parameters.json` 为准，并由 `tests/models/` 对照字典逐层验证。`scripts/generate_api_models.py` 仅为参数转换辅助工具，必须用 `--output-dir` 指定输出目录，建议输出至临时目录供人工对照。业务校验和调用逻辑维护在独立文件中。
- 类目搜索最初按截图维护，现以 PA 参数字典为契约；保留 JSON 名称 `categoryPathDTOS`，网关与权限见对应字典，真实接入仍需验证授权。
- 父规格查询按用户提供的截图维护，响应 JSON 名称保留 `parentSpecDTOS`，模型放在 `temu/category/`；没有业务参数的方法只接收 context，不创建多余的空 Request 类型。

- 编辑域方法使用 `temu/edit/` 模型，`goods.update` 当前仅用于产地修改；修改单申请返回逐项失败信息和 UID 映射，不自动串联提交或猜测 taskVersion。
- 素材接口视频/多语言图片列表使用切片指针区分省略、null（不修改）、[]（删除）；floorId 使用双层指针保留 null 新建楼层语义。JSON 解码须保留这些状态，不能因省略策略丢失清空/新建语义。

## 代码规范

- 示例生成的请求参数、响应结果及其他调试 JSON 文件，统一放在对应接口的示例目录 `examples/<业务域>/<接口名以下划线分词>/`，与该接口的 `main.go` 同目录；后续新增示例也必须遵循。读取已有请求 JSON 时使用该目录，不回退到项目根目录。示例默认从项目根目录运行，代码中的文件路径应明确包含对应示例目录。生成文件需加入 `.gitignore`，不覆盖用户已有请求参数。

- 在未经允许的情况下严禁修改任何代码
- 遵循 Go 通用惯例，使用 `gofmt` 格式化代码，优先使用标准库，按需引入第三方依赖。
- 包名使用简短的小写单词；导出标识符使用大驼峰，内部标识符使用小驼峰；文件名使用小写，必要时以下划线分隔。
- 保持函数职责清晰，避免不必要的抽象和重复代码；注释优先使用中文，说明业务规则、约束或实现原因。
- 显式处理错误，补充必要的操作上下文，包装错误时使用 `%w`；不要忽略错误或用 `panic` 处理常规业务失败。
- 网络调用接收 `context.Context`，设置合理的超时，及时关闭响应体，并分别处理 HTTP 错误和 Temu 业务错误。
- 已收到平台响应但 HTTP 状态异常、正文读取或模型解析失败时，使用 `ResponseError` 保留接口 type、状态码和原始正文；读取中断也保留已收到的部分，底层原因通过 Unwrap 暴露。不得只保留自定义错误文本而丢弃平台正文。正常业务失败继续返回接口响应与 BusinessError，不补造响应，不自动重试。
- 请求与响应优先使用明确的结构体及 `json` 标签；字段类型、必填条件和签名规则以官方文档为准，不猜测接口契约。
- SDK 方法名由接口 `type` 按点分词转为大驼峰，例如 `bg.goods.image.upload.global` 对应 `BgGoodsImageUploadGlobal`；请求、响应入口结构体统一命名为“方法名 + Request”和“方法名 + Respose”（按约定使用 Respose 拼写）。
- 嵌套结构体按字段含义命名，例如图片上传的 `Options`、`Result`，不机械追加完整方法名前缀；仅在同包确实有名称冲突时添加必要的业务限定。
- 可选整数按零值语义选择类型，不因“可选”一律改指针。分页 page/pageSize、正数 ID、时间参数及明确从 1 开始的枚举使用值类型和 omitempty，0 表示未填写；packageType、imageBizType、sizeMode、库存、开关或状态等需明确传 0 的字段使用指针，nil 省略、指向 0 则发送 0。零值含义尚未明确的字段保守保留指针并记录待核实，不推断 0 一定合法。必填整数、响应字段和整数切片元素按契约维护。普通字符串和数组优先用值类型或切片；可选对象、可空响应和清空语义按需使用指针。
- 各接口响应匿名嵌入 `common.Response`，复用 `Success`、`ErrorCode`、`ErrorMessage`，并单独定义具体类型的 `Result`。公共状态放在 `temu/common/`，避免业务模型与客户端循环依赖；JSON 字段仍平铺在响应根对象。
- 鉴权、签名和公共请求逻辑集中封装；商家业务逻辑与底层 HTTP 调用分开。
- 共用参数类型统一维护在 `temu/common/params.go`，业务模型直接引用 common 包，不在 temu 根包保留参数类型别名或重复的 params.go。
- Cookie、Token、App Secret 等凭据通过环境变量或本地配置提供，不硬编码、不提交到仓库、不写入日志；日志中的敏感字段必须脱敏。
- 示例凭据统一调用 `examples/config.LoadCredentials()`：优先读取 `TEMU_APP_KEY`、`TEMU_APP_SECRET`、`TEMU_ACCESS_TOKEN` 系统环境变量，缺失字段从 `examples/.env` 补齐，仍缺失则报错。不得在各示例硬编码或重复读取凭据；`.env` 不提交，`.env.example` 只保留空占位。`examples/product/bg_glo_goods_add/` 的商品发布模板按用户授权复制成功上架的单款 product-request.json，缺文件时由嵌入的 request.example.json 生成。该成功案例只代表用户反馈的特定店铺与参数，不推断为任意商家通用有效。
- 按用户明确要求，当前命令行调试例外：发送前将请求方法、URL、请求头和已签名请求参数原样输出到 stderr，不脱敏（含 access_token、app_key、sign 及图片签名链接）。app_secret 不属于请求参数，仍不输出；结果 JSON 独立输出到 stdout。
- 对新增、修改、删除商品等写入操作，不进行未经文档确认可安全重试的自动重试，避免重复执行。
- 为签名、参数转换和错误处理等关键逻辑编写测试；测试优先使用模拟服务，避免默认调用真实商家接口。
- 各模块测试统一放在该模块的 `tests/` 下：SDK 为根目录 `tests/`，独立示例模块为 `examples/tests/`；测试数据放在各模块 tests/testdata/，不散落到生产代码目录。通过公开入口验证行为，不为测试单独公开内部实现函数。
- 提交变更前执行 `go test ./...` 和 `go vet ./...`。若受环境或外部条件限制无法完成，明确说明原因，不将模拟测试描述为真实接口验证。
- 新增接口时，在 `docs/` 中记录文档来源、配置方式、调用示例和已知限制；示例只使用占位凭据。

类目属性查询 `bg.glo.goods.attrs.get` 按同名前缀 JSON 参数字典维护，模型放在 `temu/category/goods_attrs_get.go`；多语言字段使用动态 `map[string]string`，`productCreateTime` 为毫秒时间戳，`chooseAllQualifySpec` 按 BOOLEAN 保留 bool。已通过模型字典与模拟测试，尚未真实查询联调。

创建规格 `bg.glo.goods.spec.create` 最初按截图实现，现以 PA 字典维护（页面更新 2026-01-11 14:19:11），请求含必填 `parentSpecId`、`specName`，响应为 `result.specId`。使用字典指定的 PA 网关，权限需真实接入时核验；写操作不自动重试，不假设名称可去重。

新增 PA 类目及货品接口以 docs/API-docs/PA_Category_Product_APIs.md 为索引。可选布尔查询条件保留指针以发送 false；可选切片使用 omitzero 区分省略和空数组。无业务参数的方法只接收 context；所有方法不自动重试。旧 models 测试已迁移到新的 PA 字典目录，本地执行全量测试需要这些 docs 资料。

调用示例必须按 examples/<业务域>/<完整接口名以下划线分词>/main.go 分组；修改目录时同步所有调用命令和源码错误提示。模型文件按接口末段词序命名，如 goods_add.go。旧结果名称保留兼容别名；不要仅为风格修改已公开字段的类型。商品列表可选切片与新接口统一使用 omitzero。

可分享的业务请求模板统一命名为 request.example.json，允许通过 .gitignore 白名单提交。模板仅含业务字段和结构占位值，不含公共凭据或真实商家数据；没有业务参数的接口不创建空模板。模板不自动参与真实调用，说明中须按示例实际取参方式指导填写；本地请求和结果 JSON 继续忽略，不覆盖已有请求。

使用 JSON 文件作为业务参数的调用示例统一通过 examples/config.LoadRequest() 读取接口目录中的 request.json，不内嵌 requestJSON 字符串或添加恒定空字符串判断。缺文件时提示复制同目录 request.example.json 并填写，不能自动使用占位模板发起请求，不覆盖本地参数。公共读取器严格接受一个 JSON 对象并拒绝未知字段。商品发布原有 product-request.json 和特殊数据准备流程保持明确记录。

商品发布 request.example.json 按用户明确授权与其成功上架的 product-request.json 同步，属于业务案例模板例外（仍不含公共鉴权凭据）；其他接口模板继续使用结构占位值。发品缺本地文件时从该模板生成，不再维护重复的虚构双款 makeRequest 数据。

examples 是独立 Go Module（同路径子模块），通过 examples/go.mod 的本地 replace 使用根 SDK；主模块不得依赖示例模块。go.work 仅用于本地开发并忽略，不提交。克隆仓库后在根目录运行 go work init . ./examples，随后可从根目录运行既有示例命令。验证发布隔离时用 GOWORK=off go test ./...、GOWORK=off go vet ./...；开发环境对两个模块执行 go test ./... ./examples/...、go vet ./... ./examples/...。README 仅保留接口分类、安装和基本用法，开发发布细节放独立文档。

JIT 接口模型放 temu/jit/，调用方法放 temu/jit.go。bg.glo.jitmode.activate 当前字典只有必填 productId、productSkcId；result 是未展开 OBJECT。不要根据“开通/关闭”简介猜测开关字段或反转行为，不自动重试。
