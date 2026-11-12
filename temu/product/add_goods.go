// 参数来源：docs/API-json/bg.glo.goods.add.{request,response}.parameters.json

package product

import "github.com/devlin229/temu-sdk-go/temu/common"

// 请求只含业务参数；公共鉴权和签名参数由客户端添加。
// 普通字段使用值类型；可选对象及有明确零值、null 或清空语义的字段按需保留指针。
// 可选切片使用 omitzero：nil 省略，非 nil 空切片发送 []。
// 必填标记仅表示字段定义；条件必填及业务有效性仍须按商家模式、类目和站点核对。

// BgGloGoodsAddRequest 是 bg.glo.goods.add 的业务请求。
type BgGloGoodsAddRequest struct {
	ProductSemiManagedReq         *ProductSemiManagedRequest         `json:"productSemiManagedReq,omitempty"`         // 可选；Semi-managed Merchant Information
	CarouselImageURLs             []string                           `json:"carouselImageUrls,omitzero"`              // 可选；Commodity Carousel, not required for apparel category, will be aggregated from skc
	ProductOuterPackageImageReqs  []ProductOuterPackageImageRequest  `json:"productOuterPackageImageReqs,omitzero"`   // 可选；Outer Packaging Image
	CopyFromProductID             int64                              `json:"copyFromProductId,omitempty"`             // 可选；Replicated Source Product ID
	ProductName                   string                             `json:"productName"`                             // 必填；Product Name
	ProductI18nReqs               []ProductI18nRequest               `json:"productI18nReqs,omitzero"`                // 可选；Product Multilingual Information Request
	SellOutProductIDSrc           int64                              `json:"sellOutProductIdSrc,omitempty"`           // 可选；Sold Out Product ID Src
	ProductShipmentReq            *ProductShipmentRequest            `json:"productShipmentReq,omitempty"`            // 可选；Product Delivery Information Request
	Cat8ID                        int64                              `json:"cat8Id"`                                  // 必填；Eight-level Category ID, pass 0 if none
	Cat4ID                        int64                              `json:"cat4Id"`                                  // 必填；Fourth-Level Category ID, pass 0 if none
	Cat6ID                        int64                              `json:"cat6Id"`                                  // 必填；Sixth-Level Category ID, pass 0 if none
	ShowSizeTemplateIDs           []int64                            `json:"showSizeTemplateIds,omitzero"`            // 可选；Key Display Size Table Template ID List
	Cat2ID                        int64                              `json:"cat2Id"`                                  // 必填；Secondary Category ID, pass 0 if none
	ProductSpecPropertyReqs       []ProductSpecPropertyRequest       `json:"productSpecPropertyReqs"`                 // 必填；Product Specification Attribute
	VehicleLibraryRelationReqList []VehicleLibraryRelationRequest    `json:"vehicleLibraryRelationReqList,omitzero"`  // 可选；Vehicle Model Library Configuration
	ProductCarouseVideoReqList    []ProductCarouseVideoRequest       `json:"productCarouseVideoReqList,omitzero"`     // 可选；Carousel Video
	GoodsLayerDecorationReqs      []GoodsLayerDecorationRequest      `json:"goodsLayerDecorationReqs,omitzero"`       // 可选；Product Details Decoration
	ProductPropertyReqs           []ProductPropertyRequest           `json:"productPropertyReqs"`                     // 必填；Product Attributes
	ProductGuideFileReqs          []ProductGuideFileRequest          `json:"productGuideFileReqs,omitzero"`           // 可选；Product Manual File Multilingual
	MaterialMultiLanguages        []string                           `json:"materialMultiLanguages,omitzero"`         // 可选；Image Multilingual List
	ProductWarehouseRouteReq      *ProductWarehouseRouteRequest      `json:"productWarehouseRouteReq,omitempty"`      // 可选；Commodity Warehouse Routing Request
	GoodsModelReqs                []GoodsModelRequest                `json:"goodsModelReqs,omitzero"`                 // 可选；Product Model List Request
	SizeTemplateID                int64                              `json:"sizeTemplateId,omitempty"`                // 可选；Size Chart Template ID
	ProductOuterPackageReq        *ProductOuterPackageRequest        `json:"productOuterPackageReq,omitempty"`        // 可选；Product Outer Packaging Information
	SourceInvitationID            int64                              `json:"sourceInvitationId,omitempty"`            // 可选；Source Invitation ID
	SensitiveTransNormalFileReqs  []SensitiveTransNormalFileRequest  `json:"sensitiveTransNormalFileReqs,omitzero"`   // 可选；Sensitive Product Conversion to General Proof Document List
	Cat7ID                        int64                              `json:"cat7Id"`                                  // 必填；Seventh-Level Category ID, pass 0 if none
	SellOutProductID              string                             `json:"sellOutProductId,omitempty"`              // 可选；Sold Out Product ID
	Cat9ID                        int64                              `json:"cat9Id"`                                  // 必填；9th Level Category ID, pass 0 if none
	Cat3ID                        int64                              `json:"cat3Id"`                                  // 必填；Third-Level Category ID, pass 0 if none
	ProductDetailVideoReqList     []ProductDetailVideoRequest        `json:"productDetailVideoReqList,omitzero"`      // 可选；Detail Video
	Cat5ID                        int64                              `json:"cat5Id"`                                  // 必填；Fifth-Level Category ID, pass 0 if none
	Cat1ID                        int64                              `json:"cat1Id"`                                  // 必填；First-Level Category ID
	CarouselImageI18nReqs         []CarouselImageI18nRequest         `json:"carouselImageI18nReqs,omitzero"`          // 可选；Product Carousel Multi-Language Info Request
	SizeTemplateIDs               []int64                            `json:"sizeTemplateIds,omitzero"`                // 可选；Size Chart Template ID List
	ProductWhExtAttrReq           *ProductWhExtAttrRequest           `json:"productWhExtAttrReq,omitempty"`           // 可选；Commodity Warehouse and Supply Chain Extension Properties Request
	ProductSKCReqs                []ProductSKCRequest                `json:"productSkcReqs"`                          // 必填；Product SKC List
	ProductSaleExtAttrReq         *ProductSaleExtAttrRequest         `json:"productSaleExtAttrReq,omitempty"`         // 可选；Product Sales Side Extended Attribute Request
	PersonalizationSwitch         *int64                             `json:"personalizationSwitch,omitempty"`         // 可选；Whether to Support Customized Template, 0: Not Supported, 1: Supported
	ProductCustomReq              *ProductCustomRequest              `json:"productCustomReq,omitempty"`              // 可选；Commodity Customs Information
	Cat10ID                       int64                              `json:"cat10Id"`                                 // 必填；Tenth-Level Category ID, pass 0 if none
	MaterialImgURL                string                             `json:"materialImgUrl"`                          // 必填；Material Image
	ProductComplianceStatementReq *ProductComplianceStatementRequest `json:"productComplianceStatementReq,omitempty"` // 可选；Compliance Signing Agreement
}

