package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	flag.String("host", "", "host")
	flag.Parse()

	var args = os.Args

	if len(args) != 3 {
		log.Fatalf("expected 2 command line args , but got %d", len(args))
	}

	host := args[2]

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatalf("failed to get ip of host . err: %v", err)
	}

	if len(ips) == 0 {
		fmt.Println("no ips found for the host", host)
	}

	for _, ipV4 := range ips {
		if ipV4.To4() != nil {
			log.Println("IPV4", ipV4)
		}
	}

	for _, ipV6 := range ips {
		if ipV6.To4() == nil {
			log.Println("IPV6", ipV6)
		}
	}

	// fmt.Println("IPS", ips)
}
