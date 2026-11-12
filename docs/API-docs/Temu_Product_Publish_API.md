# Temu 商品发布接口文档

版本：1.1 ｜ 更新日期：2026-11-12

通用接入规则参见配套文件 **Temu_Developer_Guide.md**（《Temu 开发者文档》）第 2～6 节。

## 1. 接口范围与用途

本文件仅定义商品发布接口 `bg.glo.goods.add`（官方名称：上传供应商货品）。用于创建、提交货品信息；接口成功返回不应直接解释为已审核通过或已面向消费者可售。商品审核和后续销售状态不在本文范围内。

适用范围以该接口的商家模式、店铺权限及类目规则为准。不要将本接口直接套用到其他 Temu 开放平台区域或不同商家体系。

## 2. 接口信息

| 项目 | 值 |
| --- | --- |
| 网关 | PA 网关，完整地址与授权步骤见《Temu 开发者文档》第 2 节 |
| 接口 type | `bg.glo.goods.add` |
| 迁移说明列出的等价 type | `temu.goods.add`；建议本文统一使用前者 |
| 旧 type | `bg.goods.add`，不要与 PA 网关混用 |
| 数据格式 | JSON |
| 权限包 | 货品API组 |
| 应用类型 | 官方权限表列为 He uses type、Self use type |
| 接入方式 | 下文按 POST JSON 组织调用示例；接口页未单独列出 HTTP 方法 |

请求公共参数与签名见《Temu 开发者文档》第 3、4 节；以下仅定义本接口业务字段。

## 3. 核心业务参数

“是/否”是官方字段级标记；可选对象内部的必填字段，仅在使用该对象时适用。标为可选也可能因商家模式、类目、站点成为条件必填。

| 字段 | 类型 | 官方必填 | 中文说明 |
| --- | --- | --- | --- |
| productName | STRING | 是 | 商品名称 |
| cat1Id | INTEGER | 是 | 一级类目 ID |
| cat2Id～cat10Id | INTEGER | 是 | 类目路径各层 ID，无对应层级按官方说明传 0 |
| materialImgUrl | STRING | 是 | 素材图 URL |
| productPropertyReqs | LIST | 是 | 商品属性，使用目标类目的属性模板数据 |
| productSpecPropertyReqs | LIST | 是 | 商品规格属性，使用目标类目的规格模板数据 |
| productSkcReqs | LIST | 是 | SKC 列表，内含 SKU |
| carouselImageUrls | LIST | 否 | 商品轮播图；服饰类目可由 SKC 聚合 |
| productI18nReqs | LIST | 否 | 多语言标题，元素含 language、productName |
| productWhExtAttrReq | OBJECT | 否 | 产地、外部商品链接、产地证明 |
| productSemiManagedReq | OBJECT | 否 | 半托管信息，含绑定站点 |
| productShipmentReq | OBJECT | 否 | 运费模板与承诺发货时效 |
| productWarehouseRouteReq | OBJECT | 否 | 站点与仓库路由 |
| goodsLayerDecorationReqs | LIST | 否 | 商品详情装修 |
| productGuideFileReqs | LIST | 否 | 多语言说明书 |
| productComplianceStatementReq | OBJECT | 否 | 合规声明协议版本与链接，不能自行编造协议值 |

其他可选字段及上述对象的全部嵌套字段见第 8 节。

### 3.1 商品属性与规格

`productPropertyReqs[]` 包含 `pid`、`templatePid`、`refPid`、`vid`、`propName`、`propValue`、`valueUnit`，以及可选 `numberInputValue`。
`productSpecPropertyReqs[]` 还包含 `specId`、`parentSpecId`、`valueGroupId`、`valueGroupName`。
属性、属性值和规格 ID 必须互相匹配，不应凭名称生成 ID。官方允许无值时使用 0 或空字符串的字段，详见完整字段表；不代表所有 ID 都能填 0。

### 3.2 SKC、SKU 与价格

| 层级 | 关键字段 | 说明 |
| --- | --- | --- |
| SKC | extCode | 外部编码，无编码时官方允许空字符串 |
| SKC | mainProductSkuSpecReqs | 主销售规格列表 |
| SKC | previewImgUrls | 预览图列表；表中标必填，说明指出非服饰类目非必填 |
| SKC | productSkuReqs | SKU 列表，官方说明服饰类目最多 10 个 |
| SKU | extCode、thumbUrl | SKU 外部编码与预览图；SKU extCode 的官方描述写为 SKC 编码（疑似文案笔误），应按其 SKU 层级使用 |
| SKU | currencyType | CNY 或 USD，必须与店铺支持币种一致 |
| SKU | supplierPrice | 供货价；半托管场景已废弃此字段，使用分站点价格 |
| SKU | siteSupplierPrices[] | 半托管分站点价格，元素为 siteId、supplierPrice |
| SKU | productSkuSpecReqs | SKU 规格，包含 specId、parentSpecId、specName、parentSpecName |
| SKU | productSkuWhExtAttrReq | 重量、体积、敏感属性等扩展信息 |
| SKU | productSkuStockQuantityReq | 仓库库存，按商家模式和站点要求使用 |

分站点 `supplierPrice` 单位为人民币分或美元美分，例如 77.8 元应传 7780。单字段 `supplierPrice` 在当前字段表未明确单位，接入前应核对实际模式规则，不能用商品零售价格代替供货申报价。

### 3.3 单位与图片约束

| 项目 | 要求 |
| --- | --- |
| SKU 重量 value | 毫克 mg；100 g = 100000 mg |
| SKU 体积 len/width/height | 毫米 mm；需满足最长边 ≥ 次长边 ≥ 最短边 |
| 体积字段文案冲突 | len 的英文描述写最短边（疑似文案笔误），但 inputLen 描述为最长边；联调时按最长边、次长边、最短边关系校验 |
| 净含量 netContentNumber | 实际值乘 1000；例如 1 L 传 1000 |
| 发货时效 shipmentLimitSecond | 秒，选项 86400、172800、259200；259200 仅定制商品可用 |
| 非服饰素材图、轮播图、SKU 预览图 | 1:1，宽高均 ≥800 px，≤2M |
| 服饰 SKC 轮播图 | 3:4，宽 ≥1340 px、高 ≥1785 px，≤2M |
| 图片格式 | JPG、JPEG、PNG |
| 图片语言 | 按目的地官方语言和站点要求提供 |

以上图片限制来自当前接口错误码说明，具体类目、语言版本仍应以平台校验为准。图片应使用平台认可的有效 URL；不要把本地路径或 Base64 当作 URL 传入。

## 4. 请求结构示例

以下是普通非敏感商品的结构示意，并非可直接提交的完整业务样例。演示 ID、图片地址、规格值必须替换；属性数组不得因示例为空而在真实请求中留空。半托管还需按实际情况补充站点、仓库、分站点价格等字段。

