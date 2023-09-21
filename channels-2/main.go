package main

import (
	"channels-2/helpers"
	"fmt"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumbers(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)

	fmt.Println(helpers.MinimumNumber(1, 3))
	fmt.Println(helpers.MaximumNumber(10, 7))
}
