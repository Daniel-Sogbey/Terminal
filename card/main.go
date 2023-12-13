package main

import (
	"fmt"
	"sort"
)

func main() {

	// cards := newDeckFromFile("my_cards")

	// cards.shuffle()
	// cards.print()
	// array := []int{4, 6}
	// targetSum := 10

	// fmt.Println(twoNumberSum(array, targetSum))
	im := make(map[int]bool)

	

	fmt.Println(im)
}

func twoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[i]+array[j] == target && array[i] != array[j] {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{1, 2}
}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.

   sort.Ints(array)
    
    left := 0
    right := len(array) - 1

    for left < right {
        currentSum := array[left] + array[right]

        if currentSum == target {
            return []int{array[left], array[right]}
        }

        left = left + 1
        
    }

    
    
	return []int{}
}