```json
{
  "type": "bg.glo.goods.add",
  "app_key": "YOUR_APP_KEY",
  "timestamp": "UNIX_SECONDS",
  "data_type": "JSON",
  "access_token": "YOUR_ACCESS_TOKEN",
  "productName": "示例玩具",
  "cat1Id": 10001,
  "cat2Id": 10002,
  "cat3Id": 10003,
  "cat4Id": 0,
  "cat5Id": 0,
  "cat6Id": 0,
  "cat7Id": 0,
  "cat8Id": 0,
  "cat9Id": 0,
  "cat10Id": 0,
  "materialImgUrl": "https://example.com/material.png",
  "carouselImageUrls": [
    "https://example.com/carousel.png"
  ],
  "productPropertyReqs": [],
  "productSpecPropertyReqs": [],
  "productWhExtAttrReq": {
    "outerGoodsUrl": "",
    "productOrigin": {
      "region1ShortName": "CN"
    }
  },
  "productSkcReqs": [
    {
      "extCode": "TOY-RED",
      "mainProductSkuSpecReqs": [
        {
          "specId": 20001,
          "parentSpecId": 20000,
          "parentSpecName": "颜色",
          "specName": "红色"
        }
      ],
      "previewImgUrls": [
        "https://example.com/skc.png"
      ],
      "productSkuReqs": [
        {
          "extCode": "TOY-RED-ONE",
          "currencyType": "CNY",
          "supplierPrice": 7780,
          "thumbUrl": "https://example.com/sku.png",
          "productSkuSpecReqs": [
            {
              "specId": 30001,
              "parentSpecId": 30000,
              "parentSpecName": "规格",
              "specName": "单件"
            }
          ],
          "productSkuWhExtAttrReq": {
            "productSkuWeightReq": {
              "value": 100000
            },
            "productSkuVolumeReq": {
              "len": 200,
              "width": 100,
              "height": 50
            },
            "productSkuSensitiveLimitReq": {},
            "productSkuSensitiveAttrReq": {
              "isSensitive": 0
            }
          }
        }
      ]
    }
  ],
  "sign": "CALCULATED_UPPERCASE_MD5"
}
```

上述 supplierPrice 为占位示例值；其单位须按商家模式复核（第 3.2 节）。半托管改用 `siteSupplierPrices`，不要直接沿用该单字段。

完整请求需合并通用参数并计算签名；发送方式参见《Temu 开发者文档》第 5 节。

## 5. 返回参数

| 参数接口 | 参数类型 | 说明 |
| --- | --- | --- |
| result | OBJECT | result |
| productSkuList | LIST | SKU List |
| $item | OBJECT | - |
| productSkuId | INTEGER | sku id |
| extCode | STRING | Sku External Code |
| skuSpecList | LIST | Sku Specification |
| $item | OBJECT | - |
| specId | INTEGER | Specification ID |
| parentSpecName | STRING | Parent Specification Name |
| parentSpecId | INTEGER | Parent Specification ID |
| specName | STRING | Specification Name |
| productSkcId | INTEGER | skc id |
| productId | INTEGER | Product ID |
| productSkcList | LIST | SKC List |
| $item | OBJECT | - |
| productSkcId | INTEGER | skc id |
| success | BOOLEAN | status |
| errorCode | INTEGER | error code |
| errorMsg | STRING | error message |

`productSkuList`、`productId`、`productSkcList` 位于 result 内；SKU 结果中的规格位于 skuSpecList。以 `success` 判断业务是否成功，并保留 errorCode、errorMsg 便于排查。

下例仅展示返回结构，ID 为示意值；当前页面未提供可核验的成功 errorCode 固定值，因此示例不假定其值。

```json
{
  "success": true,
  "result": {
    "productId": 90000001,
    "productSkcList": [{"productSkcId": 90000002}],
    "productSkuList": [{
      "productSkuId": 90000003,
      "productSkcId": 90000002,
      "extCode": "TOY-RED-ONE",
      "skuSpecList": [{
        "specId": 30001,
        "parentSpecId": 30000,
        "parentSpecName": "规格",
        "specName": "单件"
      }]
    }]
  }
}
```

## 6. 发品流程与异常处理

1. 准备目标类目路径、属性/规格模板、素材 URL、产地及实际商品重量体积；按模式补齐站点、仓库、价格与物流信息。
2. 构造请求并检查必填和条件必填字段，生成当前时间戳、计算签名后提交。
3. 成功后保存 productId、productSkcId、productSkuId 与本地编码映射。
4. 参数、类目、规格、图片或权限错误应修正后重试；超时或系统错误先核对是否已创建，避免盲目重复发品。当前页面未列出幂等键，不能假定 extCode 自动去重。

本文不展开类目、图片上传、库存、查询或审核接口；这些数据依赖只说明所需输入，不扩展其接口定义。

| 本接口场景 | 处理方式 |
| --- | --- |
| 2000168 相同或高度相似商品 | 先核对是否已创建，避免重复提交 |
| 2000197 今日发品上限 | 次日再发或联系运营 |

网关、权限、签名和时间戳排错见《Temu 开发者文档》第 6 节。

## 7. 官方业务错误码

