package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, world!")
	result, err := divide(2, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of division is ", result)

}

func divide(x float32, y float32) (float32, error) {
	var result float32
	if y == 0.0 {
		return result, errors.New("zero division error")
	}

	result = x / y

	return result, nil
}
