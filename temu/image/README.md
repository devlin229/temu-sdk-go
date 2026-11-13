# 图片接口

图片接口模型从 `github.com/devlin229/temu-sdk-go/temu/image` 导入，方法由根包的 `*temu.Temu` 调用。

| 方法 | 用途 |
| --- | --- |
| BgGoodsImageUploadGlobal | Base64 图片上传 |
| BgAlgoImageTranslateGlobal | 提交图片翻译任务 |
| BgAlgoImageTranslateResultGlobal | 查询翻译结果 |
| BgCompliancepictureGetGlobal | 批量识别牛皮癣图片 |
| BgGloColorimageurlGet | 获取色块图 |
| BgGloFancyImageCm2in | 图片中 cm 转 inch |
| BgGloPicturecompressionGet | 高清图片压缩 |
| BgGoodsTexttopictureAddGlobal | 文字转图片 |

请求/响应统一命名为方法名加 Request/Respose，继续使用项目约定的 Respose 拼写。方法接收 context 和请求指针；nil 请求返回英文 SDK 错误。

翻译提交返回的 TaskID 是字符串，再通过查询方法获取结果；CustomTaskID 是调用方提供的 int64 自定义去重 ID。IsContainDetail 必填，false 也会发送。不自动重试、轮询或将 ResultCode 解释为完成状态。结果中的文本框、坐标和颜色数组保持原顺序。

色块图请求的 RGBA 与 Coor 为可选对象，内部整数使用指针：nil 省略，指向0发送0；不自动补 alpha 默认值或中心点。响应 Confidence 与识别结果 HasSpot 按字典保留为字符串。压缩 URL 数量建议不超过30，SDK 不截断、不去重、不分批，结果保留原图 URL、字节数和像素宽高。

**cm2in 使用非标准响应**：仅有 Code、ErrMsg、ResultURL。文档未定义成功码，SDK 不补 Success 等字段，不根据占位值0推测成功；HTTP正常且可解析时返回原响应和 nil error，包括非零 Code。调用方需核对响应，nil error 仅表示传输及解析成功。其他接口按 Success 判断业务失败，保留完整响应并返回 BusinessError。

全部接口保留平台原始错误文本，不自动输出请求、凭据或签名日志。HTTP/解析失败通过 ResponseError 保留原始正文，网络原因可通过 errors.Is/As 检查。SDK 只返回图片 URL，不下载或转存图片。

完整运行示例在 examples/image/，新增七个接口从同目录 request.json 读取业务参数；先复制 request.example.json，替换占位值并删除不用的可选字段。调用会真实提交或查询图片服务。当前新增接口通过本地契约和模拟测试，尚未真实商家联调。