// ProductSemiManagedRequest Semi-managed Merchant Information
type ProductSemiManagedRequest struct {
	SemiLanguageStrategy int64   `json:"semiLanguageStrategy,omitempty"` // 可选；Semi-Managed - Material Language Strategy
	BindSiteIDs          []int64 `json:"bindSiteIds"`                    // 必填；Bound Site List
	SemiManagedSiteMode  int64   `json:"semiManagedSiteMode,omitempty"`  // 可选；Semi-managed Site Sales Mode
}

// ProductOuterPackageImageRequest 对应 bg.glo.goods.add.productOuterPackageImageReqs[] 的对象字段。
type ProductOuterPackageImageRequest struct {
	ImageURL string `json:"imageUrl"` // 必填；Image Link
}

// ProductI18nRequest 对应 bg.glo.goods.add.productI18nReqs[] 的对象字段。
type ProductI18nRequest struct {
	Language    string `json:"language"`    // 必填；Language Code
	ProductName string `json:"productName"` // 必填；Product Name
}

// ProductShipmentRequest Product Delivery Information Request
type ProductShipmentRequest struct {
	FreightTemplateID   string `json:"freightTemplateId"`   // 必填；Shipping Template ID
	ShipmentLimitSecond int64  `json:"shipmentLimitSecond"` // 必填；Promised Delivery Time (unit: s). Basic options: 86400, 172800, 259200 (available only for custom items).
}

// ProductSpecPropertyRequest 对应 bg.glo.goods.add.productSpecPropertyReqs[] 的对象字段。
type ProductSpecPropertyRequest struct {
	VID              int64  `json:"vid"`                        // 必填；Basic Attribute Value ID, pass 0 if none
	SpecID           int64  `json:"specId"`                     // 必填；Specification ID
	ValueGroupID     int64  `json:"valueGroupId"`               // 必填；Attribute Value Group ID, 0 if none
	ParentSpecID     int64  `json:"parentSpecId"`               // 必填；Parent Specification ID
	ValueGroupName   string `json:"valueGroupName"`             // 必填；Attribute Group Name, pass empty string if none
	ValueUnit        string `json:"valueUnit"`                  // 必填；Unit of Attribute Value, empty string if not available
	PID              int64  `json:"pid"`                        // 必填；Attribute ID
	TemplatePID      int64  `json:"templatePid"`                // 必填；Template Attribute ID
	NumberInputValue string `json:"numberInputValue,omitempty"` // 可选；Numerical Input
	PropValue        string `json:"propValue"`                  // 必填；Basic Property Value
	PropName         string `json:"propName"`                   // 必填；Reference Property Name
	RefPID           int64  `json:"refPid"`                     // 必填；Reference Property ID
}

// VehicleLibraryRelationRequest 对应 bg.glo.goods.add.vehicleLibraryRelationReqList[] 的对象字段。
type VehicleLibraryRelationRequest struct {
	ProductPropValueDependencyReqList []ProductPropValueDependencyRequest `json:"productPropValueDependencyReqList,omitzero"` // 可选；Attribute Value Dependency Configuration
	VehicleLibraryID                  int64                               `json:"vehicleLibraryId"`                           // 必填；Vehicle Model ID
}

