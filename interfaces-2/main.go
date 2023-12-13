package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Square struct {
	l float64
}

type Rectangle struct {
	l float64
	b float64
}

func main() {
	square := Square{
		l: 5,
	}

	rectangle := Rectangle{
		l: 5,
		b: 7,
	}
	printArea(rectangle)
	printArea(square)

	// fmt.Println(areaOfRectange(rectangle))
	// fmt.Println(areaOfSquare(square))
}

// func areaOfRectange(r Rectangle) float64 {
// 	return r.l * r.b
// }

// func areaOfSquare(s Square) float64 {
// 	return float64(math.Pow(float64(s.l), 2))
// }

func printArea(s Shape) {
	fmt.Println(s.area())
}

func (s Square) area() float64 {
	return math.Pow(float64(s.l), 2)
}

func (r Rectangle) area() float64 {
	return r.l * r.b
}
