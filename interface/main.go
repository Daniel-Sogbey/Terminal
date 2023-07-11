package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	r := rectangle{
		width:  20,
		height: 10,
	}

	c := circle{
		radius: 15,
	}

	printGreeting(eb)
	printGreeting(sb)
	printArea(r)
	printArea(c)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