// ProductPropValueDependencyRequest 对应 bg.glo.goods.add.vehicleLibraryRelationReqList[].productPropValueDependencyReqList[] 的对象字段。
type ProductPropValueDependencyRequest struct {
	PropertyValueDependencyID5  int64 `json:"propertyValueDependencyId5,omitempty"`  // 可选；Attribute Value Depends on id5
	PropertyValueDependencyID4  int64 `json:"propertyValueDependencyId4,omitempty"`  // 可选；Attribute Value Depends on id4
	PropertyValueDependencyID3  int64 `json:"propertyValueDependencyId3,omitempty"`  // 可选；Attribute Value Dependent on id3
	PropertyValueDependencyID2  int64 `json:"propertyValueDependencyId2,omitempty"`  // 可选；Attribute Value Depends on id2
	PropertyValueDependencyID9  int64 `json:"propertyValueDependencyId9,omitempty"`  // 可选；Attribute Value Depends on id9
	PropertyValueDependencyID8  int64 `json:"propertyValueDependencyId8,omitempty"`  // 可选；Attribute Value Depends on id8
	PropertyValueDependencyID7  int64 `json:"propertyValueDependencyId7,omitempty"`  // 可选；Attribute Value Depends on id7
	PropertyValueDependencyID6  int64 `json:"propertyValueDependencyId6,omitempty"`  // 可选；Attribute Value Depends on id6
	PropertyValueDependencyID10 int64 `json:"propertyValueDependencyId10,omitempty"` // 可选；Attribute Value Depends on id10
	PropertyValueDependencyID1  int64 `json:"propertyValueDependencyId1,omitempty"`  // 可选；Attribute Value Dependent ID 1
}

// ProductCarouseVideoRequest 对应 bg.glo.goods.add.productCarouseVideoReqList[] 的对象字段。
type ProductCarouseVideoRequest struct {
	VID      string `json:"vid"`      // 必填；Video VID
	CoverURL string `json:"coverUrl"` // 必填；Video Cover Image (B-side stores the first frame image)
	VideoURL string `json:"videoUrl"` // 必填；Video URL
	Width    int64  `json:"width"`    // 必填；Video Width
	Height   int64  `json:"height"`   // 必填；Video Height
}

// GoodsLayerDecorationRequest 对应 bg.glo.goods.add.goodsLayerDecorationReqs[] 的对象字段。
type GoodsLayerDecorationRequest struct {
	FloorID     *int64    `json:"floorId,omitempty"` // 可选；Floor ID, null: Create new, otherwise Update
	GoodsID     int64     `json:"goodsId,omitempty"` // 可选；Product ID
	Lang        string    `json:"lang"`              // 必填；Language Type
	Type        string    `json:"type"`              // 必填；Component Type, image - image, text - text
	Priority    int64     `json:"priority"`          // 必填；Floor Sorting
	ContentList []Content `json:"contentList"`       // 必填；Floor Content
	Key         string    `json:"key"`               // 必填；Floor Type Key, currently defaults to 'DecImage'
}

// Content 对应 bg.glo.goods.add.goodsLayerDecorationReqs[].contentList[] 的对象字段。
type Content struct {
	ImgURL            string             `json:"imgUrl,omitempty"`            // 可选；Image Address--General
	TextModuleDetails *TextModuleDetails `json:"textModuleDetails,omitempty"` // 可选；Text Module Details
	Width             int64              `json:"width,omitempty"`             // 可选；Image Width--General
	Text              string             `json:"text,omitempty"`              // 可选；Text Information--Text Module
	Height            int64              `json:"height,omitempty"`            // 可选；Image Height--General
}

// TextModuleDetails Text Module Details
type TextModuleDetails struct {
	BackgroundColor string `json:"backgroundColor"`      // 必填；Background Color
	FontFamily      int64  `json:"fontFamily,omitempty"` // 可选；Font Type
	FontSize        int64  `json:"fontSize"`             // 必填；Text Module Font Size
	Align           string `json:"align"`                // 必填；Text Alignment, left--Left Aligned；right--Right Aligned；center--Centered；justify--Justified
	FontColor       string `json:"fontColor"`            // 必填；Text Color
}

// ProductPropertyRequest 对应 bg.glo.goods.add.productPropertyReqs[] 的对象字段。
type ProductPropertyRequest struct {
	VID              int64  `json:"vid"`                        // 必填；Basic Attribute Value ID, pass 0 if none
	ValueUnit        string `json:"valueUnit"`                  // 必填；Unit of Attribute Value, empty string if not available
	PID              int64  `json:"pid"`                        // 必填；Attribute ID
	TemplatePID      int64  `json:"templatePid"`                // 必填；Template Attribute ID
	NumberInputValue string `json:"numberInputValue,omitempty"` // 可选；Numerical Input
	PropValue        string `json:"propValue"`                  // 必填；Basic Property Value
	PropName         string `json:"propName"`                   // 必填；Reference Property Name
	RefPID           int64  `json:"refPid"`                     // 必填；Reference Property ID
}

// ProductGuideFileRequest 对应 bg.glo.goods.add.productGuideFileReqs[] 的对象字段。
type ProductGuideFileRequest struct {
	FileName      string   `json:"fileName"`      // 必填；File Name
	PdfMaterialID int64    `json:"pdfMaterialId"` // 必填；PDF File ID
	Languages     []string `json:"languages"`     // 必填；Language
}

// ProductWarehouseRouteRequest Commodity Warehouse Routing Request
type ProductWarehouseRouteRequest struct {
	TargetRouteList  []TargetRouteList  `json:"targetRouteList"`           // 必填；Target Self-Delivery Site-Warehouse Relationship
	CurrentRouteList []CurrentRouteList `json:"currentRouteList,omitzero"` // 可选；Current Self-Delivery Station-Warehouse Relationship
}

