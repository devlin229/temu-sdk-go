# temu-sdk-go

Temu Go SDK，支持公共签名、接口调用和请求/响应模型。需要 Go 1.27.1 或更高版本。

## 已接入分类

| 分类 | 数量 | 接口功能 |
| --- | --- | --- |
| 类目 `category` | 11 | 子类目、类目搜索、属性模板、必填信息、父规格、创建规格、包装清单、OEM、属性映射、类目纠正记录及纠正 |
| 商品 `product` | 11 | 商品发布、列表、详情、品牌、搬运、废弃列表、爆款售罄、生命周期、发货仓库、招标单、运费模板 |
| 图片 `image` | 1 | 图片上传 |

## 安装

```sh
go get github.com/devlin229/temu-sdk-go/temu@latest
```

## 使用

```go
package main

import (
    "context"
    "encoding/json"
    "log"
    "os"

    "github.com/devlin229/temu-sdk-go/temu"
    "github.com/devlin229/temu-sdk-go/temu/category"
)

func main() {
    client := temu.NewClient(
        os.Getenv("TEMU_APP_KEY"),
        os.Getenv("TEMU_APP_SECRET"),
        os.Getenv("TEMU_ACCESS_TOKEN"),
    )
    response, err := client.BgGloGoodsCatsGet(
        context.Background(),
        &category.BgGloGoodsCatsGetRequest{},
    )
    if err != nil {
        log.Fatal(err)
    }
    if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
        log.Fatal(err)
    }
}
```

方法名由接口名按点分词转为大驼峰，例如 `bg.glo.goods.add` 对应 `BgGloGoodsAdd`。业务模型分别从 `temu/category`、`temu/product`、`temu/image` 导入；鉴权参数和签名由客户端添加。

完整调用示例见 [examples](examples/README.md)，该目录为独立模块，不包含在 SDK 的模块下载包中。
