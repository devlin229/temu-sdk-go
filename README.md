# temu-tool

使用 Go 对接 Temu 开放接口，辅助商家执行商品管理等操作。已实现公共客户端和商品发布接口 `bg.glo.goods.add`，仅使用 Go 标准库。当前通过本地模拟测试，尚未进行真实店铺发品联调。

## 目录与职责

```text
.
├── main.go                    # 信号处理与程序入口
├── go.mod
├── .env.example               # 环境变量示例，不会自动加载
├── AGENTS.md                  # 开发约定
├── internal/
│   ├── cli/                   # 命令行解析、配置检查与商品操作
│   ├── config/                # 配置加载及检查
│   └── temu/
│       ├── client.go          # 公共参数、超时与单次 HTTP 调用
│       ├── sign.go            # JSON 编码与签名
│       ├── errors.go          # HTTP 错误
│       └── product/           # 发布请求模型、校验、调用与结果解析
├── tests/                     # 集中维护测试
│   ├── cli/                   # 命令行测试
│   ├── config/                # 配置测试
│   ├── temu/                  # 公共客户端与签名测试
│   ├── product/               # 商品发布测试
│   └── testdata/              # 测试数据
├── examples/product-add.json  # 占位商品业务参数，不能直接用于真实发品
├── scripts/                   # 从文档生成请求模型的维护脚本
└── docs/                      # API 文档与本地使用说明
```

## 运行

使用 `go.mod` 声明的 Go 版本。运行帮助无需配置凭据：

```sh
go run . -h
```

参考 `.env.example`，通过 shell 或部署环境设置 `TEMU_APP_KEY`、`TEMU_APP_SECRET`、`TEMU_ACCESS_TOKEN`。`TEMU_GATEWAY` 默认使用文档中的 PA 网关，`TEMU_TIMEOUT` 默认为 `30s`。程序不自动读取 `.env` 文件。不要提交真实凭据。

```sh
go run . -check-config
```

此命令只检查配置是否齐全和格式是否有效，不访问 Temu，不验证授权权限或凭据真实性，也不会创建商品。

## 商品发布

临时商品查询调试入口为 `go run . -test`，使用环境变量中的凭据调用 `bg.glo.goods.list.get`，stderr 输出完整请求，stdout 输出完整 JSON 响应。该 type 沿用 `main.go` 中的测试定义；项目尚未提供其业务文档，目前仅发送公共参数，实际网关适用性、权限、分页与筛选字段需根据对应文档确认。输出业务响应不表示业务一定成功，需检查平台返回内容。

复制 `examples/product-add.json` 到本地文件，按实际商品和类目模板替换所有示例 ID、规格、图片及价格。文件只包含业务字段，公共鉴权与签名由客户端添加。

```sh
# 离线基础校验，无需凭据
go run . -validate-product product.json

# 使用环境配置实际创建货品，并将返回的 ID/编码映射保存为 JSON
go run . -publish-product product.json > product-result.json
```

仅第二条命令发送商品创建请求。提交成功表示货品信息已创建，不表示审核通过或消费者可购买。超时、响应不完整、结果保存失败时先核对平台记录，不直接重复提交。

发送前会向 stderr 打印请求方法、URL、请求头和完整的已签名 JSON 参数，不脱敏。`> product-result.json` 仍只保存结果；如需保存调试输出，可追加 `2> request-debug.log`。`app_secret` 不在请求体中，也不会输出。

详细说明、Go 调用示例及本地校验边界见 [商品发布使用说明](docs/product-publish.md)。

## 扩展接口

1. 在 `docs/API-docs/` 整理接口文档，确认商家模式、网关、type、权限和参数要求。
2. 商品接口放在 `internal/temu/product/`；其他业务按实际需要增加如 `internal/temu/order/` 的包，不提前创建空模块。
3. 定义明确的请求、响应结构和业务方法，复用 `temu.Client.Do(ctx, apiType, request)`。请求结构只包含业务字段；公共层添加鉴权、时间戳和签名。若接口明确要求版本，可在请求中包含 `version`，公共层默认不添加。
4. `Do` 目前只支持当前文档约定的 POST JSON，返回完整 JSON 响应，业务模块自行解析并检查业务成功标记。不同调用格式需核对文档后扩展。可用 `errors.As` 识别 `*temu.HTTPError`。
5. 上层商家操作流程在需要时另建 `internal/service/`，组合业务接口；不把编排逻辑放入签名或 HTTP 层。
6. 不同网关或不同店铺授权分别创建客户端。公共层不自动重试，发生写操作超时后由业务流程核对执行结果。

当前客户端要求 HTTPS 网关、禁止跟随重定向，并限制响应体为 16 MiB。测试通过自定义 HTTP Transport 模拟响应，不依赖真实凭据或外网。

## 验证

```sh
go test ./...
go vet ./...
go build -o bin/temu-tool .
```

测试统一放在 `tests/` 下按模块分目录，通过业务包的公开入口验证行为；测试数据统一放在 `tests/testdata/`。只运行测试目录也可使用 `go test ./tests/...`。

本地测试不代表已完成真实 Temu 接口联调。类目模板、图片规格、商家模式和店铺权限等动态规则需由平台校验。
