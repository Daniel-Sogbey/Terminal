package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	printArea()
	perimeter() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float32
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func (r rectangle) perimeter() float64 {
	return 2 * ((r.width) + (r.height))
}

// Assignment
func (s square) getArea() float64 {
	return s.sideLength * 2
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}

func (s square) printArea() {
	fmt.Println(s.getArea())
}
