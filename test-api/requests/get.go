package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Status struct {
	verified  bool
	sentCount int
}

type Fact struct {
	response []Response
}

// Fact response from the cat facts api
type Response struct {
	Status struct {
		Verified  bool `json:"verified"`
		SentCount int  `json:"sentCount"`
	} `json:"status"`
	ID        string `json:"_id"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Version   int    `json:"__v"`
	Source    string `json:"source"`
	UpdatedAt string `json:"updatedAt"`
	Type      string `json:"type"`
	CreatedAt string `json:"createdAt"`
	Deleted   bool   `json:"deleted"`
	Used      bool   `json:"used"`
}

type Client struct {
	ApiKey string
	Client *http.Client
}

func getUrl(service string) string {
	if service == "fact" {
		return "https://cat-fact.herokuapp.com/facts"
	}
	return ""
}

func formatResponse(response *http.Response) (Fact, error) {
	var res []Response
	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&res); err != nil {
		return Fact{}, err
	}

	data := []Response{}

	for _, fact := range res {
		fmt.Println(fact.ID)
		response := Response{
			Status:    fact.Status,
			ID:        fact.ID,
			User:      fact.User,
			Text:      fact.Text,
			Version:   fact.Version,
			Source:    fact.Source,
			UpdatedAt: fact.UpdatedAt,
			Type:      fact.Type,
			CreatedAt: fact.CreatedAt,
			Deleted:   fact.Deleted,
			Used:      fact.Used,
		}

		data = append(data, response)

	}

	return Fact{
		response: data,
	}, nil

}

func (c *Client) Get() (Fact, error) {
	fmt.Println("Making a get request")
	c.Client = &http.Client{}
	req, err := http.NewRequest("GET", getUrl("fact"), nil)

	if err != nil {
		log.Fatalf("Error creating new HTTP Request %v", err)
	}

	resp, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error sending HTTP Request %v", err)
	}

	defer resp.Body.Close()

	fmt.Printf("RESPONSE %v", resp.StatusCode)

	return formatResponse(resp)

}
