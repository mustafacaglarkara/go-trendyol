package Dtos

// ProductFilter
type ProductFilter struct {
	Approved      bool          `url:"approved,omitempty"`
	Barcode       string        `url:"barcode,omitempty"`
	StartDate     int           `url:"startDate,omitempty"`
	EndDate       int           `url:"endDate,omitempty"`
	Page          int           `url:"page,omitempty"`
	DateQueryType DateQueryType `url:"dateQueryType,omitempty"`
	Size          int           `url:"size,omitempty"`
	SupplierId    int64         `url:"supplierId,omitempty"`
	StockCode     string        `url:"stockCode,omitempty"`
	Archived      bool          `url:"archived,omitempty"`
	ProductMainId string        `url:"productMainId,omitempty"`
	OnSale        bool          `url:"onSale,omitempty"`
	Rejected      bool          `url:"rejected,omitempty"`
	BlackListed   bool          `url:"blackListed,omitempty"`
	//BrandIds      []interface{} `url:"brandIds,omitempty"`
}

type DateQueryType string

const (
	CreatedDate      DateQueryType = "CREATED_DATE"
	LastModifiedDate DateQueryType = "LAST_MODIFIED_DATE"
)

// Brands Filter
type BrandsFilter struct {
	Page int `url:"page,omitempty"`
	Size int `url:"size,omitempty"`
}