| 错误码 | 错误描述 | 解决办法 |
| --- | --- | --- |
| 1000001 | 服务器开小差 | 一般是系统抖动，可参考具体报错文案尝试解决或重试，如果还不行请联系管理员 |
| 1000003 | 参数错误 | 结合参数错误的具体原因排查 |
| 1000005 | 系统异常 | 尝试重试，如果还不行请联系管理员 |
| 2000011 | 自定义规格属性校验失败 | 请结合校验失败具体原因检查规格入参 |
| 2000044 | 商品体积录入有误，请遵循最长边 ≥ 次长边 ≥ 最短边 | 商品体积录入有误，请遵循最长边 ≥ 次长边 ≥ 最短边 |
| 2000081 | 不合法或不可用的品牌 | 输入的品牌信息不可用或id不正确 |
| 2000096 | 体积内容必须遵循最长边 ≥ 次长边 ≥ 最短边原则 | 体积内容必须遵循最长边 ≥ 次长边 ≥ 最短边原则 |
| 2000135 | 当前类目净含量必填 | 请填写净含量 |
| 2000177 | 半托管商品英文标题最少需要x个字 | 英文标题不满足字数要求，请重新输入 |
| 2000200 | 属性值id[x]的属性值名称有误 | 属性值id与属性值不匹配，请检查入参 |
| 2000202 | 商品不符合大件标准，不可使用大件商品运费模版 | 请选择非大件运费模板，或者重新维护商品体积 |
| 6000002 | 货品属性校验失败 | 请结合校验失败具体原因检查属性入参 |
| 6000012 | 尺码表校验失败 | 请结合校验失败具体原因检查尺码表入参 |
| 6000059 | 童鞋适用年龄和鞋子不匹配，请确认后填写 | 请确认填写内容 |
| 2000004 | 不合法的规格属性 | specId与specName不匹配，请检查入参 |
| 2000009 | 不合法的类目 | 入参类目id不合法，或者查询类目信息接口抖动，请重试，如果不行请更换类目 |
| 2000010 | 属性模板查询失败 | 接口抖动，或者当前类目未配置属性模板，请尝试重试，如果不行联系管理员处理 |
| 2000014 | 服饰类目skc轮播图[x]校验失败，应符合宽高比例为3:4，宽>=1340px，高>=1785px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000017 | 素材图[x]校验失败，应符合图片宽高比例为1:1，宽>=800px，高>=800px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000018 | 非服饰类目货品轮播图x校验失败，应符合宽高比例为1:1，宽>=800px，高>=800px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000020 | 非服饰类目sku预览图[x]校验失败，应符合宽高比例为1:1，宽>=800px，高>=800px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000021 | 图片[x]格式校验失败，只允许.JPG .JPEG .PNG | 请上传JPG、JPEG、PNG的图片 |
| 2000025 | 当前图片中文字请使用销售目的地官方语言，请重新上传[x] | 图片存在中文，请重新上传 |
| 2000026 | 图片是否存在牛皮癣校验失败,请重试 | 查询接口失败，请重试 |
| 2000031 | 服装类目skc下价格需要保持一致，请进行调整 | 服装类目skc下价格需要保持一致，请进行调整 |
| 2000037 | 市场部店铺发布货品需要外部商品链接 | 请补齐外部商品链接 |
| 2000060 | 请选择正确的币种 | 输入的币种入参不合法 |
| 2000061 | 暂不支持变更币种 | 报价格币种需和店铺支持币种保持一致 |
| 2000077 | 模特数据校验失败 | 请结合具体失败原因解决 |
| 2000079 | 录入申报价格大于x元，商品无法创建成功 | 达到申报价格上限，请合理输入 |
| 2000094 | 说明书文件[x]校验失败，单页应<=[x]M，长x宽应为1600*1200 | 上传说明书不符合格式要求，请重新上传 |
| 2000102 | 当前不支持设置备货区域 | 当前店铺不支持设置备货区域，如有疑问请咨询管理员 |
| 2000114 | 图片不合法，不支持空url | 请勿传入空url |
| 2000125 | 运费模板不存在 | 输入的运费模板id不正确，或者查询运费模板信息失败，请稍后重试，如果不行请联系管理员 |
| 2000127 | 货品运费模板校验失败 | 运费模板运费、区域信息校验失败，请结合具体错误信息解决 |
| 2000146 | 产地必填 | 请补充产地 |
| 2000148 | 说明书未上传 | 请上传说明书 |
| 2000154 | 说明书英文内容不合格，请重新上传 | 说明书英文内容不合格，请重新上传 |
| 2000158 | URL域名校验不通过或URL包含不合法字符串 | 请检查入参位URL的字段 |
| 2000159 | 说明书缺少必要语言 | 说明书缺少必要语言 |
| 2000161 | 说明书语言错误 | 说明书语言入参不合法，请检查 |
| 2000165 | 商品标签未填写，请重新上传 | 请填写商品标签 |
| 2000168 | 创建失败。您已创建相同或高度相似的商品，建议您重新编辑商品信息，避免重复创建相同或高度相似的商品 | 重新编辑商品信息，避免重复创建相同或高度相似的商品 |
| 2000171 | 创建失败，存在重复的货品属性 | 重新编辑商品信息，避免重复创建相同或高度相似的商品 |
| 2000173 | 您当前帐户预留金额不足，无法发布商品，请前往【结算管理-资金中心】充值 | 请前往【结算管理-资金中心】充值 |
| 2000184 | 合规声明未签署 | 请签署合规声明 |
| 2000187 | 您当前账户预留金额不足，无法选择定制商品，请前往【结算管理-资金中心】充值 | 您当前账户预留金额不足，无法选择定制商品，请前往【结算管理-资金中心】充值 |
| 2000188 | 当前类目不支持定制 | 当前类目不支持定制 |
| 2000193 | jit商品不允许开启定制 | jit商品不允许开启定制 |
| 2000197 | 已超出今日发品数量限制，若有需要请联系运营 | 明日再发品，或者联系管理员加白 |
| 2000198 | 您当前账户预留金额不足，无法选择定制商品，请等待货款回款足额后再进行开启定制 | 您当前账户预留金额不足，无法选择定制商品，请等待货款回款足额后再进行开启定制 |
| 2000199 | 合规声明信息错误 | 合规声明信息错误，请检查入参 |
| 2000204 | 分站点申报价格校验失败 | 检查分站点报价格入参，比如站点信息是否和经营站点一致 |
| 2000301 | 商详装修楼层ID不合法 | 请检查商详装修楼层ID入参 |
| 2000302 | 商详装修楼层优先级不合法 | 商详装修楼层优先级不能重复，请检查入参 |
| 2000306 | 第x个楼层的商详装修图片不合法 | 上传图片不符合格式要求，请重新上传 |
| 2000307 | 第x个楼层的商详装修文字数量不合法 | 文字数量达到上限，具体阈值请咨询管理员 |
| 2000317 | 服饰类目多语言skc轮播图[x]校验失败，应符合宽高比例为3:4，宽>=1340px，高>=1785px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000318 | 非服饰类目货品多语言轮播图[x]校验失败，应符合宽高比例为1:1，宽>=800px，高>=800px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000319 | 非服饰类目sku多语言预览图[x]校验失败，应符合宽高比例为1:1，宽>=800px，高>=800px，<=2M | 上传图片不符合格式要求，请重新上传 |
| 2000320 | 视频未转码 | 请先转码 |
| 2000322 | 当前已操作退店，发品失败 | 联系管理员解决 |
| 6000009 | 发布商品失败 | 请结合具体错误原因解决 |
| 6000011 | 所选类目不合法 | 类目路径不合法，请检查入参 |
| 6000018 | 货号不规范，请使用字母、数字和标点符号维护货号！ | 货号不规范，请使用字母、数字和标点符号维护货号！ |
| 6000027 | 视频比例仅允许1:1、 4:3、16:9 | 视频比例仅允许1:1、 4:3、16:9 |
| 6000033 | 尺码发布请按正确类型勾选后提交 | 尺码发布请按正确类型勾选后提交 |
| 6000056 | 请上传中东英语SKU预览图 | 请上传中东英语SKU预览图 |
| 6000058 | 请上传英国英语SKU预览图 | 请上传英国英语SKU预览图 |
| 6000064 | 输入内容存在违规内容，请重新调整货号后输入 | 输入内容存在违规内容，请重新调整货号后输入 |
| 6000081 | 库存信息校验失败 | 请结合具体报错信息解决 |
| 6000096 | 视频大小不能超过xMB | 检查视频大小 |
| 6000097 | 视频时长不能超过x秒 | 检查视频时长 |
| 6000108 | 产地证明文件[x]校验失败，应符合[x]格式，且<=[x]M | 上传的产地证明文件不符合格式要求，请重新上传 |
| 6000135 | 泛欧售卖需包含所有欧盟站点 | 泛欧售卖需包含所有欧盟站点 |
| 6000136 | 当前商品不满足非泛欧售卖条件，请提交反馈，或联系对接运营处理 | 当前商品不满足非泛欧售卖条件，请提交反馈，或联系对接运营处理 |
| 6000137 | 仅欧盟支持设置站点售卖模式 | 仅欧盟支持设置站点售卖模式 |
| 6000139 | 当前商品不支持选择未开站站点：xxx | 移除提示的未开站站点 |
| 6000141 | 站点不合法：xxx | 移除不合法的站点 |
| 7000035 | 多语言规格名称重复，x：规格id：x 翻译重复，请联系运营修改翻译内容 | 请联系运营修改翻译内容 |
| 7000048 | 请上传当地语种预览图 | 请上传当地语种预览图 |
| 7000049 | 请上传当地语种轮播图 | 请上传当地语种轮播图 |
| 7000050 | 请填写当地语种商品名称 | 请填写当地语种商品名称 |

