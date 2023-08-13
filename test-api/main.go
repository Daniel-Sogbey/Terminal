package main

import (
	"fmt"
	"log"

	"github.com/Daniel-Sogbey/test-api/requests"
)

func main() {
	fmt.Println("Testing API")
	c := &requests.Client{}

	resp, err := c.Get()

	if err != nil {
		log.Fatalf("Error making GET Request %v", err)
	}

	fmt.Println(resp)
}
