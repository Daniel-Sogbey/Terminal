package main

import "fmt"

type numbers []int

func checkForOddAndEven() {
	numbers := numbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for _, number := range numbers {
	// if number%2 == 1 {
	// fmt.Println("odd")
	// } else {
	// fmt.Println("even")
	// }
	// }

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
