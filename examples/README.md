# 接口调用示例

examples 是独立 Go Module，依赖本地上级 SDK；主 SDK 的 go get 下载包不会包含此目录。克隆完整仓库后，首次在项目根目录创建本地工作区（若已有 go.work 则跳过）：

```sh
go work init . ./examples
```

之后继续从项目根目录运行下方 go run 命令。本地 go.work 和 go.work.sum 已被 Git 忽略，不作为 SDK 发布内容。

验证两个模块使用 `go test ./... ./examples/...` 和 `go vet ./... ./examples/...`；示例配置测试位于 examples/tests/config。

后续新增或修改示例时，生成的请求参数、响应结果及其他调试 JSON 一律保存到该接口的示例目录，与 main.go 放在一起。已有请求也从该目录读取，不在项目根目录生成或查找。文件路径按从项目根目录执行组织，本地请求和生成文件由 .gitignore 排除，只有不含凭据的 request.example.json 模板允许提交；用户编辑的请求不得被覆盖。

示例按 SDK 业务域分组，每个接口独立一个 main 包，目录名统一为接口 type 将点替换为下划线：

```text
examples/
├── category/ # 类目、属性、规格及类目纠正，11 个接口
├── product/  # 商品、仓库、运费、搬运及生命周期，11 个接口
├── image/    # 图片上传，1 个接口
├── jit/      # JIT 操作，1 个接口
├── config/   # 公共凭据加载
└── .env.example # 空配置模板
```

从项目根目录运行，例如：

```sh
go run ./examples/category/bg_glo_goods_cats_get
go run ./examples/product/bg_glo_goods_list_get
go run ./examples/image/bg_goods_image_upload_global
```

凭据由 `config.LoadCredentials()` 统一读取，业务参数仍在对应示例中准备。使用公共读取方法的新 PA 示例从同目录 request.json 加载业务字段；已有上传、商品发布及规格等示例保留各自的数据准备流程。从项目根目录运行时，请求 JSON 与响应结果统一读写到各自接口示例目录。结果输出到 stdout，调试请求输出到 stderr。运行会真实调用接口，写接口不自动重试；自动化模拟测试仍位于 tests/。

| 分组 | 接口 | 用途 | 目录 |
| --- | --- | --- | --- |
| category | `bg.glo.goods.accessories.get` | 货品包装清单类型查询 | `examples/category/bg_glo_goods_accessories_get` |
| category | `bg.glo.goods.attrs.get` | 货品发布类目属性模板查询 | `examples/category/bg_glo_goods_attrs_get` |
| category | `bg.glo.goods.category.match` | 类目搜索 | `examples/category/bg_glo_goods_category_match` |
| category | `bg.glo.goods.cats.get` | 子类目查询 | `examples/category/bg_glo_goods_cats_get` |
| category | `bg.glo.goods.catsmandatory.get` | 类目必填信息接口 | `examples/category/bg_glo_goods_catsmandatory_get` |
| category | `bg.glo.goods.parentspec.get` | 查询父规格列表 | `examples/category/bg_glo_goods_parentspec_get` |
| category | `bg.glo.goods.spec.create` | 创建规格 | `examples/category/bg_glo_goods_spec_create` |
| category | `bg.glo.oem.page.list` | 检索OE/OEM数据 | `examples/category/bg_glo_oem_page_list` |
| category | `bg.goods.attribute.mapping.global` | 内外属性映射 | `examples/category/bg_goods_attribute_mapping_global` |
| category | `bg.goods.redress.correctrecord.query` | 查询商品类目纠正列表 | `examples/category/bg_goods_redress_correctrecord_query` |
| category | `bg.goods.redress.optionalcategory.correct` | 纠正商品类目 | `examples/category/bg_goods_redress_optionalcategory_correct` |
| jit | `bg.glo.jitmode.activate` | JIT 操作 | `examples/jit/bg_glo_jitmode_activate` |
| image | `bg.goods.image.upload.global` | bas64图片上传-global | `examples/image/bg_goods_image_upload_global` |
| product | `bg.btg.goods.stock.warehouse.list.get` | 根据站点查询可绑定的发货仓库信息接口 | `examples/product/bg_btg_goods_stock_warehouse_list_get` |
| product | `bg.glo.best.seller.invitation.query` | best seller招标单查询 | `examples/product/bg_glo_best_seller_invitation_query` |
| product | `bg.glo.goods.add` | 上传供应商货品 | `examples/product/bg_glo_goods_add` |
| product | `bg.glo.goods.brand.get` | 货品品牌查询 | `examples/product/bg_glo_goods_brand_get` |
| product | `bg.glo.goods.detail.get` | 商品详情查询接口 | `examples/product/bg_glo_goods_detail_get` |
| product | `bg.glo.goods.list.get` | 商品列表查询 | `examples/product/bg_glo_goods_list_get` |
| product | `bg.glo.goods.migrate` | 货品搬运接口 | `examples/product/bg_glo_goods_migrate` |
| product | `bg.glo.goods.removed.get` | 卖家中心已废弃列表 | `examples/product/bg_glo_goods_removed_get` |
| product | `bg.glo.goods.topselling.soldout.get` | 批量查询爆款售罄商品 | `examples/product/bg_glo_goods_topselling_soldout_get` |
| product | `bg.glo.logistics.template.get` | 查询运费模板列表 | `examples/product/bg_glo_logistics_template_get` |
| product | `bg.glo.product.search` | 查询货品生命周期状态 | `examples/product/bg_glo_product_search` |

