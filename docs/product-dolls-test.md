# 双款娃娃发品测试

运行示例：[examples/product_add/main.go](../examples/product_add/main.go)。完整测试请求已经生成到项目根目录的 `product-request.json`，自动化测试使用相同的初始数据 `tests/testdata/doll-product-add.request.json`。

## 图片和款式

使用用户提供的图片 URL：

```text
https://img.kwcdn.com/product/open/51b943f3f6014cf6a2e002c0078421dc-goods.jpeg
```

图片中可见两款娃娃，按用户要求作为两个可选 SKU 编写。每款放在一个 SKC 下，每个 SKC 一个 SKU；两只不是默认一起出售的套装，照片中的吉他、滑板和木托盘不在标题中宣称附送。未推断品牌、BJD 结构、材质、身高、比例或适用年龄。

| 款式 | SKC 编码 | SKU 编码 | 虚构规格 ID | 虚构 supplierPrice | 虚构重量 |
| --- | --- | --- | --- | --- | --- |
| 灰蓝上衣款 | TEST-DOLL-SKC-GRAY-BLUE | TEST-DOLL-SKU-GRAY-BLUE | 20001 | 7780 | 180000 mg |
| 酒红上衣款 | TEST-DOLL-SKC-BURGUNDY | TEST-DOLL-SKU-BURGUNDY | 20002 | 7980 | 190000 mg |

用户实际执行后，product-result.json 返回 3000000：Main Sales Specification List Cannot Be Empty，已确认当前发品请求不接受空主销列表。当前根据用户新的调试要求，两款 mainProductSkuSpecReqs 各保留一个对象，其 specId、parentSpecId 设为 0，parentSpecName、specName 设为空字符串；默认值是否被平台接受尚未验证。商品规格属性和 SKU 规格仍使用 parentspec-result.json 返回的“颜色 / 1001”，本次不清空它们。

子规格 ID 20001/20002、VID 50011/50012、PID/模板 PID/引用 PID 等仍是测试值。父规格查询只有父规格名称和 ID，未返回子规格或类目属性模板，不能认为这些测试 ID 已因更换父规格而变得有效；后续规格校验仍需真实数据。本次仅修改请求与示例，没有重新提交商品，也没有断言发品已经成功。

素材、轮播、两款 SKC 预览图和 SKU 缩略图都使用所给 URL，不重新上传或裁切图片。

根据用户执行类目搜索后保存的 category-match-result.json，类目路径已替换为“玩具与游戏（25439）→ 娃娃及配件（25567）→ 娃娃（25589）”，四至十级填 0。该路径的父子 ID 关系、叶子标记和非隐藏状态已对照查询结果核验；并非根据搜索结果顺序盲选。请求 JSON、示例代码及自动化测试数据已同步。

按用户“其他数据随便编，测试用”的授权，属性及规格 ID、价格、尺寸 200×120×100 mm、重量、产地 CN 和非敏感标记仍为虚构测试值，未从平台查询或通过实物确认。当前按全托模板使用 supplierPrice，单字段单位未得到核实，不将数字解释为实际售价。平台可能拒绝这些测试参数。

已检查源图片是 JPEG，3278×4370，文件约 2.35 MB。当前商品文档对非服饰主图有 1:1、≤2M 等约束，本示例保留原图，是否符合该商品实际类目的要求仍需平台核验，不把图片上传成功等同于发品图片校验成功。

## 运行

从项目根目录执行：

```sh
go run ./examples/product_add
```

示例优先读取当前目录已有的 `product-request.json`，用户编辑后重新运行会使用该文件，不会重新生成或覆盖。仅在文件不存在时生成模板；JSON 格式、字段或类型不正确时直接报错，不回退模板覆盖原文件。

凭据默认是空字符串，未填写时在准备请求后退出，不发送请求。在 main.go 中填写 appKey、appSecret、accessToken 后，运行同一命令会真实调用 BgGloGoodsAdd。

示例读取已有业务请求 JSON（或在缺失时生成），再调用 SDK；平台返回响应时，完整响应同时保存到 `product-result.json` 并输出到 stdout，包括失败错误码或成功后的商品、SKC、SKU ID 与编码映射。最终已签名请求由 SDK 输出到 stderr。没有自动重试。

## 验证

```sh
go test ./tests/temu -run TestBgGloGoodsAddDollVariants -count=1
go test ./...
go vet ./...
```

自动化测试仅使用模拟 HTTP Transport，检查图片 URL、两款结构、规格匹配和返回 ID/编码映射。未使用真实凭据，未执行真实商品创建；填入虚构 ID 的联调请求不能被描述为已经符合 Temu 类目规则。
