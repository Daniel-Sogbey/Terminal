package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
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
	return math.Pi
}