每个接口的请求和结果 JSON 与 main.go 放在同一目录，继续从项目根目录执行上述命令。例如商品发布读取 `examples/product/bg_glo_goods_add/product-request.json`，结果写入同目录的 `product-result.json`；已有 6 个根目录 JSON 已原样迁移，生成文件仍由 .gitignore 排除。

## 公共凭据配置

全部 24 个接口示例调用 examples/config.LoadCredentials()，按字段优先使用系统环境变量：

```dotenv
TEMU_APP_KEY=
TEMU_APP_SECRET=
TEMU_ACCESS_TOKEN=
```

环境变量不存在或为空时，从项目 examples/.env 补齐该字段；如果仍有缺失，报错列出缺失变量和填写方式，不输出凭据值。环境变量已经完整时不会读取 .env。加载本地文件不会修改进程环境。

可以从项目根目录准备本地文件（已有文件不覆盖）：

```sh
cp -n examples/.env.example examples/.env
```

随后编辑 examples/.env 填写三个值，不再逐个修改 main.go 的凭据。examples/.env 已被 Git 忽略，.env.example 保留空值用于分享。

文件支持 KEY=value、可选 export 前缀、单引号、双引号与注释；双引号支持转义，不执行 shell、不展开变量。重复键以最后一项为准。格式错误提示行号，不输出整行内容。公共加载器可从项目根目录或子目录定位文件；接口示例的 JSON 读写仍按从项目根目录运行组织。

## 可分享的请求模板

22 个含业务参数的接口各提供 request.example.json。商品发布模板按用户授权复制其成功上架的单款请求；其余模板来自本地 PA 字段结构。全部模板不含公共鉴权凭据。父规格列表、爆款售罄查询没有业务参数，不创建空请求模板。

除商品发布成功案例外，其他模板中的空字符串、0、false 和单元素数组只是类型占位，并非合法业务值或默认值。复制后按真实业务填写必填和条件必填项，删去不使用的可选字段；模板本身不自动发送。

商品发布示例读取 product-request.json。首次准备可以在项目根目录执行：

```sh
cp -n examples/product/bg_glo_goods_add/request.example.json examples/product/bg_glo_goods_add/product-request.json
```

如果本地文件已经存在，保留已有文件，避免覆盖。编辑 product-request.json 后再运行示例；缺文件时从嵌入的 request.example.json 生成，和可分享模板使用同一份数据，不再生成旧的虚构双款参数。

原来内嵌 requestJSON 的 16 个 PA 示例已统一调用 config.LoadRequest()，读取接口目录中的 request.json，不再修改源码里的 JSON 字符串。先复制同目录的 request.example.json 为 request.json，填好业务参数、删去不需要的可选字段后运行；全可选接口可以将本地文件改成 {}。缺文件、JSON 格式错误或字段不匹配时直接报错，不自动创建、不覆盖、不回退到占位模板。原有类目搜索、属性查询和规格创建示例仍按 main.go 中明确的结构体字段取参。图片上传继续按 imagePath 读取图片并编码，商品发布继续读取 product-request.json。

Git 规则为：

```gitignore
/examples/**/*.json
!/examples/**/request.example.json
```

因此别人检出仓库能获得模板；本地请求、响应及调试 JSON 仍保持忽略。模板中不得添加 Token、App Secret、真实签名或本地配置。商品发布案例按本次用户授权保留业务 ID 和图片；复用前仍应核对自己的店铺、类目、规格、图片与商品信息。

例如包装清单查询，从项目根目录执行：

```sh
cp -n examples/category/bg_glo_goods_accessories_get/request.example.json examples/category/bg_glo_goods_accessories_get/request.json
# 编辑 request.json，填写业务值或按需要使用 {}
go run ./examples/category/bg_glo_goods_accessories_get
```

JIT 示例运行 `go run ./examples/jit/bg_glo_jitmode_activate`；先复制同目录 request.example.json 为 request.json 并填写真实货品和 SKC ID。该命令会提交 JIT 写操作，不自动重试。