// TargetRouteList 对应 bg.glo.goods.add.productWarehouseRouteReq.targetRouteList[] 的对象字段。
type TargetRouteList struct {
	SiteIDList  []int64 `json:"siteIdList"`  // 必填；Site ID.
	WarehouseID string  `json:"warehouseId"` // 必填；Warehouse ID
}

// CurrentRouteList 对应 bg.glo.goods.add.productWarehouseRouteReq.currentRouteList[] 的对象字段。
type CurrentRouteList struct {
	SiteIDList  []int64 `json:"siteIdList"`  // 必填；Site ID.
	WarehouseID string  `json:"warehouseId"` // 必填；Warehouse ID
}

// GoodsModelRequest 对应 bg.glo.goods.add.goodsModelReqs[] 的对象字段。
type GoodsModelRequest struct {
	ModelProfileURL string `json:"modelProfileUrl"`           // 必填；Model Portrait
	SizeSpecName    string `json:"sizeSpecName"`              // 必填；Trial Size Specification Name
	ModelID         int64  `json:"modelId"`                   // 必填；Model ID, not transmitted for new virtual model scenarios
	SizeSpecID      int64  `json:"sizeSpecId"`                // 必填；Fitting Size Specification ID
	ModelWaist      string `json:"modelWaist,omitempty"`      // 可选；Model Waistline Text
	ModelType       int64  `json:"modelType,omitempty"`       // 可选；Model Type, 1: Garment Model, 2: Shoe Model
	ModelName       string `json:"modelName"`                 // 必填；Model Name
	ModelHeight     string `json:"modelHeight,omitempty"`     // 可选；Model Height Text
	ModelFeature    int64  `json:"modelFeature,omitempty"`    // 可选；Model Properties, 1: Real Model, 2: Virtual Model
	ModelFootWidth  string `json:"modelFootWidth,omitempty"`  // 可选；Model Foot Width Text
	ModelBust       string `json:"modelBust,omitempty"`       // 可选；Model Bust Text
	ModelFootLength string `json:"modelFootLength,omitempty"` // 可选；Model Long Text
	TryOnResult     int64  `json:"tryOnResult,omitempty"`     // 可选；Try-on Experience，TRUE_TO_SIZE(1, "Comfortable")TOO_SMALL(2, "Tight-fitting")TOO_LARGE(3, \"Relaxed\")
	ModelHip        string `json:"modelHip,omitempty"`        // 可选；Model Hip Text
}

// ProductOuterPackageRequest Product Outer Packaging Information
type ProductOuterPackageRequest struct {
	PackageShape int64 `json:"packageShape,omitempty"` // 可选；Outer Packaging Shape
	PackageType  int64 `json:"packageType,omitempty"`  // 可选；Outer Packaging Type
}

// SensitiveTransNormalFileRequest 对应 bg.glo.goods.add.sensitiveTransNormalFileReqs[] 的对象字段。
type SensitiveTransNormalFileRequest struct {
	FileName string `json:"fileName"` // 必填；File Name
	FileURL  string `json:"fileUrl"`  // 必填；File Path
}

// ProductDetailVideoRequest 对应 bg.glo.goods.add.productDetailVideoReqList[] 的对象字段。
type ProductDetailVideoRequest struct {
	VID      string `json:"vid"`      // 必填；Video VID
	CoverURL string `json:"coverUrl"` // 必填；Video Cover Image (B-side stores the first frame image)
	VideoURL string `json:"videoUrl"` // 必填；Video URL
	Width    int64  `json:"width"`    // 必填；Video Width
	Height   int64  `json:"height"`   // 必填；Video Height
}

// CarouselImageI18nRequest 对应 bg.glo.goods.add.carouselImageI18nReqs[] 的对象字段。
type CarouselImageI18nRequest struct {
	ImgURLList *[]string `json:"imgUrlList,omitempty"` // 可选；Image List, empty list implies deletion, null implies no change
	Language   string    `json:"language"`             // 必填；Language
}

// ProductWhExtAttrRequest Commodity Warehouse and Supply Chain Extension Properties Request
type ProductWhExtAttrRequest struct {
	ProductOriginCertFiles []ProductOriginCertFiles `json:"productOriginCertFiles,omitzero"` // 可选；Product Origin Certificate Files
	OuterGoodsURL          string                   `json:"outerGoodsUrl"`                   // 必填；Off-Site Product Link (Pass Empty String as Fallback)
	ProductOrigin          ProductOrigin            `json:"productOrigin"`                   // 必填；Product Origin
}

// ProductOriginCertFiles 对应 bg.glo.goods.add.productWhExtAttrReq.productOriginCertFiles[] 的对象字段。
type ProductOriginCertFiles struct {
	FileName string `json:"fileName"` // 必填；File Name
	FileURL  string `json:"fileUrl"`  // 必填；File Url
}

// ProductOrigin Product Origin
type ProductOrigin struct {
	Region2ID        int64  `json:"region2Id,omitempty"` // 可选；Secondary Region ID
	Region1ShortName string `json:"region1ShortName"`    // 必填；First-Level Region Abbreviation (Two-Character Code)
}

