package Dtos

// Update İşlemleri İçin Gerekli Olan DTO
type UpdateProductItem struct {
	Barcode           string `json:"barcode"`
	Title             string `json:"title"`
	ProductMainId     string `json:"productMainId"`
	BrandId           int    `json:"brandId"`
	CategoryId        int    `json:"categoryId"`
	StockCode         string `json:"stockCode"`
	DimensionalWeight int    `json:"dimensionalWeight"`
	Description       string `json:"description"`
	DeliveryDuration  int    `json:"deliveryDuration"`
	VatRate           int    `json:"vatRate"`
	DeliveryOption    struct {
		DeliveryDuration int    `json:"deliveryDuration"`
		FastDeliveryType string `json:"fastDeliveryType"`
	} `json:"deliveryOption"`
	Images []struct {
		Url string `json:"url"`
	} `json:"images"`
	Attributes []struct {
		AttributeId          int    `json:"attributeId"`
		AttributeValueId     int    `json:"attributeValueId,omitempty"`
		CustomAttributeValue string `json:"customAttributeValue,omitempty"`
	} `json:"attributes"`
	CargoCompanyId     int `json:"cargoCompanyId"`
	ShipmentAddressId  int `json:"shipmentAddressId"`
	ReturningAddressId int `json:"returningAddressId"`
}
type UpdateProductIntegration struct {
	Items []UpdateProductItem `json:"items"`
}

// Insert İşlemi İçin Gerekli olan DTO
type TransferProductsDto struct {
	Items []struct {
		Barcode           string  `json:"barcode"`
		Title             string  `json:"title"`
		ProductMainID     string  `json:"productMainId"`
		BrandID           int     `json:"brandId"`
		CategoryID        int     `json:"categoryId"`
		Quantity          int     `json:"quantity"`
		StockCode         string  `json:"stockCode"`
		DimensionalWeight int     `json:"dimensionalWeight"`
		Description       string  `json:"description"`
		CurrencyType      string  `json:"currencyType"`
		ListPrice         float64 `json:"listPrice"`
		SalePrice         float64 `json:"salePrice"`
		VatRate           int     `json:"vatRate"`
		CargoCompanyID    int     `json:"cargoCompanyId"`
		DeliveryOption    struct {
			DeliveryDuration int    `json:"deliveryDuration"`
			FastDeliveryType string `json:"fastDeliveryType"`
		} `json:"deliveryOption"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
		Attributes []struct {
			AttributeID          int    `json:"attributeId"`
			AttributeValueID     int    `json:"attributeValueId,omitempty"`
			CustomAttributeValue string `json:"customAttributeValue,omitempty"`
		} `json:"attributes"`
	} `json:"items"`
}

// Fiyat ve Stok Güncellemesi İçin Gerekli Olan DTO
type UpdatePriceAndInventoryDto struct {
	Items []struct {
		Barcode   string  `json:"barcode"`
		Quantity  int     `json:"quantity"`
		SalePrice float64 `json:"salePrice"`
		ListPrice float64 `json:"listPrice"`
	} `json:"items"`
}
