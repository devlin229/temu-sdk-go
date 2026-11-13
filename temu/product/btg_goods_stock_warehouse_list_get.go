// 参数来源：docs/Temu_PA_APIs/product/bg.btg.goods.stock.warehouse.list.get/parameters.json
// 页面更新时间：2026-01-15 14:58:57；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// OpenApiUser 用户信息
type OpenApiUser struct {
	SupplierID int64 `json:"supplierId"` // 必填；供应商id
}

// BgBtgGoodsStockWarehouseListGetRequest 是 bg.btg.goods.stock.warehouse.list.get 的业务请求。
type BgBtgGoodsStockWarehouseListGetRequest struct {
	OpenApiUser OpenApiUser `json:"openApiUser"` // 必填；用户信息
	SiteIDList  []int64     `json:"siteIdList"`  // 必填；站点列表.
}

// ValidWarehouse 对应接口中的嵌套对象。
type ValidWarehouse struct {
	WarehouseDisable bool   `json:"warehouseDisable"` // 仓库是否失效
	WarehouseID      string `json:"warehouseId"`      // 仓库id
	WarehouseName    string `json:"warehouseName"`    // 仓库名称
	ManagementType   string `json:"managementType"`   // 仓库类型 0: 三方仓,1:自建仓,2:家庭仓,3:其他(仅适用于9个工作日发货时效的商品)
}

// WarehouseDTO 对应接口中的嵌套对象。
type WarehouseDTO struct {
	ValidWarehouseList []ValidWarehouse `json:"validWarehouseList"` // 可选发货仓列表
	SiteID             int64            `json:"siteId"`             // 站点id
	SiteName           string           `json:"siteName"`           // 站点名称
}

// GoodsStockWarehouseResult result
type GoodsStockWarehouseResult struct {
	WarehouseDTOList []WarehouseDTO `json:"warehouseDTOList"` // 站点可选发货仓列表
}

// BgBtgGoodsStockWarehouseListGetRespose 是 bg.btg.goods.stock.warehouse.list.get 的完整响应。
type BgBtgGoodsStockWarehouseListGetRespose struct {
	common.Response
	Result *GoodsStockWarehouseResult `json:"result"` // result
}