// ProductSKCRequest 对应 bg.glo.goods.add.productSkcReqs[] 的对象字段。
type ProductSKCRequest struct {
	ExtCode                         string                               `json:"extCode"`                                  // 必填；Product SKC External Code, pass empty string if not available
	ProductSKUReqs                  []ProductSKURequest                  `json:"productSkuReqs"`                           // 必填；Product SKU List (up to 10 for Apparel Category)
	MainProductSKUSpecReqs          []MainProductSKUSpecRequest          `json:"mainProductSkuSpecReqs"`                   // 必填；Main Sales Specification List
	PreviewImgURLs                  []string                             `json:"previewImgUrls"`                           // 必填；List of Preview Images, not required for non-apparel categories
	ProductSKCCarouselImageI18nReqs []ProductSKCCarouselImageI18nRequest `json:"productSkcCarouselImageI18nReqs,omitzero"` // 可选；SKC Carousel Multi-Language Information Request
	IsBasePlate                     int64                                `json:"isBasePlate,omitempty"`                    // 可选；Whether Baseplate
	ColorImageURL                   string                               `json:"colorImageUrl,omitempty"`                  // 可选；SKC Color Block Diagram
}

// ProductSKURequest 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[] 的对象字段。
type ProductSKURequest struct {
	CurrencyType                  string                             `json:"currencyType"`                            // 必填；Currency (CNY: Chinese Yuan, USD: US Dollar) (Default: Chinese Yuan)
	ProductSKUMultiPackReq        *ProductSKUMultiPackRequest        `json:"productSkuMultiPackReq,omitempty"`        // 可选；Product Multi-Package Request
	ProductSKUSuggestedPriceReq   *ProductSKUSuggestedPriceRequest   `json:"productSkuSuggestedPriceReq,omitempty"`   // 可选；Product SKU Suggested Price Request
	SiteSupplierPrices            []SiteSupplierPrice                `json:"siteSupplierPrices,omitzero"`             // 可选；Site Supply Price List, for semi_managed merchant scenario only
	SupplierPrice                 int64                              `json:"supplierPrice,omitempty"`                 // 可选；Supply Price, deprecated in semi_managed merchant scenario
	ProductSKUUSSuggestedPriceReq *ProductSKUUSSuggestedPriceRequest `json:"productSkuUsSuggestedPriceReq,omitempty"` // 可选；Product SKU US Suggested Price Request
	ProductSKUStockQuantityReq    *ProductSKUStockQuantityRequest    `json:"productSkuStockQuantityReq,omitempty"`    // 可选；Product SKU Inventory Request
	ExtCode                       string                             `json:"extCode"`                                 // 必填；Product SKC External Code, pass empty string if not available
	ProductSKUThumbURLI18nReqs    []ProductSKUThumbURLI18nRequest    `json:"productSkuThumbUrlI18nReqs,omitzero"`     // 可选；SKU Preview Image Multilingual Information Request
	ProductSKUAccessoriesReq      *ProductSKUAccessoriesRequest      `json:"productSkuAccessoriesReq,omitempty"`      // 可选；Product SKU Accessories Request
	ThumbURL                      string                             `json:"thumbUrl"`                                // 必填；Preview Image
	ProductSKUWhExtAttrReq        ProductSKUWhExtAttrRequest         `json:"productSkuWhExtAttrReq"`                  // 必填；Product SKU Extended Attributes
	ProductSKUSpecReqs            []ProductSKUSpecRequest            `json:"productSkuSpecReqs"`                      // 必填；Product SKU Specification List
}

// ProductSKUMultiPackRequest Product Multi-Package Request
type ProductSKUMultiPackRequest struct {
	NumberOfPieces          int64                        `json:"numberOfPieces,omitempty"`          // 可选；Quantity, default is 1 for single item
	IndividuallyPacked      int64                        `json:"individuallyPacked,omitempty"`      // 可选；Whether to use independent packaging (pass -1 to clear)
	ProductSKUNetContentReq *ProductSKUNetContentRequest `json:"productSkuNetContentReq,omitempty"` // 可选；Net Content Request, passing an empty object indicates clearing
	MixedType               int64                        `json:"mixedType,omitempty"`               // 可选；Mixed set type, 1: different products, 2: same product with different specifications
	TotalNetContent         *TotalNetContent             `json:"totalNetContent,omitempty"`         // 可选；Total Net Content
	PieceNewUnitCode        int64                        `json:"pieceNewUnitCode,omitempty"`        // 可选；Unit, 1: piece
	SKUClassification       int64                        `json:"skuClassification,omitempty"`       // 可选；Sku Category, 1: Single Product, 2: Combination Set, 3: Mixed Suite
	NumberOfPiecesNew       int64                        `json:"numberOfPiecesNew,omitempty"`       // 可选；Total quantity contained
	PieceUnitCode           int64                        `json:"pieceUnitCode,omitempty"`           // 可选；Unit per Item, 1: Piece, 2: Pair, 3: Pack
}

// ProductSKUNetContentRequest Net Content Request, passing an empty object indicates clearing
type ProductSKUNetContentRequest struct {
	NetContentUnitCode int64 `json:"netContentUnitCode,omitempty"` // 可选；Net Content Unit, 1: Fluid Ounce, 2: Milliliter, 3: Gallon, 4: Liter, 5: Gram, 6: Kilogram, 7: Troy Ounce, 8: Pound
	NetContentNumber   int64 `json:"netContentNumber,omitempty"`   // 可选；Net Content Value, the target value needs to be multiplied by 1000 before passing the value. For example, if you are passing "1 L", you need to pass "1000" in "netContentNumber".
}

