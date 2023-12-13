package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	b := []string{"c", "d", "e", "f"}

	fmt.Println(common(a, b))
	fmt.Println(commonV1(a, b))
}

// Time : O(n^2) | Space : O(1)
func common(a []string, b []string) []string {
	commonItems := []string{}

	for _, itemInA := range a {
		for _, itemInB := range b {
			if itemInA == itemInB {
				commonItems = append(commonItems, []string{itemInA}...)
			}
		}
	}
	return commonItems
}

// Time : O(n*m) | Space : O(1)
func commonV1(a []string, b []string) []string {
	hashMap := make(map[string]bool)
	var commonItems []string

	for _, itemInA := range a {
		hashMap[itemInA] = true
	}

	for _, itemInB := range b {
		if hashMap[itemInB] {
			commonItems = append(commonItems, []string{itemInB}...)
		}
	}

	return commonItems
}
