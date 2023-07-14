package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	file, err := os.Open(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)

	// fmt.Println(*file)

	// data, err := os.ReadFile(args[1])

	// if err != nil {
	// fmt.Println(err)
	// os.Exit(1)
	// }

	// os.Stdout.Write(data)

}
