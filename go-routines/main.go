package main

import (
	"fmt"
	"time"
)

// func f(from string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from, ":", i)
// 	}
// }

func main() {

	// //how a normal function would run
	// f("direct")

	// //go routine 1 and 2 run concurrently (haphazardly)

	// //go routine 1
	// go f("goroutine")

	// //go routine 2
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Millisecond * 1)

	// fmt.Println("done")

	message := make(chan string)
	var msg string

	before := time.Now()

	go func() {
		message <- "ping"
	}()

	go func() {
		msg = <-message
	}()

	after := time.Now()

	time.Sleep(time.Second * 5)

	fmt.Println(msg, "Duration:", after.Sub(before))
}
