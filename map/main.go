package main

import "fmt"

func main() {

	var colors map[string]string = map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors := make(map[string]string)

	printMap(colors)

	fmt.Println(colors)
	var c byte = 'a'

	var z rune = 'a'

	fmt.Println(c)
	fmt.Println(z)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
