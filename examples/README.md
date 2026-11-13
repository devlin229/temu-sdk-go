# 接口调用示例

后续新增或修改示例时，生成的请求参数、响应结果及其他调试 JSON 一律保存到该接口的示例目录，与 main.go 放在一起。已有请求也从该目录读取，不在项目根目录生成或查找。文件路径按从项目根目录执行组织，本地请求和生成文件由 .gitignore 排除，只有不含凭据的 request.example.json 模板允许提交；用户编辑的请求不得被覆盖。

示例按 SDK 业务域分组，每个接口独立一个 main 包，目录名统一为接口 type 将点替换为下划线：

```text
examples/
├── category/ # 类目、属性、规格及类目纠正，11 个接口
├── product/  # 商品、仓库、运费、搬运及生命周期，11 个接口
├── image/    # 图片上传，1 个接口
├── config/   # 公共凭据加载
└── .env.example # 空配置模板
```

从项目根目录运行，例如：

```sh
go run ./examples/category/bg_glo_goods_cats_get
go run ./examples/product/bg_glo_goods_list_get
go run ./examples/image/bg_goods_image_upload_global
```

凭据由 `config.LoadCredentials()` 统一读取，业务参数仍在对应示例中准备。新 PA 示例的 requestJSON 只填写业务字段；已有上传、商品发布及规格等示例保留各自的数据准备流程。从项目根目录运行时，请求 JSON 与响应结果统一读写到各自接口示例目录。结果输出到 stdout，调试请求输出到 stderr。运行会真实调用接口，写接口不自动重试；自动化模拟测试仍位于 tests/。

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

全部 23 个接口示例调用 examples/config.LoadCredentials()，按字段优先使用系统环境变量：

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

21 个含业务参数的接口各提供 request.example.json，内容来自本地 PA 字段结构，不含公共鉴权字段或真实商家数据。父规格列表、爆款售罄查询没有业务参数，不创建空请求模板。

模板中的空字符串、0、false 和单元素数组只是类型占位，并非合法业务值或默认值。复制后按真实业务填写必填和条件必填项，删去不使用的可选字段；模板本身不自动发送。

商品发布示例读取 product-request.json。首次准备可以在项目根目录执行：

```sh
cp -n examples/product/bg_glo_goods_add/request.example.json examples/product/bg_glo_goods_add/product-request.json
```

如果本地文件已经存在，保留已有文件，避免覆盖。编辑 product-request.json 后再运行示例；原有缺文件时生成测试数据的行为保留。

其他 PA 示例目前仍从 main.go 的 requestJSON 字符串或结构体字段取参数：可以复制模板内容、填好业务值后粘贴到 requestJSON（建议使用 Go 反引号字符串），或据此填写结构体。单独复制成 request.json 不会自动被这些示例读取。图片上传示例继续按 imagePath 读取图片并编码，模板只用于查看上传选项结构。

Git 规则为：

```gitignore
/examples/**/*.json
!/examples/**/request.example.json
```

因此别人检出仓库能获得模板；本地请求、响应及调试 JSON 仍保持忽略。模板中不得添加 Token、App Secret、真实签名、真实商家参数或其他本地配置。
