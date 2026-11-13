// 参数来源：docs/Temu_PA_APIs/product/bg.glo.goods.migrate/parameters.json
// 页面更新时间：2025-08-27 16:57:38；导出日期：2026-11-12。
package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 可选切片 nil 时省略，非 nil 空切片发送 []；可选布尔值用指针区分省略与 false。

// GoodsMigrateProductSemiManagedRequest Product Semi Managed Request
type GoodsMigrateProductSemiManagedRequest struct {
	BindSiteIDs         []int64 `json:"bindSiteIds"`                   // 必填；Bound Site List
	SemiManagedSiteMode *int64  `json:"semiManagedSiteMode,omitempty"` // 可选；Semi-managed Site Sales Mode
}

// TargetRoute 对应接口中的嵌套对象。
type TargetRoute struct {
	SiteIDList  []int64 `json:"siteIdList"`  // 必填；Site ID.
	WarehouseID string  `json:"warehouseId"` // 必填；Warehouse ID
}

// GoodsMigrateProductWarehouseRouteRequest Product Warehouse Route Request
type GoodsMigrateProductWarehouseRouteRequest struct {
	TargetRouteList []TargetRoute `json:"targetRouteList"` // 必填；Target Self-Delivery Site-Warehouse Relationship
}

// Spec 对应接口中的嵌套对象。
type Spec struct {
	SpecID         int64  `json:"specId"`         // 必填；Specification ID
	ParentSpecName string `json:"parentSpecName"` // 必填；Parent Specification Name
	ParentSpecID   int64  `json:"parentSpecId"`   // 必填；Parent Specification ID
	SpecName       string `json:"specName"`       // 必填；Specification Name
}

// SiteSupplierPrices 对应接口中的嵌套对象。
type SiteSupplierPrices struct {
	SiteID        int64 `json:"siteId"`        // 必填；Declared Price Site ID
	SupplierPrice int64 `json:"supplierPrice"` // 必填；Site Declared Price, Unit: RMB: Fen, USD: Cent
}

// SKUDetails 对应接口中的嵌套对象。
type SKUDetails struct {
	CurrencyType       string               `json:"currencyType"`       // 必填；Currency Type
	SpecList           []Spec               `json:"specList"`           // 必填；SKU Spec List
	SiteSupplierPrices []SiteSupplierPrices `json:"siteSupplierPrices"` // 必填；Site Supplier Price List
}

// SKCDetails 对应接口中的嵌套对象。
type SKCDetails struct {
	SKUDetails []SKUDetails `json:"skuDetails"` // 必填；Sku Detail List
	SpecList   []Spec       `json:"specList"`   // 必填；SKC Spec List
}

// GoodsMigrateProductShipmentRequest Product Shipment Request
type GoodsMigrateProductShipmentRequest struct {
	FreightTemplateID   string `json:"freightTemplateId"`   // 必填；Shipping Template ID
	ShipmentLimitSecond int64  `json:"shipmentLimitSecond"` // 必填；Promised Delivery Time (unit: s)
}

// Migration 对应接口中的嵌套对象。
type Migration struct {
	ProductSemiManagedReq      GoodsMigrateProductSemiManagedRequest    `json:"productSemiManagedReq"`      // 必填；Product Semi Managed Request
	ProductWarehouseRouteReq   GoodsMigrateProductWarehouseRouteRequest `json:"productWarehouseRouteReq"`   // 必填；Product Warehouse Route Request
	SKCDetails                 []SKCDetails                             `json:"skcDetails"`                 // 必填；Skc Detail List
	SourceProductID            int64                                    `json:"sourceProductId"`            // 必填；Source Product ID
	ProductShipmentReq         GoodsMigrateProductShipmentRequest       `json:"productShipmentReq"`         // 必填；Product Shipment Request
	WarehouseSKUStockAvailable int64                                    `json:"warehouseSkuStockAvailable"` // 必填；Warehouse Sku Stock Available
}

// BgGloGoodsMigrateRequest 是 bg.glo.goods.migrate 的业务请求。
type BgGloGoodsMigrateRequest struct {
	MigrationList []Migration `json:"migrationList"` // 必填；Copy Request List
}

// MigrationRet 对应接口中的嵌套对象。
type MigrationRet struct {
	ProductDraftID  int64 `json:"productDraftId"`  // Product Draft ID
	SourceProductID int64 `json:"sourceProductId"` // Source Product ID`
}

// GoodsMigrateResult result
type GoodsMigrateResult struct {
	MigrationRetList []MigrationRet `json:"migrationRetList"` // Copy Fully Managed Product Result List
}

// BgGloGoodsMigrateRespose 是 bg.glo.goods.migrate 的完整响应。
type BgGloGoodsMigrateRespose struct {
	common.Response
	Result *GoodsMigrateResult `json:"result"` // result
}
