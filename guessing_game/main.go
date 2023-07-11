package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	guessNumber()
}

func guessNumber() {
	min, max := 1, 100

	var randomNumber int = rand.Intn(max-min) + min
	var userGuess int = getUserInput()

	fmt.Println("The secret number is ", randomNumber, " but you guessed ", userGuess)
}

func getUserInput() int {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occurred whiles reading input", err)
		return 0
	}

	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value", err)
		return 0
	}

	fmt.Println("Your guess is ", guess)

	return guess
}
