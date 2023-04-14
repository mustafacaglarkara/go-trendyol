package Dtos

type DeliveryOption struct {
	DeliveryDuration int    `json:"deliveryDuration"`
	FastDeliveryType string `json:"fastDeliveryType"`
}

type Image struct {
	URL string `json:"url"`
}

type Attribute struct {
	AttributeID          int    `json:"attributeId"`
	AttributeValueID     int    `json:"attributeValueId,omitempty"`
	CustomAttributeValue string `json:"customAttributeValue,omitempty"`
}

type Item struct {
	Barcode           string         `json:"barcode"`
	Title             string         `json:"title"`
	ProductMainID     string         `json:"productMainId"`
	BrandID           int            `json:"brandId"`
	CategoryID        int            `json:"categoryId"`
	Quantity          int            `json:"quantity"`
	StockCode         string         `json:"stockCode"`
	DimensionalWeight int            `json:"dimensionalWeight"`
	Description       string         `json:"description"`
	CurrencyType      string         `json:"currencyType"`
	ListPrice         float64        `json:"listPrice"`
	SalePrice         float64        `json:"salePrice"`
	VATRate           int            `json:"vatRate"`
	CargoCompanyID    int            `json:"cargoCompanyId"`
	DeliveryOption    DeliveryOption `json:"deliveryOption"`
	Images            []Image        `json:"images"`
	Attributes        []Attribute    `json:"attributes"`
}

type JSONData struct {
	Items []Item `json:"items"`
}
