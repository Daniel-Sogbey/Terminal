package main

import "fmt"

func main() {

	strArray := []string{"hi", "hello", "Bye"}
	intArray := []int{1, 2, 3}

	Print(strArray)
	Print(intArray)

}

//without generics

// func Print(s []string) {
// 	for i, v := range s {
// 		fmt.Println(i, v)
// 	}
// }

// func Print(s []int) {
// 	for i, v := range s {
// 		fmt.Println(i, v)
// 	}
// }

//with generics

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
