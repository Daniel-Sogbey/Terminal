package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello,world!")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}

	factorial := fact(5)
	fmt.Println("Factorial", factorial)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fact(n-1)
}
