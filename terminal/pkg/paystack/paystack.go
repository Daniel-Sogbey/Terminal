package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Payment struct {
	Email  string `json:"email"`
	Amount int    `json:"amount"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

func Initialize(p *Payment) (*Response, error) {
	url := "https://api.paystack.co/transaction/initialize"
	client := &http.Client{}

	js, err := json.Marshal(&p)

	if err != nil {
		log.Println("unable to marshal json. ERR: ", err)
		return nil, err
	}

	reqBody := bytes.NewBuffer(js)

	req, err := http.NewRequest(http.MethodPost, url, reqBody)

	req.Header.Set("authorization", "Bearer sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb")
	req.Header.Set("content-type", "application/json")

	if err != nil {
		log.Println("unable to create an http request. ERR : ", err)
		return nil, err
	}

	fmt.Println("REQUEST : ", req)

	resp, err := client.Do(req)

	if err != nil {
		log.Println("unable to perform request. ERR : ", err)
		return nil, err
	}

	fmt.Println("STATUS CODE: ", resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("unable to read response body. ERR : ", err)
		return nil, err
	}

	fmt.Printf("RESPONSE BODY %v\n", string(respBody))

	var response *Response

	err = json.Unmarshal(respBody, &response)

	if err != nil {
		log.Println("unable to unmarshal response data. Err : ", err)
		return nil, err
	}

	return response, nil
}

func VerifyTransaction(reference string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.paystack.co/transaction/verify/%s", reference)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("authorization", "Bearer sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb")

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(respBody, &data)

	return data, nil

}
