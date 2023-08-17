package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type logWriter struct {
// }

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// io.Reader.Read(resp.Body, bs)

	// fmt.Println(string(bs))
	// 0xc000132000
	// 0xc0000c8260

	file, err := os.Create("google.txt")

	if err != nil {
		os.Exit(1)
	}

	io.Copy(file, resp.Body)

}