## 8. 完整业务字段字典

以下按官方页面顺序分组，保留全部字段及英文原始说明，避免翻译丢失约束。字段路径中的 `[]` 表示数组元素。类型为官方类型；列表元素行单独保留，以明确元素类型。“必填”是当前字段在其所属对象内的官方标记。


### productSemiManagedReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productSemiManagedReq` | OBJECT | 否 | Semi-managed Merchant Information |
| `productSemiManagedReq.semiLanguageStrategy` | INTEGER | 否 | Semi-Managed - Material Language Strategy |
| `productSemiManagedReq.bindSiteIds` | LIST | 是 | Bound Site List |
| `productSemiManagedReq.bindSiteIds[]` | INTEGER | 否 | - |
| `productSemiManagedReq.semiManagedSiteMode` | INTEGER | 否 | Semi-managed Site Sales Mode |

### carouselImageUrls

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `carouselImageUrls` | LIST | 否 | Commodity Carousel, not required for apparel category, will be aggregated from skc |
| `carouselImageUrls[]` | STRING | 否 | - |

### productOuterPackageImageReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productOuterPackageImageReqs` | LIST | 否 | Outer Packaging Image |
| `productOuterPackageImageReqs[]` | OBJECT | 否 | - |
| `productOuterPackageImageReqs[].imageUrl` | STRING | 是 | Image Link |

### copyFromProductId

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `copyFromProductId` | INTEGER | 否 | Replicated Source Product ID |

### productName

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productName` | STRING | 是 | Product Name |

### productI18nReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productI18nReqs` | LIST | 否 | Product Multilingual Information Request |
| `productI18nReqs[]` | OBJECT | 否 | - |
| `productI18nReqs[].language` | STRING | 是 | Language Code |
| `productI18nReqs[].productName` | STRING | 是 | Product Name |

### sellOutProductIdSrc

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `sellOutProductIdSrc` | INTEGER | 否 | Sold Out Product ID Src |

### productShipmentReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productShipmentReq` | OBJECT | 否 | Product Delivery Information Request |
| `productShipmentReq.freightTemplateId` | STRING | 是 | Shipping Template ID |
| `productShipmentReq.shipmentLimitSecond` | INTEGER | 是 | Promised Delivery Time (unit: s). Basic options: 86400, 172800, 259200 (available only for custom items). |

### cat8Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat8Id` | INTEGER | 是 | Eight-level Category ID, pass 0 if none |

### cat4Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat4Id` | INTEGER | 是 | Fourth-Level Category ID, pass 0 if none |

### cat6Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat6Id` | INTEGER | 是 | Sixth-Level Category ID, pass 0 if none |

### showSizeTemplateIds

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `showSizeTemplateIds` | LIST | 否 | Key Display Size Table Template ID List |
| `showSizeTemplateIds[]` | INTEGER | 否 | - |

### cat2Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat2Id` | INTEGER | 是 | Secondary Category ID, pass 0 if none |

### productSpecPropertyReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productSpecPropertyReqs` | LIST | 是 | Product Specification Attribute |
| `productSpecPropertyReqs[]` | OBJECT | 否 | - |
| `productSpecPropertyReqs[].vid` | INTEGER | 是 | Basic Attribute Value ID, pass 0 if none |
| `productSpecPropertyReqs[].specId` | INTEGER | 是 | Specification ID |
| `productSpecPropertyReqs[].valueGroupId` | INTEGER | 是 | Attribute Value Group ID, 0 if none |
| `productSpecPropertyReqs[].parentSpecId` | INTEGER | 是 | Parent Specification ID |
| `productSpecPropertyReqs[].valueGroupName` | STRING | 是 | Attribute Group Name, pass empty string if none |
| `productSpecPropertyReqs[].valueUnit` | STRING | 是 | Unit of Attribute Value, empty string if not available |
| `productSpecPropertyReqs[].pid` | INTEGER | 是 | Attribute ID |
| `productSpecPropertyReqs[].templatePid` | INTEGER | 是 | Template Attribute ID |
| `productSpecPropertyReqs[].numberInputValue` | STRING | 否 | Numerical Input |
| `productSpecPropertyReqs[].propValue` | STRING | 是 | Basic Property Value |
| `productSpecPropertyReqs[].propName` | STRING | 是 | Reference Property Name |
| `productSpecPropertyReqs[].refPid` | INTEGER | 是 | Reference Property ID |

### vehicleLibraryRelationReqList

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `vehicleLibraryRelationReqList` | LIST | 否 | Vehicle Model Library Configuration |
| `vehicleLibraryRelationReqList[]` | OBJECT | 否 | - |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList` | LIST | 否 | Attribute Value Dependency Configuration |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[]` | OBJECT | 否 | - |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId5` | INTEGER | 否 | Attribute Value Depends on id5 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId4` | INTEGER | 否 | Attribute Value Depends on id4 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId3` | INTEGER | 否 | Attribute Value Dependent on id3 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId2` | INTEGER | 否 | Attribute Value Depends on id2 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId9` | INTEGER | 否 | Attribute Value Depends on id9 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId8` | INTEGER | 否 | Attribute Value Depends on id8 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId7` | INTEGER | 否 | Attribute Value Depends on id7 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId6` | INTEGER | 否 | Attribute Value Depends on id6 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId10` | INTEGER | 否 | Attribute Value Depends on id10 |
| `vehicleLibraryRelationReqList[].productPropValueDependencyReqList[].propertyValueDependencyId1` | INTEGER | 否 | Attribute Value Dependent ID 1 |
| `vehicleLibraryRelationReqList[].vehicleLibraryId` | INTEGER | 是 | Vehicle Model ID |

### productCarouseVideoReqList

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productCarouseVideoReqList` | LIST | 否 | Carousel Video |
| `productCarouseVideoReqList[]` | OBJECT | 否 | - |
| `productCarouseVideoReqList[].vid` | STRING | 是 | Video VID |
| `productCarouseVideoReqList[].coverUrl` | STRING | 是 | Video Cover Image (B-side stores the first frame image) |
| `productCarouseVideoReqList[].videoUrl` | STRING | 是 | Video URL |
| `productCarouseVideoReqList[].width` | INTEGER | 是 | Video Width |
| `productCarouseVideoReqList[].height` | INTEGER | 是 | Video Height |

### goodsLayerDecorationReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `goodsLayerDecorationReqs` | LIST | 否 | Product Details Decoration |
| `goodsLayerDecorationReqs[]` | OBJECT | 否 | - |
| `goodsLayerDecorationReqs[].floorId` | INTEGER | 否 | Floor ID, null: Create new, otherwise Update |
| `goodsLayerDecorationReqs[].goodsId` | INTEGER | 否 | Product ID |
| `goodsLayerDecorationReqs[].lang` | STRING | 是 | Language Type |
| `goodsLayerDecorationReqs[].type` | STRING | 是 | Component Type, image - image, text - text |
| `goodsLayerDecorationReqs[].priority` | INTEGER | 是 | Floor Sorting |
| `goodsLayerDecorationReqs[].contentList` | LIST | 是 | Floor Content |
| `goodsLayerDecorationReqs[].contentList[]` | OBJECT | 否 | - |
| `goodsLayerDecorationReqs[].contentList[].imgUrl` | STRING | 否 | Image Address--General |
| `goodsLayerDecorationReqs[].contentList[].textModuleDetails` | OBJECT | 否 | Text Module Details |
| `goodsLayerDecorationReqs[].contentList[].textModuleDetails.backgroundColor` | STRING | 是 | Background Color |
| `goodsLayerDecorationReqs[].contentList[].textModuleDetails.fontFamily` | INTEGER | 否 | Font Type |
| `goodsLayerDecorationReqs[].contentList[].textModuleDetails.fontSize` | INTEGER | 是 | Text Module Font Size |
| `goodsLayerDecorationReqs[].contentList[].textModuleDetails.align` | STRING | 是 | Text Alignment, left--Left Aligned；right--Right Aligned；center--Centered；justify--Justified |
| `goodsLayerDecorationReqs[].contentList[].textModuleDetails.fontColor` | STRING | 是 | Text Color |
| `goodsLayerDecorationReqs[].contentList[].width` | INTEGER | 否 | Image Width--General |
| `goodsLayerDecorationReqs[].contentList[].text` | STRING | 否 | Text Information--Text Module |
| `goodsLayerDecorationReqs[].contentList[].height` | INTEGER | 否 | Image Height--General |
| `goodsLayerDecorationReqs[].key` | STRING | 是 | Floor Type Key, currently defaults to 'DecImage' |

### productPropertyReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productPropertyReqs` | LIST | 是 | Product Attributes |
| `productPropertyReqs[]` | OBJECT | 否 | - |
| `productPropertyReqs[].vid` | INTEGER | 是 | Basic Attribute Value ID, pass 0 if none |
| `productPropertyReqs[].valueUnit` | STRING | 是 | Unit of Attribute Value, empty string if not available |
| `productPropertyReqs[].pid` | INTEGER | 是 | Attribute ID |
| `productPropertyReqs[].templatePid` | INTEGER | 是 | Template Attribute ID |
| `productPropertyReqs[].numberInputValue` | STRING | 否 | Numerical Input |
| `productPropertyReqs[].propValue` | STRING | 是 | Basic Property Value |
| `productPropertyReqs[].propName` | STRING | 是 | Reference Property Name |
| `productPropertyReqs[].refPid` | INTEGER | 是 | Reference Property ID |

### productGuideFileReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productGuideFileReqs` | LIST | 否 | Product Manual File Multilingual |
| `productGuideFileReqs[]` | OBJECT | 否 | - |
| `productGuideFileReqs[].fileName` | STRING | 是 | File Name |
| `productGuideFileReqs[].pdfMaterialId` | INTEGER | 是 | PDF File ID |
| `productGuideFileReqs[].languages` | LIST | 是 | Language |
| `productGuideFileReqs[].languages[]` | STRING | 否 | - |

### materialMultiLanguages

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `materialMultiLanguages` | LIST | 否 | Image Multilingual List |
| `materialMultiLanguages[]` | STRING | 否 | - |

### productWarehouseRouteReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productWarehouseRouteReq` | OBJECT | 否 | Commodity Warehouse Routing Request |
| `productWarehouseRouteReq.targetRouteList` | LIST | 是 | Target Self-Delivery Site-Warehouse Relationship |
| `productWarehouseRouteReq.targetRouteList[]` | OBJECT | 否 | - |
| `productWarehouseRouteReq.targetRouteList[].siteIdList` | LIST | 是 | Site ID. |
| `productWarehouseRouteReq.targetRouteList[].siteIdList[]` | INTEGER | 否 | - |
| `productWarehouseRouteReq.targetRouteList[].warehouseId` | STRING | 是 | Warehouse ID |
| `productWarehouseRouteReq.currentRouteList` | LIST | 否 | Current Self-Delivery Station-Warehouse Relationship |
| `productWarehouseRouteReq.currentRouteList[]` | OBJECT | 否 | - |
| `productWarehouseRouteReq.currentRouteList[].siteIdList` | LIST | 是 | Site ID. |
| `productWarehouseRouteReq.currentRouteList[].siteIdList[]` | INTEGER | 否 | - |
| `productWarehouseRouteReq.currentRouteList[].warehouseId` | STRING | 是 | Warehouse ID |

### goodsModelReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `goodsModelReqs` | LIST | 否 | Product Model List Request |
| `goodsModelReqs[]` | OBJECT | 否 | - |
| `goodsModelReqs[].modelProfileUrl` | STRING | 是 | Model Portrait |
| `goodsModelReqs[].sizeSpecName` | STRING | 是 | Trial Size Specification Name |
| `goodsModelReqs[].modelId` | INTEGER | 是 | Model ID, not transmitted for new virtual model scenarios |
| `goodsModelReqs[].sizeSpecId` | INTEGER | 是 | Fitting Size Specification ID |
| `goodsModelReqs[].modelWaist` | STRING | 否 | Model Waistline Text |
| `goodsModelReqs[].modelType` | INTEGER | 否 | Model Type, 1: Garment Model, 2: Shoe Model |
| `goodsModelReqs[].modelName` | STRING | 是 | Model Name |
| `goodsModelReqs[].modelHeight` | STRING | 否 | Model Height Text |
| `goodsModelReqs[].modelFeature` | INTEGER | 否 | Model Properties, 1: Real Model, 2: Virtual Model |
| `goodsModelReqs[].modelFootWidth` | STRING | 否 | Model Foot Width Text |
| `goodsModelReqs[].modelBust` | STRING | 否 | Model Bust Text |
| `goodsModelReqs[].modelFootLength` | STRING | 否 | Model Long Text |
| `goodsModelReqs[].tryOnResult` | INTEGER | 否 | Try-on Experience，TRUE_TO_SIZE(1, "Comfortable")TOO_SMALL(2, "Tight-fitting")TOO_LARGE(3, \"Relaxed\") |
| `goodsModelReqs[].modelHip` | STRING | 否 | Model Hip Text |

### sizeTemplateId

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `sizeTemplateId` | INTEGER | 否 | Size Chart Template ID |

### productOuterPackageReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productOuterPackageReq` | OBJECT | 否 | Product Outer Packaging Information |
| `productOuterPackageReq.packageShape` | INTEGER | 否 | Outer Packaging Shape |
| `productOuterPackageReq.packageType` | INTEGER | 否 | Outer Packaging Type |

### sourceInvitationId

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `sourceInvitationId` | INTEGER | 否 | Source Invitation ID |

### sensitiveTransNormalFileReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `sensitiveTransNormalFileReqs` | LIST | 否 | Sensitive Product Conversion to General Proof Document List |
| `sensitiveTransNormalFileReqs[]` | OBJECT | 否 | - |
| `sensitiveTransNormalFileReqs[].fileName` | STRING | 是 | File Name |
| `sensitiveTransNormalFileReqs[].fileUrl` | STRING | 是 | File Path |

### cat7Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat7Id` | INTEGER | 是 | Seventh-Level Category ID, pass 0 if none |

### sellOutProductId

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `sellOutProductId` | STRING | 否 | Sold Out Product ID |

### cat9Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat9Id` | INTEGER | 是 | 9th Level Category ID, pass 0 if none |

### cat3Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat3Id` | INTEGER | 是 | Third-Level Category ID, pass 0 if none |

### productDetailVideoReqList

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productDetailVideoReqList` | LIST | 否 | Detail Video |
| `productDetailVideoReqList[]` | OBJECT | 否 | - |
| `productDetailVideoReqList[].vid` | STRING | 是 | Video VID |
| `productDetailVideoReqList[].coverUrl` | STRING | 是 | Video Cover Image (B-side stores the first frame image) |
| `productDetailVideoReqList[].videoUrl` | STRING | 是 | Video URL |
| `productDetailVideoReqList[].width` | INTEGER | 是 | Video Width |
| `productDetailVideoReqList[].height` | INTEGER | 是 | Video Height |

### cat5Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat5Id` | INTEGER | 是 | Fifth-Level Category ID, pass 0 if none |

### cat1Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat1Id` | INTEGER | 是 | First-Level Category ID |

### carouselImageI18nReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `carouselImageI18nReqs` | LIST | 否 | Product Carousel Multi-Language Info Request |
| `carouselImageI18nReqs[]` | OBJECT | 否 | - |
| `carouselImageI18nReqs[].imgUrlList` | LIST | 否 | Image List, empty list implies deletion, null implies no change |
| `carouselImageI18nReqs[].imgUrlList[]` | STRING | 否 | - |
| `carouselImageI18nReqs[].language` | STRING | 是 | Language |

### sizeTemplateIds

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `sizeTemplateIds` | LIST | 否 | Size Chart Template ID List |
| `sizeTemplateIds[]` | INTEGER | 否 | - |

### productWhExtAttrReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productWhExtAttrReq` | OBJECT | 否 | Commodity Warehouse and Supply Chain Extension Properties Request |
| `productWhExtAttrReq.productOriginCertFiles` | LIST | 否 | Product Origin Certificate Files |
| `productWhExtAttrReq.productOriginCertFiles[]` | OBJECT | 否 | - |
| `productWhExtAttrReq.productOriginCertFiles[].fileName` | STRING | 是 | File Name |
| `productWhExtAttrReq.productOriginCertFiles[].fileUrl` | STRING | 是 | File Url |
| `productWhExtAttrReq.outerGoodsUrl` | STRING | 是 | Off-Site Product Link (Pass Empty String as Fallback) |
| `productWhExtAttrReq.productOrigin` | OBJECT | 是 | Product Origin |
| `productWhExtAttrReq.productOrigin.region2Id` | INTEGER | 否 | Secondary Region ID |
| `productWhExtAttrReq.productOrigin.region1ShortName` | STRING | 是 | First-Level Region Abbreviation (Two-Character Code) |

### productSkcReqs

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productSkcReqs` | LIST | 是 | Product SKC List |
| `productSkcReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].extCode` | STRING | 是 | Product SKC External Code, pass empty string if not available |
| `productSkcReqs[].productSkuReqs` | LIST | 是 | Product SKU List (up to 10 for Apparel Category) |
| `productSkcReqs[].productSkuReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].currencyType` | STRING | 是 | Currency (CNY: Chinese Yuan, USD: US Dollar) (Default: Chinese Yuan) |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq` | OBJECT | 否 | Product Multi-Package Request |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.numberOfPieces` | INTEGER | 否 | Quantity, default is 1 for single item |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.individuallyPacked` | INTEGER | 否 | Whether to use independent packaging (pass -1 to clear) |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.productSkuNetContentReq` | OBJECT | 否 | Net Content Request, passing an empty object indicates clearing |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.productSkuNetContentReq.netContentUnitCode` | INTEGER | 否 | Net Content Unit, 1: Fluid Ounce, 2: Milliliter, 3: Gallon, 4: Liter, 5: Gram, 6: Kilogram, 7: Troy Ounce, 8: Pound |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.productSkuNetContentReq.netContentNumber` | INTEGER | 否 | Net Content Value, the target value needs to be multiplied by 1000 before passing the value. For example, if you are passing "1 L", you need to pass "1000" in "netContentNumber". |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.mixedType` | INTEGER | 否 | Mixed set type, 1: different products, 2: same product with different specifications |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.totalNetContent` | OBJECT | 否 | Total Net Content |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.totalNetContent.netContentUnitCode` | INTEGER | 否 | Net Content Unit, 1: Fluid Ounce, 2: Milliliter, 3: Gallon, 4: Liter, 5: Gram, 6: Kilogram, 7: Troy Ounce, 8: Pound |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.totalNetContent.netContentNumber` | INTEGER | 否 | Net Content Value, the target value needs to be multiplied by 1000 before passing the value. For example, if you are passing "1 L", you need to pass "1000" in "netContentNumber". |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.pieceNewUnitCode` | INTEGER | 否 | Unit, 1: piece |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.skuClassification` | INTEGER | 否 | Sku Category, 1: Single Product, 2: Combination Set, 3: Mixed Suite |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.numberOfPiecesNew` | INTEGER | 否 | Total quantity contained |
| `productSkcReqs[].productSkuReqs[].productSkuMultiPackReq.pieceUnitCode` | INTEGER | 否 | Unit per Item, 1: Piece, 2: Pair, 3: Pack |
| `productSkcReqs[].productSkuReqs[].productSkuSuggestedPriceReq` | OBJECT | 否 | Product SKU Suggested Price Request |
| `productSkcReqs[].productSkuReqs[].productSkuSuggestedPriceReq.suggestedPriceCurrencyType` | STRING | 否 | Recommended Price Currency |
| `productSkcReqs[].productSkuReqs[].productSkuSuggestedPriceReq.suggestedPrice` | INTEGER | 否 | Suggested Price |
| `productSkcReqs[].productSkuReqs[].productSkuSuggestedPriceReq.specialSuggestedPrice` | STRING | 否 | Special Suggested Price |
| `productSkcReqs[].productSkuReqs[].siteSupplierPrices` | LIST | 否 | Site Supply Price List, for semi_managed merchant scenario only |
| `productSkcReqs[].productSkuReqs[].siteSupplierPrices[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].siteSupplierPrices[].siteId` | INTEGER | 是 | Declared Price Site ID |
| `productSkcReqs[].productSkuReqs[].siteSupplierPrices[].supplierPrice` | INTEGER | 是 | Site Declared Price, Unit: RMB: Fen, USD: Cent |
| `productSkcReqs[].productSkuReqs[].supplierPrice` | INTEGER | 否 | Supply Price, deprecated in semi_managed merchant scenario |
| `productSkcReqs[].productSkuReqs[].productSkuUsSuggestedPriceReq` | OBJECT | 否 | Product SKU US Suggested Price Request |
| `productSkcReqs[].productSkuReqs[].productSkuUsSuggestedPriceReq.suggestedPriceCurrencyType` | STRING | 否 | Recommended Price Currency |
| `productSkcReqs[].productSkuReqs[].productSkuUsSuggestedPriceReq.suggestedPrice` | INTEGER | 否 | Suggested Price |
| `productSkcReqs[].productSkuReqs[].productSkuUsSuggestedPriceReq.specialSuggestedPrice` | STRING | 否 | Special Suggested Price |
| `productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq` | OBJECT | 否 | Product SKU Inventory Request |
| `productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq.warehouseStockQuantityReqs` | LIST | 是 | Outbound Warehouse Inventory Request List |
| `productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq.warehouseStockQuantityReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq.warehouseStockQuantityReqs[].targetStockAvailable` | INTEGER | 是 | Target Inventory |
| `productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq.warehouseStockQuantityReqs[].warehouseId` | STRING | 是 | Warehouse ID |
| `productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq.warehouseStockQuantityReqs[].currentStockAvailable` | INTEGER | 否 | Current Stock |
| `productSkcReqs[].productSkuReqs[].extCode` | STRING | 是 | Product SKC External Code, pass empty string if not available |
| `productSkcReqs[].productSkuReqs[].productSkuThumbUrlI18nReqs` | LIST | 否 | SKU Preview Image Multilingual Information Request |
| `productSkcReqs[].productSkuReqs[].productSkuThumbUrlI18nReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuThumbUrlI18nReqs[].imgUrlList` | LIST | 否 | Image List, empty list implies deletion, null implies no change |
| `productSkcReqs[].productSkuReqs[].productSkuThumbUrlI18nReqs[].imgUrlList[]` | STRING | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuThumbUrlI18nReqs[].language` | STRING | 是 | Language |
| `productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq` | OBJECT | 否 | Product SKU Accessories Request |
| `productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq.productSkuAccessories` | LIST | 是 | Accessories List |
| `productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq.productSkuAccessories[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq.productSkuAccessories[].vid` | INTEGER | 是 | Accessory Property Value ID |
| `productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq.productSkuAccessories[].num` | INTEGER | 是 | Accessory Quantity |
| `productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq.productSkuAccessories[].unitCode` | INTEGER | 是 | Unit Code |
| `productSkcReqs[].productSkuReqs[].thumbUrl` | STRING | 是 | Preview Image |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq` | OBJECT | 是 | Product SKU Extended Attributes |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuWeightReq` | OBJECT | 是 | Product SKU Weight |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuWeightReq.inputUnit` | STRING | 否 | Input Unit |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuWeightReq.inputValue` | STRING | 否 | Input Weight Value |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuWeightReq.value` | INTEGER | 是 | Weight Value, unit mg |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSameReferPriceReq` | OBJECT | 否 | Same Style Reference |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSameReferPriceReq.url` | STRING | 否 | Same Style URL |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq` | OBJECT | 是 | Product SKU Sensitive Attribute Restriction Request |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.maxBatteryCapacityHp` | INTEGER | 否 | Maximum Battery Capacity (mWh) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.maxBatteryCapacity` | INTEGER | 否 | Maximum Battery Capacity (Wh) (Prioritizes maxBatteryCapacityHp) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.maxLiquidCapacity` | INTEGER | 否 | Maximum Liquid Capacity (mL) (Prefer using maxLiquidCapacityHp) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.maxLiquidCapacityHp` | INTEGER | 否 | Maximum Liquid Capacity (μL) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.maxKnifeLength` | INTEGER | 否 | Maximum Tool Length (mm) (Prefer using maxKnifeLengthHp) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.maxKnifeLengthHp` | INTEGER | 否 | Maximum Tool Length (μm) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.knifeTipAngle` | OBJECT | 否 | Blade Angle |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveLimitReq.knifeTipAngle.degrees` | INTEGER | 是 | Degree |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq` | OBJECT | 是 | Product SKU Volume |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.inputUnit` | STRING | 否 | Input Unit |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.len` | INTEGER | 是 | Shortest Side, unit mm |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.inputLen` | STRING | 否 | Length of the Longest Input Side |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.inputHeight` | STRING | 否 | Length of the Shortest Input Side |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.width` | INTEGER | 是 | Secondary Length, unit mm |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.inputWidth` | STRING | 否 | Input Secondary Length |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuVolumeReq.height` | INTEGER | 是 | Shortest Side, unit mm |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuNetWeightReq` | OBJECT | 否 | Product SKU Net Weight |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuNetWeightReq.inputUnit` | STRING | 否 | Input Net Weight Unit |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuNetWeightReq.inputValue` | STRING | 否 | Input Net Weight Value |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuNetWeightReq.value` | INTEGER | 否 | Net Weight Value, unit mg |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuBarCodeReqs` | LIST | 否 | Product SKU Barcode |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuBarCodeReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuBarCodeReqs[].code` | STRING | 否 | Barcode |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuBarCodeReqs[].codeType` | INTEGER | 否 | Barcode Type (1: EAN, 2: UPC, 3: ISBN) |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveAttrReq` | OBJECT | 是 | Product SKU Sensitive Attribute Request |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveAttrReq.sensitiveTypes` | LIST | 否 | Sensitive Type PURE_ELECTRIC(1, "Pure Electric"), INTERNAL_ELECTRIC(2, "Internal Electric"), MAGNETISM(3, "Magnetism"), LIQUID(4, "Liquid"), POWDER(5, "Powder"), PASTE(6, "Paste"), CUTTER(7, "Tool") |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveAttrReq.sensitiveTypes[]` | INTEGER | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveAttrReq.isSensitive` | INTEGER | 否 | Whether Sensitive Attribute, 0: Non-Sensitive, 1: Sensitive |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveAttrReq.sensitiveList` | LIST | 否 | Sensitive Type PURE_ELECTRIC(110001,"Pure Electric"), INTERNAL_ELECTRIC(120001, "Internal Electric"), MAGNETISM(130001, "Magnetism"), LIQUID(140001, "Liquid"), POWDER(150001, "Powder"), PASTE(160001, "Ointment"), CUTTER(170001, "Tool") |
| `productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuSensitiveAttrReq.sensitiveList[]` | INTEGER | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuSpecReqs` | LIST | 是 | Product SKU Specification List |
| `productSkcReqs[].productSkuReqs[].productSkuSpecReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkuReqs[].productSkuSpecReqs[].specId` | INTEGER | 是 | Specification ID |
| `productSkcReqs[].productSkuReqs[].productSkuSpecReqs[].parentSpecName` | STRING | 是 | Parent Specification Name |
| `productSkcReqs[].productSkuReqs[].productSkuSpecReqs[].parentSpecId` | INTEGER | 是 | Parent Specification ID |
| `productSkcReqs[].productSkuReqs[].productSkuSpecReqs[].specName` | STRING | 是 | Specification Name |
| `productSkcReqs[].mainProductSkuSpecReqs` | LIST | 是 | Main Sales Specification List |
| `productSkcReqs[].mainProductSkuSpecReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].mainProductSkuSpecReqs[].specId` | INTEGER | 是 | Specification ID |
| `productSkcReqs[].mainProductSkuSpecReqs[].parentSpecName` | STRING | 是 | Parent Specification Name |
| `productSkcReqs[].mainProductSkuSpecReqs[].parentSpecId` | INTEGER | 是 | Parent Specification ID |
| `productSkcReqs[].mainProductSkuSpecReqs[].specName` | STRING | 是 | Specification Name |
| `productSkcReqs[].previewImgUrls` | LIST | 是 | List of Preview Images, not required for non-apparel categories |
| `productSkcReqs[].previewImgUrls[]` | STRING | 否 | - |
| `productSkcReqs[].productSkcCarouselImageI18nReqs` | LIST | 否 | SKC Carousel Multi-Language Information Request |
| `productSkcReqs[].productSkcCarouselImageI18nReqs[]` | OBJECT | 否 | - |
| `productSkcReqs[].productSkcCarouselImageI18nReqs[].imgUrlList` | LIST | 否 | Image List, empty list implies deletion, null implies no change |
| `productSkcReqs[].productSkcCarouselImageI18nReqs[].imgUrlList[]` | STRING | 否 | - |
| `productSkcReqs[].productSkcCarouselImageI18nReqs[].language` | STRING | 是 | Language |
| `productSkcReqs[].isBasePlate` | INTEGER | 否 | Whether Baseplate |
| `productSkcReqs[].colorImageUrl` | STRING | 否 | SKC Color Block Diagram |

### productSaleExtAttrReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productSaleExtAttrReq` | OBJECT | 否 | Product Sales Side Extended Attribute Request |
| `productSaleExtAttrReq.inventoryRegion` | INTEGER | 否 | Inventory Area |
| `productSaleExtAttrReq.productSecondHandReq` | OBJECT | 否 | Second-hand Goods Information |
| `productSaleExtAttrReq.productSecondHandReq.isSecondHand` | BOOLEAN | 否 | Whether Second-hand |
| `productSaleExtAttrReq.productSecondHandReq.secondHandLevel` | INTEGER | 否 | Second-hand Grade (Condition) |
| `productSaleExtAttrReq.discreetShipping` | BOOLEAN | 否 | This parameter is required for adult products (category ID: 16571). Setting this service will display the corresponding logo on the consumer order page, which can improve payment conversion rates. Merchants cannot display the details of the purchased items on the shipping packaging. If a consumer complains about failure to maintain confidentiality and the complaint is verified, the merchant may be held liable for after-sales service. |
| `productSaleExtAttrReq.customizedTechnologyReq` | OBJECT | 否 | Customized Craft Request |
| `productSaleExtAttrReq.customizedTechnologyReq.twiceType` | LIST | 否 | Secondary Process |
| `productSaleExtAttrReq.customizedTechnologyReq.twiceType[]` | INTEGER | 否 | - |
| `productSaleExtAttrReq.customizedTechnologyReq.firstType` | INTEGER | 是 | Primary Process |
| `productSaleExtAttrReq.customizedTechnologyReq.technologyType` | INTEGER | 是 | Craft Type |
| `productSaleExtAttrReq.manufacturingLocation` | OBJECT | 否 | Manufacturing Location |
| `productSaleExtAttrReq.manufacturingLocation.region1ShortNames` | LIST | 否 | List of First-Level Region Two-Letter Short Names |
| `productSaleExtAttrReq.manufacturingLocation.region1ShortNames[]` | STRING | 否 | - |
| `productSaleExtAttrReq.productNoChargerReq` | OBJECT | 否 | Product No Charger Version Information (empty list must be sent to clear from has to none) |
| `productSaleExtAttrReq.productNoChargerReq.noChargerProductIds` | LIST | 是 | No Charger Version Product ID (pass empty list to clear) |
| `productSaleExtAttrReq.productNoChargerReq.noChargerProductIds[]` | INTEGER | 否 | - |
| `productSaleExtAttrReq.ipCodes` | LIST | 否 | ip code list, If ipBizType=1 (bg.glo.goods.brand.get), ipCodes is required. |
| `productSaleExtAttrReq.ipCodes[]` | STRING | 否 | - |

### personalizationSwitch

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `personalizationSwitch` | INTEGER | 否 | Whether to Support Customized Template, 0: Not Supported, 1: Supported |

### productCustomReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productCustomReq` | OBJECT | 否 | Commodity Customs Information |
| `productCustomReq.goodsLabelName` | STRING | 否 | Product Tag |
| `productCustomReq.isRecommendedTag` | BOOLEAN | 是 | Whether to select recommended tags |

### cat10Id

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `cat10Id` | INTEGER | 是 | Tenth-Level Category ID, pass 0 if none |

### materialImgUrl

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `materialImgUrl` | STRING | 是 | Material Image |

### productComplianceStatementReq

| 字段路径 | 类型 | 必填 | 官方说明 |
| --- | --- | --- | --- |
| `productComplianceStatementReq` | OBJECT | 否 | Compliance Signing Agreement |
| `productComplianceStatementReq.protocolVersion` | STRING | 是 | Protocol Version Number |
| `productComplianceStatementReq.protocolUrl` | STRING | 是 | Protocol Link |

## 9. 官方说明差异记录

| 字段 | 官方原文差异 | 处理建议 |
| --- | --- | --- |
| productSkcReqs[].previewImgUrls | 标记必填，但说明非服饰类目非必填 | 视为存在类目条件，按目标类目校验 |
| SKU 体积 len | 英文描述为最短边，而 inputLen 描述为最长边 | 可能为文案笔误；按最长边 ≥ 次长边 ≥ 最短边校验，并在联调时确认字段映射 |
| productSkcReqs[].productSkuReqs[].extCode | 位于 SKU 对象，说明写成 SKC 外部编码 | 按 SKU 层级理解，疑似复制文案问题 |
| SKU supplierPrice | 当前表未明确单字段金额单位；分站点 supplierPrice 明确为分/美分 | 单字段接入前核对模式规则，不将分站点单位直接视为已确认 |

原始英文说明保留于第 8 节。version 的公共参数差异见《Temu 开发者文档》第 7 节。

## 10. 来源与版本

[官方商品发布接口](https://agentpartner.temu.com/document?cataId=875198836203&docId=925526695187)：bg.glo.goods.add，页面更新时间 2026-08-25 22:03:37。网关与签名来源见《Temu 开发者文档》第 9 节。

本文件为官方页面的开发整理版。请求与返回示例为结构示意，未进行真实发品联调，未创建商品。

