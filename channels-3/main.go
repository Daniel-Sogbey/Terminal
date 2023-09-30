package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	result := <-ch

	fmt.Printf("Result from channel %d", result)
}