// TotalNetContent Total Net Content
type TotalNetContent struct {
	NetContentUnitCode int64 `json:"netContentUnitCode,omitempty"` // 可选；Net Content Unit, 1: Fluid Ounce, 2: Milliliter, 3: Gallon, 4: Liter, 5: Gram, 6: Kilogram, 7: Troy Ounce, 8: Pound
	NetContentNumber   int64 `json:"netContentNumber,omitempty"`   // 可选；Net Content Value, the target value needs to be multiplied by 1000 before passing the value. For example, if you are passing "1 L", you need to pass "1000" in "netContentNumber".
}

// ProductSKUSuggestedPriceRequest Product SKU Suggested Price Request
type ProductSKUSuggestedPriceRequest struct {
	SuggestedPriceCurrencyType string `json:"suggestedPriceCurrencyType,omitempty"` // 可选；Recommended Price Currency
	SuggestedPrice             int64  `json:"suggestedPrice,omitempty"`             // 可选；Suggested Price
	SpecialSuggestedPrice      string `json:"specialSuggestedPrice,omitempty"`      // 可选；Special Suggested Price
}

// SiteSupplierPrice 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[].siteSupplierPrices[] 的对象字段。
type SiteSupplierPrice struct {
	SiteID        int64 `json:"siteId"`        // 必填；Declared Price Site ID
	SupplierPrice int64 `json:"supplierPrice"` // 必填；Site Declared Price, Unit: RMB: Fen, USD: Cent
}

// ProductSKUUSSuggestedPriceRequest Product SKU US Suggested Price Request
type ProductSKUUSSuggestedPriceRequest struct {
	SuggestedPriceCurrencyType string `json:"suggestedPriceCurrencyType,omitempty"` // 可选；Recommended Price Currency
	SuggestedPrice             int64  `json:"suggestedPrice,omitempty"`             // 可选；Suggested Price
	SpecialSuggestedPrice      string `json:"specialSuggestedPrice,omitempty"`      // 可选；Special Suggested Price
}

// ProductSKUStockQuantityRequest Product SKU Inventory Request
type ProductSKUStockQuantityRequest struct {
	WarehouseStockQuantityReqs []WarehouseStockQuantityRequest `json:"warehouseStockQuantityReqs"` // 必填；Outbound Warehouse Inventory Request List
}

// WarehouseStockQuantityRequest 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[].productSkuStockQuantityReq.warehouseStockQuantityReqs[] 的对象字段。
type WarehouseStockQuantityRequest struct {
	TargetStockAvailable  int64  `json:"targetStockAvailable"`            // 必填；Target Inventory
	WarehouseID           string `json:"warehouseId"`                     // 必填；Warehouse ID
	CurrentStockAvailable *int64 `json:"currentStockAvailable,omitempty"` // 可选；Current Stock
}

// ProductSKUThumbURLI18nRequest 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[].productSkuThumbUrlI18nReqs[] 的对象字段。
type ProductSKUThumbURLI18nRequest struct {
	ImgURLList *[]string `json:"imgUrlList,omitempty"` // 可选；Image List, empty list implies deletion, null implies no change
	Language   string    `json:"language"`             // 必填；Language
}

// ProductSKUAccessoriesRequest Product SKU Accessories Request
type ProductSKUAccessoriesRequest struct {
	ProductSKUAccessories []ProductSKUAccessory `json:"productSkuAccessories"` // 必填；Accessories List
}

// ProductSKUAccessory 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[].productSkuAccessoriesReq.productSkuAccessories[] 的对象字段。
type ProductSKUAccessory struct {
	VID      int64 `json:"vid"`      // 必填；Accessory Property Value ID
	Num      int64 `json:"num"`      // 必填；Accessory Quantity
	UnitCode int64 `json:"unitCode"` // 必填；Unit Code
}

// ProductSKUWhExtAttrRequest Product SKU Extended Attributes
type ProductSKUWhExtAttrRequest struct {
	ProductSKUWeightReq         ProductSKUWeightRequest          `json:"productSkuWeightReq"`                   // 必填；Product SKU Weight
	ProductSKUSameReferPriceReq *ProductSKUSameReferPriceRequest `json:"productSkuSameReferPriceReq,omitempty"` // 可选；Same Style Reference
	ProductSKUSensitiveLimitReq ProductSKUSensitiveLimitRequest  `json:"productSkuSensitiveLimitReq"`           // 必填；Product SKU Sensitive Attribute Restriction Request
	ProductSKUVolumeReq         ProductSKUVolumeRequest          `json:"productSkuVolumeReq"`                   // 必填；Product SKU Volume
	ProductSKUNetWeightReq      *ProductSKUNetWeightRequest      `json:"productSkuNetWeightReq,omitempty"`      // 可选；Product SKU Net Weight
	ProductSKUBarCodeReqs       []ProductSKUBarCodeRequest       `json:"productSkuBarCodeReqs,omitzero"`        // 可选；Product SKU Barcode
	ProductSKUSensitiveAttrReq  ProductSKUSensitiveAttrRequest   `json:"productSkuSensitiveAttrReq"`            // 必填；Product SKU Sensitive Attribute Request
}

// ProductSKUWeightRequest Product SKU Weight
type ProductSKUWeightRequest struct {
	InputUnit  string `json:"inputUnit,omitempty"`  // 可选；Input Unit
	InputValue string `json:"inputValue,omitempty"` // 可选；Input Weight Value
	Value      int64  `json:"value"`                // 必填；Weight Value, unit mg
}

