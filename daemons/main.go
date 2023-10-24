package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGTERM)

	go func() {
		for {
			select {
			case sig := <-signals:
				if sig == syscall.SIGTERM {
					fmt.Println("Received termination signal. Exiting ....")
					os.Exit(0)
				}

			default:
				fmt.Println("Hello, World!")
				time.Sleep(time.Second * 30)
			}
		}
	}()

	<-signals
}
