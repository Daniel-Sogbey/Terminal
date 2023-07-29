package main

import (
	"log"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

func main() {
	log.Println("Hello, world!")
	r := Rectangle{width: 10, length: 20}
	s := Square{length: 10}
	c := Circle{radius: 5}

	printArea(r)
	printArea(s)
	printArea(c)
}

func printArea(x Shape) {
	log.Println(x.area())
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (s Square) area() float64 {
	return math.Pow(s.length, 2)
}
