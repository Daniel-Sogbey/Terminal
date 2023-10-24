package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./IMG-8203.jpg")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	fmt.Println(bufio.NewReader(file))
}
