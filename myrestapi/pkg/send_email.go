package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Personalization struct {
	To      []Email `json:"to"`
	Subject string  `json:"subject"`
}

type Email struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type From struct {
	Email string `json:"email"`
	Name  string `json:"Name"`
}

type ReplyTo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type EmailRequest struct {
	Personalization []Personalization `json:"personalizations"`
	Content         []Content         `json:"content"`
	From            From              `json:"from"`
	ReplyTo         ReplyTo           `json:"reply_to"`
}

type EmailClient struct {
	Url    string
	ApiKey string
	Client *http.Client
}

func NewEmailClient() *EmailClient {
	fmt.Println(os.Getenv("SEND_GRID_API_KEY"))
	return &EmailClient{
		Url:    "https://api.sendgrid.com/v3/mail/send",
		ApiKey: os.Getenv("SEND_GRID_API_KEY"),
		Client: &http.Client{},
	}
}

func (ec *EmailClient) SendEmail(message, to string) (int, error) {

	er := &EmailRequest{
		Personalization: []Personalization{
			{
				To: []Email{
					{
						Email: to,
					},
				},
				Subject: message,
			},
		},
		Content: []Content{
			{
				Type:  "text/plain",
				Value: "Hello, From Go Server",
			},
		},
		From: From{
			Email: "gift.alchemy.developer@gmail.com",
			Name:  "Daniel",
		},
		ReplyTo: ReplyTo{
			Email: "gift.alchemy.developer@gmail.com",
			Name:  "Daniel",
		},
	}

	fmt.Println("EMAIL CLIENT", ec)
	fmt.Println("EMAIL REQUEST", er)

	b, err := json.Marshal(er)

	if err != nil {
		return 500, err
	}

	req, err := http.NewRequest(http.MethodPost, ec.Url, bytes.NewBuffer(b))

	if err != nil {
		return 500, errors.New("couldn't create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", ec.ApiKey))

	res, err := ec.Client.Do(req)

	if err != nil {
		return 500, errors.New("couldn't make request to send grid servers")
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error reading response body into []byte : %v", err)
	}

	fmt.Println(string(responseBody))

	if res.StatusCode == http.StatusUnauthorized {
		fmt.Println("Response ", res.StatusCode)
		return res.StatusCode, errors.New("unauthorized request")
	}

	if res.StatusCode != http.StatusAccepted {
		return res.StatusCode, errors.New(err.Error())
	}

	return res.StatusCode, nil

}
