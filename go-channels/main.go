package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

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
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// func fact(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	fmt.Println(n)
// 	return n * fact(n-1)
// }

// func main() {
// 	fmt.Println("Enter your amount : ")
// 	reader := bufio.NewReader(os.Stdin)

// 	input, err := reader.ReadString('\n')

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	input = strings.TrimSpace(input)

// 	amount, err := strconv.ParseFloat(input, 64)

// 	tip := fmt.Sprintf("%v", calculateTip(amount))
// 	io.Copy(os.Stdout, strings.NewReader(tip))
// }

// func calculateTip(amount float64) float64 {
// 	return 0.2 * amount
// }

// fmt.Println("Hello,world!")
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		c1 <- "one"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		c2 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("Received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("Received", msg2)
// 		}
// 	}

// 	// factorial := fact(5)
// 	// fmt.Println("Factorial", factorial)
