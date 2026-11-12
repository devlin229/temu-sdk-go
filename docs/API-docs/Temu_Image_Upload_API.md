# 图片上传 SDK

接口：`bg.goods.image.upload.global`。模型依据 `docs/API-json/bg.goods.image.upload.global.{request,response}.parameters.json`，导出日期 2026-11-12；原页面更新时间 2025-06-22 20:00:52，来源为[图片上传接口页](https://agentpartner.temu.com/document?cataId=875198836203&docId=929743122710)。

通用授权、签名及公共参数见 [Temu_Developer_Guide.md](Temu_Developer_Guide.md)。本方法按本地导出定义使用 PA 网关 `https://openapi-b-partner.temu.com/openapi/router`，发送 POST JSON。适用商家模式与具体权限名称未包含在该参数导出中，真实接入时仍需按官方权限页核对，不能因本地测试通过而认定授权有效。

## 调用

可直接运行的示例位于 [examples/basic/main.go](../../examples/basic/main.go)。本地测试按用户要求直接在源码中填写 `appKey`、`appSecret`、`accessToken` 和 `imagePath`，不读取环境变量或命令行参数。图片路径默认为 `./product.png`，相对路径以运行命令时的当前目录为准。填写后在项目根目录执行：

```sh
go run ./examples/basic
# 如需保存响应：
# go run ./examples/basic > upload-result.json
```

上述上传命令会发出真实请求。

示例按用户提供的接口示例格式构造 Data URL：`data:<实际图片 MIME 类型>;base64,<Base64 内容>`，例如 JPEG 使用 `data:image/jpeg;base64,`，PNG 使用 `data:image/png;base64,`。MIME 类型根据文件内容识别，不按扩展名猜测。上传前将包含此前缀的完整字符串写入当前目录的 `image-base64.txt`，不添加换行；从项目根目录运行时，该文件位于项目根目录。文件内容与请求中的 image 字段一致，即使随后上传失败也会保留。
示例读取图片并编码为 Base64，设置 30 秒超时，成功时向 stdout 输出完整 JSON；业务失败也输出平台响应，同时向 stderr 输出错误并以非零状态退出。未设置凭据或图片读取失败时不会发送请求。示例默认只发送 image，外包装类型和图片处理选项按需在 request 中设置。

```go
package main

import (
    "context"
    "encoding/base64"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/devlin229/temu-sdk-go/temu"
    "github.com/devlin229/temu-sdk-go/temu/image"
)

func main() {
    appKey := "" // 填写 app_key
    appSecret := "" // 填写 app_secret
    accessToken := "" // 填写卖家授权 access_token
    if appKey == "" || appSecret == "" || accessToken == "" {
        log.Fatal("请填写 appKey、appSecret、accessToken")
    }
    content, err := os.ReadFile("product.png")
    if err != nil {
        log.Fatal(err)
    }
    encoded := "data:" + http.DetectContentType(content) + ";base64," + base64.StdEncoding.EncodeToString(content)
    client := temu.NewClient(appKey, appSecret, accessToken)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    response, err := client.BgGoodsImageUploadGlobal(ctx, &image.BgGoodsImageUploadGlobalRequest{
        Image: encoded,
    })
    if err != nil {
        log.Fatal(err)
    }
    if response.Result != nil {
        log.Printf("上传结果：%+v", response.Result)
    }
}
```

示例执行时会真实上传；本次验证仅使用模拟 Transport，未执行该示例。客户端沿用现有单例行为，仅第一次 `NewClient` 的凭据生效，默认 HTTP 超时为 30 秒。

## 请求字段与约束

请求结构体为 `image.BgGoodsImageUploadGlobalRequest`。所有业务字段在参数字典中均为可选；SDK 只拒绝 nil 请求，不自行新增必填、格式、尺寸或枚举校验，实际业务有效性由平台判断。

| 字段 | Go 类型 | 说明 |
| --- | --- | --- |
| image | string | 图片文件的 Base64 编码，支持 jpg/jpeg、png 等格式；SDK 按原值发送，不读取文件或自动编码 |
| imageBizType | int64 | 0 或 1；1 返回的 URL 用于货品发布外包装 |
| options | *image.Options | 图片处理选项 |
| options.cateId | int64 | 叶子类目 ID，智能裁剪时生效 |
| options.doIntelligenceCrop | bool | 智能裁剪；说明为按 sizeMode 返回原图和裁剪图 |
| options.boost | bool | AI 清晰度提升 |
| options.sizeMode | int64 | 0 原图大小，1 为 800×800，2 为 1350×1800 |

普通字段使用值类型，调用时直接赋值，例如 `Image: encoded`、`ImageBizType: 1`。`Options` 保留指针，nil 时省略整个对象，设置时使用 `&image.Options{...}`。`imageBizType` 总是发送，包括零值 0；设置 Options 后，其 `doIntelligenceCrop`、`boost` 和 `sizeMode` 也总是发送，包括 false 和 0，避免使用 omitempty 丢掉有效枚举或开关值。空 image 和零值 cateId 省略。

公共鉴权参数、秒级时间戳和签名由客户端添加至同一 JSON 根对象；不额外包装业务参数，不发送 app_secret，默认省略 version。

按当前调试要求，公共客户端在发送前向 stderr 输出请求方法、URL、请求头及签名后的完整 JSON 请求参数，包含 app_key、access_token、sign 和完整图片 Base64，不截断、不脱敏；app_secret 不输出。请求参数日志与实际发送的请求体使用同一份字节，stdout 仍只输出结果。需要分别保存时运行 `go run ./examples/basic > upload-result.json 2> request-debug.log`。

## 响应与错误

返回 `(*image.BgGoodsImageUploadGlobalRespose, error)`，响应保留 `success`、`errorCode`、`errorMsg` 和 `result`；结果包含 `url`、`imageUrl`、`urls`，SDK 不推断这三个字段的优先级。

- `success=true`：返回完整响应和 nil 错误；调用方检查 result 是否存在。
- `success=false`：返回完整响应和 `*temu.BusinessError`，可用 `errors.As` 提取 `APIType`、`Code`、`Message`。
- HTTP 非 2xx、网络错误或响应解码失败：返回 nil 响应和错误；取消、超时等底层错误经 `%w` 保留，可用 `errors.Is` 判断。

本地导出未包含业务错误码清单，SDK 原样保留平台错误信息，不自行定义错误码含义。方法不进行自动重试，出现超时或异常时由调用方核对结果后决定是否再次上传。

## 验证范围

`tests/temu/image_upload_test.go` 使用模拟 HTTP Transport 验证网关与接口名、公共参数、实际请求签名、可选零值、强类型响应、业务失败、HTTP 失败、非法响应、取消、传输错误及不重试行为。`tests/models/` 继续对照参数字典核对模型。

检查命令为 `go test ./...` 和 `go vet ./...`。这些检查不验证真实网关可用性、卖家权限、图片格式与平台处理效果。
