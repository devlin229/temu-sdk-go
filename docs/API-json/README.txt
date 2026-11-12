Temu PA 接口请求与响应 JSON
导出日期：2026-11-12

包含三个接口：
1. bg.glo.goods.list.get 商品列表查询
2. bg.glo.goods.add 商品发布
3. bg.goods.image.upload.global 图片上传

文件说明：
- *.request.json：公共参数和业务参数在同一根对象；保留全部字段，展示请求层级。
- *.response.json：全部响应字段的结构模板，不是真实响应，success=false、errorCode=0 仅为类型占位，无业务含义。
- *.request.parameters.json：请求参数定义，含官方类型、必填标记、说明及 children 嵌套层级。
- *.response.parameters.json：响应参数定义；官方未列响应必填标记，因此不臆造。
- bg.glo.goods.add.full_managed.request.json：额外的全托模板，剔除半托站点、物流、仓库路由、分站点报价和仓库库存字段；使用 supplierPrice。其余可选字段仍按实际类目选用。

重要：这些是完整字段结构模板，不是最小可用请求。所有 0、空字符串、false、示意数组都只是类型占位，不代表默认值、合法枚举或真实数据。正式请求须删除不使用的可选字段，替换 ID、金额、图片、规格、时间戳和授权信息，再对最终请求计算签名。不要直接提交完整模板。
公共 version 在官方表中标必填，但说明为无要求不传；模板按说明省略，定义文件保留原文。
图片 image 参数应为实际 Base64 字符串。
完整定义保留官方全部全托/半托字段，未因为当前使用全托而删减原始参数。
没有调用任何发布或上传接口。

Go 结构体对应文件：
- bg.glo.goods.list.get：temu/product/goods_list_get.go，入口为 GetGoodsListRequest、GetGoodsListResponse。
- bg.glo.goods.add：temu/product/add_goods.go，入口为 AddGoodsRequest、AddGoodsResponse。
- bg.goods.image.upload.global：temu/image/goods_image_upload_global.go，入口为 UploadGoodsImageRequest、UploadGoodsImageResponse。

结构体的请求部分只包含 business_parameters；common_parameters 由现有客户端填入同一 JSON 根对象。响应保留 result、success、errorCode、errorMsg，result 和嵌套对象使用明确类型。
新增发布、上传模型的 INTEGER 使用 int64，必填字段保留零值，可选字段用指针加 omitempty。可选数组的 nil 指针省略字段，指向 nil 切片发送 null，指向空切片发送 []；必填数组不使用 omitempty。必填标记不代替业务校验。
商品列表请求保留原有 PageParams 和可选字段写法；响应中 virtualStock 位于 SKU 摘要，warehouseRegionId1List、longTransport 位于 productShipment。

维护发布、上传模型：先修改对应参数字典，再运行 python3 scripts/generate_api_models.py，不直接修改生成文件。
验证：go test ./... 和 go vet ./...。tests/models/ 对三个接口逐层检查字段和类型，并验证模板编解码及新增请求模型的可选值序列化；本地验证不代表真实接口联调。
