package ProductIntegration

import (
	"encoding/json"
	"fmt"
	ProductDtos "github.com/mustafacaglarkara/go-trendyol/ProductIntegration/Dtos"
	"github.com/mustafacaglarkara/go-trendyol/Utils"
	Tools "github.com/mustafacaglarkara/tools"
	"github.com/mustafacaglarkara/tools/RestClient"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ProductIntegration struct {
	SupplierID int
	UserName   string
	Password   string
}

/*
Yapılacaklar
Get Categories ve diğerleri için geri dönüş dtos ları yazılacak
*/
func (p *ProductIntegration) ProductFiltering(productParams ProductDtos.ProductFilter) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdressTmp = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/products?", p.SupplierID)
	urlAdress := Utils.SetUrlAdres(urlAdressTmp, productParams)

	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(p.UserName, p.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, headers, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)

}
func (p *ProductIntegration) UpdateStockAndPrice(productDto ProductDtos.UpdatePriceAndInventoryDto) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/products/price-and-inventory", p.SupplierID)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(p.UserName, p.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// burada http request gönderimi yapılabilir
	payload, _ := json.Marshal(productDto)

	// İstek Gönderiliyor
	resp, err := client.Post(urlAdress, headers, "", payload)
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)

}
func (p *ProductIntegration) TransferProduct(productDto ProductDtos.TransferProductsDto) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/v2/products", p.SupplierID)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(p.UserName, p.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// burada http request gönderimi yapılabilir
	payload, _ := json.Marshal(productDto)

	// İstek Gönderiliyor
	resp, err := client.Post(urlAdress, headers, "", payload)
	if err != nil {
		fmt.Println("Post isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)

}
func (p *ProductIntegration) UpdateProducts(productDto ProductDtos.UpdateProductIntegration) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/v2/products", p.SupplierID)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(p.UserName, p.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// burada http request gönderimi yapılabilir
	payload, _ := json.Marshal(productDto)

	// İstek Gönderiliyor
	resp, err := client.Put(urlAdress, headers, "", payload)
	if err != nil {
		fmt.Println("Put isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)
}
func (p *ProductIntegration) GetBatchRequestResult(batchId string) string {

	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/products/batch-requests/%s", p.SupplierID, batchId)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(p.UserName, p.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, headers, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)

}
func (p *ProductIntegration) GetBrands(brandsFilter ProductDtos.BrandsFilter) string {
	var client RestClient.Client
	var urlAdress = Utils.SetBrandsFilterAdres("https://api.trendyol.com/sapigw/brands?", brandsFilter)

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, nil, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)

}
func (p *ProductIntegration) GetBrandsByBrandsName(brandName string) string {
	encodedBrandName := url.QueryEscape(brandName)
	var client RestClient.Client
	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/brands/by-name?name=%s", encodedBrandName)

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, nil, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)
}
func (p *ProductIntegration) GetCategories() string {
	var client RestClient.Client
	var urlAdress = "https://api.trendyol.com/sapigw/product-categories"

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, nil, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()

	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)
}
func (p *ProductIntegration) GetCategoryAttributes(catId int) string {
	var client RestClient.Client
	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/product-categories/%d/attributes", catId)

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, nil, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()

	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)
}
func (p *ProductIntegration) GetWarehouseAdressList() string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/addresses", p.SupplierID)

	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(p.UserName, p.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// İstek Gönderiliyor
	resp, err := client.Get(urlAdress, headers, "")
	if err != nil {
		fmt.Println("Get isteği başarısız oldu: ", err)
		return ""
	}
	defer resp.Body.Close()
	// Cevap kodunu ve body'sini yazdırıyoruz.
	fmt.Println("Get isteği başarılı oldu: ", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body okuma hatası: ", err)
		return ""
	}
	return string(body)
}
