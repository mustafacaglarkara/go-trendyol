# Trendyol Module

## Modüller
- [X] Product Integration
- [X] Question & Answer Integration
- [ ] Claim Integration
- [ ] Order Integration

## Yapılması Gerekenler
- Response DTO modelleri oluşturulmalı
- Detaylı Dökümantasyon hazırlanmalı
- Canlı Ortamda Test Edilmedi


## Fonksiyonların Kullanımı

### Fiyat Güncelleme
```
func main(){
    var envIslemleri EnvIslemleri.EnvIslemleri
	var trendyol ProductIntegration.ProductIntegration
	var product Dtos.UpdatePriceAndInventoryDto

	trendyol.SupplierID = "{supplierID}"
	trendyol.UserName = "{apiKey}"
	trendyol.Password = "{apiSecret}"

	product = Dtos.UpdatePriceAndInventoryDto{
		Items: []struct {
			Barcode   string  `json:"barcode"`
			Quantity  int     `json:"quantity"`
			SalePrice float64 `json:"salePrice"`
			ListPrice float64 `json:"listPrice"`
		}{
			{
				Barcode:   "um-13297169",
				Quantity:  16,
				SalePrice: 153.99,
				ListPrice: 168.99,
			},
		},
	}

	durum := trendyol.UpdateStockAndPrice(product)
	println(durum)
}
```
### Ürün Filtreleme
```
func main() {
	var trendyol ProductIntegration.ProductIntegration

	trendyol.SupplierID = "{supplierID}"
	trendyol.UserName = "{apiKey}"
	trendyol.Password = "{apiSecret}"
	
	var filtering = Dtos.ProductFilter{}
	filtering.Page = 1
	filtering.Approved = true
	durum := trendyol.ProductFiltering(filtering)

	println(durum)
}
```
### Ürün Güncelleme
```
func main() {
	var trendyol ProductIntegration.ProductIntegration

	trendyol.SupplierID = "{supplierID}"
	trendyol.UserName = "{apiKey}"
	trendyol.Password = "{apiSecret}"

	var products Dtos.UpdateProductIntegration
	var product Dtos.UpdateProductItem

	product.Barcode = "Deneme"
	product.Attributes = append(product.Attributes)
	products.Items = append(products.Items, product)
	_ = trendyol.UpdateProducts(products)
}
```
### Batch Kontrol
```
func main() {
	var trendyol ProductIntegration.ProductIntegration

	var res = trendyol.GetBatchRequestResult("22f66726-d712-11ed-a4ac-3a5fa292beed-1681085241")
	println(res)
}
```
### Marka Listeleme
```
func main() {
	var trendyol ProductIntegration.ProductIntegration
	var res = trendyol.GetBrands(Dtos.BrandsFilter{Page: 0, Size: 50})
	fmt.Println(res)
}
```


### Marka Listeleme Ada Göre
```
func main() {
	var trendyol ProductIntegration.ProductIntegration
	var res = trendyol.GetBrandsByBrandsName("ürünmoda")
	
	var brands []Dtos.GetBrands
	
	if err := json.Unmarshal([]byte(res), &brands); err != nil {
		panic(err)
	}
	fmt.Println(brands[0].Name)
	println(res)
}
```
### Kategori Listeleme
```
func main() {
	var trendyol ProductIntegration.ProductIntegration
	var res = trendyol.GetCategories()
	fmt.Println(res)
}
```
### Kategori Attributes
```
func main() {
	var trendyol ProductIntegration.ProductIntegration
	var res = trendyol.GetCategoryAttributes(592)
	fmt.Println(res)
}
```
### Depo Listeleme
```
func GetWarehauseListmain() {
	var trendyol ProductIntegration.ProductIntegration

	trendyol.SupplierID = "{supplierID}"
	trendyol.UserName = "{apiKey}"
	trendyol.Password = "{apiSecret}"

	var res = trendyol.GetWarehouseAdressList()
	println(res)
}
```
### Müşteri Sorularını Listeleme
```
func main() {
	var trendyol QuestionIntegration.QuestionAnswerIntegration
	
	trendyol.SupplierID = "{supplierID}"
	trendyol.UserName = "{apiKey}"
	trendyol.Password = "{apiSecret}"
	
	var res = trendyol.GetCustomerQuestions(QDto.QuestionFilter{
		Page: 0, Size: 50,
	})
	println(res)
}
```