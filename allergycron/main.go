package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	fmt.Println("Hello, World!")
	loc, err := time.LoadLocation("Europe/Vienna")

	if err != nil {
		panic(err)
	}

	cronJob := cron.NewWithLocation(loc)

	//seconds, minutes, hours, day of month, month, day of week
	cronJob.AddFunc("0 0 8 * * *", func() {

	})

	cronJob.Start()

	// select {}

}
