package main

import "fmt"

func main() {
	a := 4
	fmt.Println("1", a)
	squareValue(&a)
	fmt.Println("3", a)
}

func squareValue(v *int) {
	*v *= *v
	fmt.Println("2", v, *v)
}