// ProductSKUSameReferPriceRequest Same Style Reference
type ProductSKUSameReferPriceRequest struct {
	URL string `json:"url,omitempty"` // 可选；Same Style URL
}

// ProductSKUSensitiveLimitRequest Product SKU Sensitive Attribute Restriction Request
type ProductSKUSensitiveLimitRequest struct {
	MaxBatteryCapacityHp int64          `json:"maxBatteryCapacityHp,omitempty"` // 可选；Maximum Battery Capacity (mWh)
	MaxBatteryCapacity   int64          `json:"maxBatteryCapacity,omitempty"`   // 可选；Maximum Battery Capacity (Wh) (Prioritizes maxBatteryCapacityHp)
	MaxLiquidCapacity    int64          `json:"maxLiquidCapacity,omitempty"`    // 可选；Maximum Liquid Capacity (mL) (Prefer using maxLiquidCapacityHp)
	MaxLiquidCapacityHp  int64          `json:"maxLiquidCapacityHp,omitempty"`  // 可选；Maximum Liquid Capacity (μL)
	MaxKnifeLength       int64          `json:"maxKnifeLength,omitempty"`       // 可选；Maximum Tool Length (mm) (Prefer using maxKnifeLengthHp)
	MaxKnifeLengthHp     int64          `json:"maxKnifeLengthHp,omitempty"`     // 可选；Maximum Tool Length (μm)
	KnifeTipAngle        *KnifeTipAngle `json:"knifeTipAngle,omitempty"`        // 可选；Blade Angle
}

// KnifeTipAngle Blade Angle
type KnifeTipAngle struct {
	Degrees int64 `json:"degrees"` // 必填；Degree
}

// ProductSKUVolumeRequest Product SKU Volume
type ProductSKUVolumeRequest struct {
	InputUnit   string `json:"inputUnit,omitempty"`   // 可选；Input Unit
	Len         int64  `json:"len"`                   // 必填；Shortest Side, unit mm
	InputLen    string `json:"inputLen,omitempty"`    // 可选；Length of the Longest Input Side
	InputHeight string `json:"inputHeight,omitempty"` // 可选；Length of the Shortest Input Side
	Width       int64  `json:"width"`                 // 必填；Secondary Length, unit mm
	InputWidth  string `json:"inputWidth,omitempty"`  // 可选；Input Secondary Length
	Height      int64  `json:"height"`                // 必填；Shortest Side, unit mm
}

// ProductSKUNetWeightRequest Product SKU Net Weight
type ProductSKUNetWeightRequest struct {
	InputUnit  string `json:"inputUnit,omitempty"`  // 可选；Input Net Weight Unit
	InputValue string `json:"inputValue,omitempty"` // 可选；Input Net Weight Value
	Value      int64  `json:"value,omitempty"`      // 可选；Net Weight Value, unit mg
}

// ProductSKUBarCodeRequest 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[].productSkuWhExtAttrReq.productSkuBarCodeReqs[] 的对象字段。
type ProductSKUBarCodeRequest struct {
	Code     string `json:"code,omitempty"`     // 可选；Barcode
	CodeType int64  `json:"codeType,omitempty"` // 可选；Barcode Type (1: EAN, 2: UPC, 3: ISBN)
}

// ProductSKUSensitiveAttrRequest Product SKU Sensitive Attribute Request
type ProductSKUSensitiveAttrRequest struct {
	SensitiveTypes []int64 `json:"sensitiveTypes,omitzero"` // 可选；Sensitive Type PURE_ELECTRIC(1, "Pure Electric"), INTERNAL_ELECTRIC(2, "Internal Electric"), MAGNETISM(3, "Magnetism"), LIQUID(4, "Liquid"), POWDER(5, "Powder"), PASTE(6, "Paste"), CUTTER(7, "Tool")
	IsSensitive    *int64  `json:"isSensitive,omitempty"`   // 可选；Whether Sensitive Attribute, 0: Non-Sensitive, 1: Sensitive
	SensitiveList  []int64 `json:"sensitiveList,omitzero"`  // 可选；Sensitive Type PURE_ELECTRIC(110001,"Pure Electric"), INTERNAL_ELECTRIC(120001, "Internal Electric"), MAGNETISM(130001, "Magnetism"), LIQUID(140001, "Liquid"), POWDER(150001, "Powder"), PASTE(160001, "Ointment"), CUTTER(170001, "Tool")
}

// ProductSKUSpecRequest 对应 bg.glo.goods.add.productSkcReqs[].productSkuReqs[].productSkuSpecReqs[] 的对象字段。
type ProductSKUSpecRequest struct {
	SpecID         int64  `json:"specId"`         // 必填；Specification ID
	ParentSpecName string `json:"parentSpecName"` // 必填；Parent Specification Name
	ParentSpecID   int64  `json:"parentSpecId"`   // 必填；Parent Specification ID
	SpecName       string `json:"specName"`       // 必填；Specification Name
}

// MainProductSKUSpecRequest 对应 bg.glo.goods.add.productSkcReqs[].mainProductSkuSpecReqs[] 的对象字段。
type MainProductSKUSpecRequest struct {
	SpecID         int64  `json:"specId"`         // 必填；Specification ID
	ParentSpecName string `json:"parentSpecName"` // 必填；Parent Specification Name
	ParentSpecID   int64  `json:"parentSpecId"`   // 必填；Parent Specification ID
	SpecName       string `json:"specName"`       // 必填；Specification Name
}

