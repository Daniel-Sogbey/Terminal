package main

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	API_KEY             = "iTMgJHGQ4Qnj7JGhRIRkelTnu"
	API_KEY_SECRET      = "83SRq1X9anvAOKouODgYaWZfzHiDrtNrTb8YhPyng1Zo8x9V09"
	ACCESS_TOKEN        = "1043075477609611265-x8qcYL9BPRgbvt3ruxSMdH0Yc51Qx4"
	ACCESS_TOKEN_SECRET = "RsgqLdUlf7yhE3WcbecTUxi5AU6eGCBMD1aq80Od2F00c"
	CLIENT_ID           = "V0Itbzl3RF91bFB3QkU3OWREMmM6MTpjaQ"
	CLIENT_SECRET       = "-LKJ19w0Pfmqy1U9KHx7Wwwv6yzuFvxo5hO22eIX3cW8io3W4x"
)

type Credentials struct {
	ApiKey            string
	ApiKeySecret      string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(c *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(c.ApiKey, c.ApiKeySecret)

	token := oauth1.NewToken(c.AccessToken, c.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	verifyParams := twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(&verifyParams)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Hello User \n%+v\n", user)

	return client, nil

}

func main() {
	fmt.Println("Hello, World!")
	credentials := Credentials{
		ApiKey:            API_KEY,
		ApiKeySecret:      API_KEY_SECRET,
		AccessToken:       ACCESS_TOKEN,
		AccessTokenSecret: ACCESS_TOKEN_SECRET,
	}
	client, err := getClient(&credentials)
	if err != nil {
		log.Println("Error getting twitter client", err)
		return
	}

	log.Printf("\n%v\n", client)
}
