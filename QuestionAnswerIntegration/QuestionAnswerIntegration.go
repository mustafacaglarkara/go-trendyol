package QuestionAnswerIntegration

import (
	"encoding/json"
	"fmt"
	"github.com/mustafacaglarkara/go-trendyol/QuestionAnswerIntegration/Dtos"
	"github.com/mustafacaglarkara/go-trendyol/Utils"
	Tools "github.com/mustafacaglarkara/tools"
	"github.com/mustafacaglarkara/tools/RestClient"
	"io/ioutil"
	"net/http"
)

type QuestionAnswerIntegration struct {
	SupplierID int
	UserName   string
	Password   string
}

func (q *QuestionAnswerIntegration) GetCustomerQuestions(questionParams Dtos.QuestionFilter) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdressTmp = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/questions/filter?", q.SupplierID)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(q.UserName, q.Password))

	urlAdress := Utils.SetQuestionFilterAdres(urlAdressTmp, questionParams)

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
func (q *QuestionAnswerIntegration) GetCustomerQuestionById(questionId int) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/questions/%d", q.SupplierID, questionId)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(q.UserName, q.Password))

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
func (q *QuestionAnswerIntegration) CreateAnswer(message Dtos.MessageRequest, questionId int) string {
	var client RestClient.Client
	var tools Tools.Tools

	var urlAdress = fmt.Sprintf("https://api.trendyol.com/sapigw/suppliers/%d/questions/%d/answers", q.SupplierID, questionId)
	var autInfo = fmt.Sprintf("Basic %s", tools.BasicAuth(q.UserName, q.Password))

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", autInfo)

	// burada http request gönderimi yapılabilir
	payload, _ := json.Marshal(message.Text)

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