// ProductSKCCarouselImageI18nRequest 对应 bg.glo.goods.add.productSkcReqs[].productSkcCarouselImageI18nReqs[] 的对象字段。
type ProductSKCCarouselImageI18nRequest struct {
	ImgURLList *[]string `json:"imgUrlList,omitempty"` // 可选；Image List, empty list implies deletion, null implies no change
	Language   string    `json:"language"`             // 必填；Language
}

// ProductSaleExtAttrRequest Product Sales Side Extended Attribute Request
type ProductSaleExtAttrRequest struct {
	InventoryRegion         int64                        `json:"inventoryRegion,omitempty"`         // 可选；Inventory Area
	ProductSecondHandReq    *ProductSecondHandRequest    `json:"productSecondHandReq,omitempty"`    // 可选；Second-hand Goods Information
	DiscreetShipping        *bool                        `json:"discreetShipping,omitempty"`        // 可选；This parameter is required for adult products (category ID: 16571). Setting this service will display the corresponding logo on the consumer order page, which can improve payment conversion rates. Merchants cannot display the details of the purchased items on the shipping packaging. If a consumer complains about failure to maintain confidentiality and the complaint is verified, the merchant may be held liable for after-sales service.
	CustomizedTechnologyReq *CustomizedTechnologyRequest `json:"customizedTechnologyReq,omitempty"` // 可选；Customized Craft Request
	ManufacturingLocation   *ManufacturingLocation       `json:"manufacturingLocation,omitempty"`   // 可选；Manufacturing Location
	ProductNoChargerReq     *ProductNoChargerRequest     `json:"productNoChargerReq,omitempty"`     // 可选；Product No Charger Version Information (empty list must be sent to clear from has to none)
	IpCodes                 []string                     `json:"ipCodes,omitzero"`                  // 可选；ip code list, If ipBizType=1 (bg.glo.goods.brand.get), ipCodes is required.
}

// ProductSecondHandRequest Second-hand Goods Information
type ProductSecondHandRequest struct {
	IsSecondHand    *bool `json:"isSecondHand,omitempty"`    // 可选；Whether Second-hand
	SecondHandLevel int64 `json:"secondHandLevel,omitempty"` // 可选；Second-hand Grade (Condition)
}

// CustomizedTechnologyRequest Customized Craft Request
type CustomizedTechnologyRequest struct {
	TwiceType      []int64 `json:"twiceType,omitzero"` // 可选；Secondary Process
	FirstType      int64   `json:"firstType"`          // 必填；Primary Process
	TechnologyType int64   `json:"technologyType"`     // 必填；Craft Type
}

// ManufacturingLocation Manufacturing Location
type ManufacturingLocation struct {
	Region1ShortNames []string `json:"region1ShortNames,omitzero"` // 可选；List of First-Level Region Two-Letter Short Names
}

// ProductNoChargerRequest Product No Charger Version Information (empty list must be sent to clear from has to none)
type ProductNoChargerRequest struct {
	NoChargerProductIDs []int64 `json:"noChargerProductIds"` // 必填；No Charger Version Product ID (pass empty list to clear)
}

// ProductCustomRequest Commodity Customs Information
type ProductCustomRequest struct {
	GoodsLabelName   string `json:"goodsLabelName,omitempty"` // 可选；Product Tag
	IsRecommendedTag bool   `json:"isRecommendedTag"`         // 必填；Whether to select recommended tags
}

// ProductComplianceStatementRequest Compliance Signing Agreement
type ProductComplianceStatementRequest struct {
	ProtocolVersion string `json:"protocolVersion"` // 必填；Protocol Version Number
	ProtocolURL     string `json:"protocolUrl"`     // 必填；Protocol Link
}

// BgGloGoodsAddRespose 是 bg.glo.goods.add 的完整响应。
type BgGloGoodsAddRespose struct {
	common.Response
	Result *Result `json:"result"` // result
}

// Result result
type Result struct {
	ProductSKUList []ProductSKU `json:"productSkuList"` // SKU List
	ProductID      int64        `json:"productId"`      // Product ID
	ProductSKCList []ProductSKC `json:"productSkcList"` // SKC List
}

// ProductSKU 对应 bg.glo.goods.add.result.productSkuList[] 的对象字段。
type ProductSKU struct {
	ProductSKUID int64     `json:"productSkuId"` // sku id
	ExtCode      string    `json:"extCode"`      // Sku External Code
	SKUSpecList  []SKUSpec `json:"skuSpecList"`  // Sku Specification
	ProductSKCID int64     `json:"productSkcId"` // skc id
}

// SKUSpec 对应 bg.glo.goods.add.result.productSkuList[].skuSpecList[] 的对象字段。
type SKUSpec struct {
	SpecID         int64  `json:"specId"`         // Specification ID
	ParentSpecName string `json:"parentSpecName"` // Parent Specification Name
	ParentSpecID   int64  `json:"parentSpecId"`   // Parent Specification ID
	SpecName       string `json:"specName"`       // Specification Name
}

// ProductSKC 对应 bg.glo.goods.add.result.productSkcList[] 的对象字段。
type ProductSKC struct {
	ProductSKCID int64 `json:"productSkcId"` // skc id
}
