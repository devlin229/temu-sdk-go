# 官方数据字典

提供 2026-11-13 提取的十份字典快照：省份、尺码表分类、底板/模特/车型/净重/共计件数类目、名称长度、定制工艺、尺码规格分组。JSON 合计约 63 KiB，随 SDK 内嵌，不依赖 temp、docs、当前工作目录或网络。

```go
import "github.com/devlin229/temu-sdk-go/temu/dictionary"

snapshot, err := dictionary.Load()
if err != nil {
    return err
}
provinceID, found := snapshot.ProvinceID("广东省")
_ = provinceID // Temu ID 43000000000006，不是行政区划编码。
_ = found
rules := snapshot.CategoryRules(categoryID)
_ = rules.NetWeightRequired
```

已有的 `*product.BgGloGoodsAddRequest` 可显式预检查：

```go
warnings, err := request.CheckDictionaries(snapshot)
if err != nil {
    return err
}
// 每项包含 Code、Field、英文 Message 和官方 SourceURL。
// 调用方按 warnings 修正后，再自行决定是否发送商品发布请求。
_ = warnings
```

检查名称加权长度、CN 省份及全托类目的 SKU 净重、共计件数、模特信息；车型库只提示核对“见详情/通用”条件。带 ProductSemiManagedReq 的请求跳过全托条件检查。此方法不修改数据、不发请求、不进入 SDK 的自动调用流程；无警告也不代表完整请求合法。

ProductNameLimit 接收一级类目 ID。ProductNameLength 返回 Minimum/Maximum，官方未明确的非 ASCII 字母等边界字符保留计数范围，不直接使用 len(string)。SizeChartClassName 接收 classId，不能当作 sizeTemplateId。省份和规格 ID 使用 int64；未命名工艺、尺码子类型和重复 specId 均保留。

每份 Document 带官方 SourceURL、SourceUpdatedAt、ExtractedAt、PACompatibility 和 Notes，可通过 Data 读取完整结构化数据。部分原文引用旧 CN 接口，不能据此认定 PA 所有场景均适用。工艺层级外层 1/2 含义未明确，不推断为商家模式或 technologyType；发货时效表和两份半托专属字典未内置。快照缺少某个值不代表平台禁止该值，最终以平台校验为准。

每次 Load 返回独立实例，可在批量操作开始时加载一次复用。可修改实例或使用 LoadFS(os.DirFS("my-dictionaries")) 加载自行维护的完整快照，文件相对路径与 data/ 一致；不要在同一实例上并发读写。原始资料完整归档在项目本地 docs/data-dictionaries，运行时 JSON 仅省去重复原文 source_text。
