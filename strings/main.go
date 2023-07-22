package main

import "fmt"

func main() {
	a := "123434"
	b := "098700"

	fmt.Println(findCommonItem(a, b))
}

// solution 1 : Time O(n^2) | Space O(1)
func findCommonItem(a string, b string) string {
	str := ""
	for _, i := range a {
		for _, j := range b {
			if i == j {
				str += string(i)
			}
		}
	}
	return str
}
