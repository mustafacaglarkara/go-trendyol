package Utils

import (
	"fmt"
	ProductIntegration "github.com/mustafacaglarkara/go-trendyol/ProductIntegration/Dtos"
	QuestionIntegration "github.com/mustafacaglarkara/go-trendyol/QuestionAnswerIntegration/Dtos"
	"net/url"
	"strconv"
)

func SetBrandsFilterAdres(urlAdress string, brandsFilter ProductIntegration.BrandsFilter) string {
	v := url.Values{}
	if brandsFilter.Size > 0 {
		v.Set("page", fmt.Sprintf("%d", brandsFilter.Page))
	}
	if brandsFilter.Page > 0 {
		v.Set("size", fmt.Sprintf("%d", brandsFilter.Size))
	}
	urlAdress += v.Encode()
	return urlAdress
}

func SetQuestionFilterAdres(urlAdress string, questionFilter QuestionIntegration.QuestionFilter) string {
	v := url.Values{}
	if len(questionFilter.Barcode) > 0 {
		v.Set("barcode", fmt.Sprintf("%s", questionFilter.Barcode))
	}
	if questionFilter.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", questionFilter.Page))
	}
	if questionFilter.Size > 0 {
		v.Set("size", fmt.Sprintf("%d", questionFilter.Size))
	}
	if questionFilter.EndDate > 0 {
		v.Set("endDate", fmt.Sprintf("%d", questionFilter.EndDate))
	}
	if questionFilter.StartDate > 0 {
		v.Set("startDate", fmt.Sprintf("%d", questionFilter.StartDate))
	}
	if questionFilter.Status != "" {
		switch questionFilter.Status {
		case
			QuestionIntegration.WaitingForAnswer,
			QuestionIntegration.WaitingForApprove,
			QuestionIntegration.Answered,
			QuestionIntegration.Reported,
			QuestionIntegration.Rejected:
			v.Set("dateQueryType", string(questionFilter.Status))
		default:
		}
	}
	if questionFilter.OrderByField != "" {
		switch questionFilter.OrderByField {
		case
			QuestionIntegration.LastModifiedDate,
			QuestionIntegration.CreatedDate:
			v.Set("orderByField", string(questionFilter.OrderByField))
		default:
		}
	}
	if questionFilter.OrderByDirection != "" {
		switch questionFilter.OrderByDirection {
		case
			QuestionIntegration.ASC,
			QuestionIntegration.DESC:
			v.Set("orderByDirection", string(questionFilter.OrderByDirection))
		}
	}
	urlAdress += v.Encode()
	return urlAdress
}

func SetUrlAdres(urlAdress string, productParams ProductIntegration.ProductFilter) string {

	v := url.Values{}
	if productParams.Approved {
		v.Set("approved", fmt.Sprintf("%t", productParams.Approved))
	}
	if len(productParams.Barcode) > 0 {
		v.Set("barcode", productParams.Barcode)
	}

	if productParams.StartDate > 0 {
		v.Set("startDate", strconv.Itoa(productParams.StartDate))
	}
	if productParams.EndDate > 0 {
		v.Set("endDate", strconv.Itoa(productParams.EndDate))
	}
	if productParams.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", productParams.Page))
	}
	if productParams.DateQueryType != "" {
		switch productParams.DateQueryType {
		case ProductIntegration.CreatedDate, ProductIntegration.LastModifiedDate:
			v.Set("dateQueryType", string(productParams.DateQueryType))
		default:
		}
	}
	if productParams.Size > 0 {
		v.Set("size", fmt.Sprintf("%d", productParams.Size))
	}
	if productParams.SupplierId > 0 {
		v.Set("supplierId", fmt.Sprintf("%d", productParams.SupplierId))
	}
	if len(productParams.StockCode) > 0 {
		v.Set("stockCode", productParams.StockCode)
	}
	v.Set("archived", fmt.Sprintf("%t", productParams.Archived))

	if len(productParams.ProductMainId) > 0 {
		v.Set("productMainId", productParams.ProductMainId)
	}
	v.Set("onSale", fmt.Sprintf("%t", productParams.OnSale))
	v.Set("rejected", fmt.Sprintf("%t", productParams.Rejected))
	v.Set("blacklisted", fmt.Sprintf("%t", productParams.BlackListed))
	urlAdress += v.Encode()
	return urlAdress
}
