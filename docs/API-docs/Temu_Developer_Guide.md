# Temu 开发者文档

版本：1.1 ｜ 更新日期：2026-11-12

## 1. 用途与文档组织

本文维护当前 Temu Partner / PA 接入体系的通用规则。各业务接口的参数、权限、条件必填项、响应结构与业务错误码由对应接口文档维护；后续接口沿用本文件，不重复复制签名实现。不同区域、商家体系或尚未迁移的接口，须先核对其网关适用范围。

| 文档 | 内容 |
| --- | --- |
| Temu_Developer_Guide.md（本文） | 网关、授权、公共参数、签名、通用调用与排错 |
| Temu_Product_Publish_API.md | 商品发布 bg.glo.goods.add 的业务定义 |

## 2. 网关与授权

当前 PA 网关：

```text
https://openapi-b-partner.temu.com/openapi/router
```

网关迁移需同时调整网关地址、接口 type，并重新授权获取包含新接口权限的 token。未列入迁移清单的接口仍按官方说明使用原 CN 网关，不可统一替换所有接口。具体 type、权限包与应用类型见对应接口文档。

卖家授权管理：[进入授权管理](https://agentseller.temu.com/open/system-manage/client-manage)。

接口详情页只列出 `/openapi/router` 或“数据存储 CN”时，应结合迁移清单确定完整域名，不能仅凭 CN 字样判断网关。

## 3. 公共请求参数

| 参数接口 | 参数类型 | 是否必填 | 说明 |
| --- | --- | --- | --- |
| type | STRING | 是 | API接口名, 形如:bg.* |
| app_key | STRING | 是 | 已创建成功的应用标志 |
| timestamp | STRING | 是 | 时间戳，格式为UNIX时间（秒） ，长度10位，当前时间-300秒<=入参时间<=当前时间+300秒 |
| sign | STRING | 是 | API入参参数签名，签名值根据如下算法给出计算过程 |
| data_type | STRING | 否 | 请求返回的数据格式，可选参数固定为JSON |
| access_token | STRING | 是 | 用户授权令牌access_token，卖家中心—授权管理，申请授权生成 |
| version | STRING | 是 | API版本，默认为V1，无要求不传此参数 |

`version` 的官方“必填”列与说明存在不一致：表中标“是”，说明写“无要求不传”。建议无明确版本要求时按说明省略，不自行添加未知版本。

公共参数与业务字段放在同一个 JSON 根对象，不额外包成 `data` 或 `biz_content`。

本表摘自当前商品发布接口页；后续接口若明确声明不同版本或公共参数要求，以其接口页为准。

## 4. 签名规则

1. 收集实际发送的公共参数与业务参数，排除 `sign`；`app_secret` 只用于计算，不放进请求体。
2. 仅对最外层 key 按 ASCII 升序排序，内层对象不重新排序。
3. 按顺序连接每个 `key + value`，不加分隔符。对象和数组以紧凑 JSON 字符串参与拼接。
4. 在整个字符串首尾分别拼接 `app_secret`。
5. 对 UTF-8 字节进行 MD5，转为 32 位大写十六进制，填入 `sign`。

```text
sign = UPPERCASE(MD5(app_secret + 按外层key排序后拼接的key/value + app_secret))
```

参与签名的值必须与实际请求一致，尤其注意布尔值为 `true/false`，不能写成 `True/False` 或 `1/0`。不要在签名后修改价格、规格、图片、多语言字段或内层对象顺序。对可选空值，建议在构造请求时明确选择省略或传值，不依赖语言默认的 null 字符串转换。

### 签名参考实现（Python）

这是根据规则整理的实现示例，未进行真实接口联调。业务数据须满足对应接口的校验要求。

```python
import hashlib
import json

def temu_sign(params: dict, app_secret: str) -> str:
    def as_text(value):
        if isinstance(value, str):
            return value
        return json.dumps(value, ensure_ascii=False, separators=(",", ":"),
                          allow_nan=False)

    keys = sorted(k for k in params if k != "sign")
    raw = app_secret + "".join(k + as_text(params[k]) for k in keys) + app_secret
    return hashlib.md5(raw.encode("utf-8")).hexdigest().upper()
```

密钥保留在服务端。不要将含密钥的 MD5 前原文输出到普通业务日志。

## 5. 统一调用方式

当前示例按 POST JSON 组织；商品发布接口页未单独列出 HTTP 方法。新增接口时核对其方法与格式要求。

```bash
curl --request POST 'https://openapi-b-partner.temu.com/openapi/router' \
  --header 'Content-Type: application/json' \
  --data-binary '@request.json'
```

`request.json` 应包含对应接口的公共参数与业务参数，并已计算 sign。示例不含真实凭证，不可直接作为有效请求发送。

建议的调用顺序：确定接口与网关 → 确认授权 → 构造公共及业务参数 → 设置当前秒级时间戳 → 签名 → 发送 → 按接口返回结构处理结果。

## 6. 通用异常与重试

| 错误提示或场景 | 处理方式 |
| --- | --- |
| type not exists | 核对 type 拼写以及 type 与网关的对应关系 |
| access_token don't have this api access | 确认接口权限，卖家重新授权后换用新 token |
| sign is invalid, please check your sign calculation | 检查外层 key 排序、紧凑 JSON、布尔值、UTF-8 和签名后的参数变更 |
| 时间戳不合法 | 使用秒级 UNIX 时间，核对服务器时钟与接口允许误差 |
| HTTP 超时或系统异常 | 写接口先核对业务结果；无明确幂等保证时避免盲目重试 |
| 业务参数错误 | 根据对应接口 errorCode、errorMsg 修正数据 |

HTTP 请求成功不等于业务成功。当前商品发布接口使用 success、errorCode、errorMsg，新增接口需按其实际返回定义解析，不能直接假定所有接口结构一致。

日志建议保留接口 type、请求时间、业务编码与错误信息；对 access_token、app_secret 及敏感业务数据脱敏。涉及密钥的签名前原文不进入普通日志。

## 7. 官方说明差异记录

| 项目 | 官方原文差异 | 处理建议 |
| --- | --- | --- |
| version | 必填列为“是”，说明为“默认为V1，无要求不传此参数” | 无明确要求时按说明省略；若具体接口另有要求，遵循该接口 |

商品字段的条件必填、体积描述和编码文案差异记录在商品发布接口文档中。

## 8. 后续接口文档约定

每个新增业务接口单独建文档，包含：用途、适用商家模式、type、网关引用、权限、业务字段及嵌套路径、业务约束、请求示例、响应、错误码、说明差异和来源版本。

公共规则引用本文相应章节；某接口的特殊规则写在接口文档中，不修改为所有接口共用的默认规则。维护版本和来源更新时间，避免不同接口迁移进度被混淆。

## 9. 官方来源

- [PA网关调用说明](https://agentpartner.temu.com/document?cataId=875198836203&docId=929750644349)，页面更新时间 2026-11-11 17:06:27。
- [签名规则](https://agentpartner.temu.com/document?cataId=875196199516&docId=896167235113)，页面更新时间 2025-05-11 15:18:33。
- [商品发布接口页（公共参数来源）](https://agentpartner.temu.com/document?cataId=875198836203&docId=925526695187)，页面更新时间 2026-08-25 22:03:37。

本文件为官方页面的开发整理版；示例未进行真实接口联调。

