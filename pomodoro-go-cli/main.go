package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to Pomodoro CLI App")
	// currentTask = AddTask("Study algorithms and data structures")

	taskPtr := flag.String("task", "", "Task to add. (Required)")

	flag.Parse()

	if *taskPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	currentTask = AddTask(*taskPtr)

	fmt.Printf("Start the task! Focus on %s\n", currentTask)

	timer1 := time.NewTimer(25 * time.Second)

	<-timer1.C

	fmt.Println("Congrats!, Task time is complete. Take break")
}
