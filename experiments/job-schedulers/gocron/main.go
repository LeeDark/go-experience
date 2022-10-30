package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().SingletonMode().Do(func(){ fmt.Println("Every 5 seconds") })

	// strings parse to duration
	s.Every("5m").SingletonMode().Do(func(){ fmt.Println("Every 5 minutes") })

	s.Every(5).Days().SingletonMode().Do(func(){ fmt.Println("Every 5 days") })

	// cron expressions supported
	var task = func() {
		fmt.Println("Every 1 minute for test")
	}
	s.Cron("*/1 * * * *").SingletonMode().Do(task) // every minute

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	s.StartAsync()
	// starts the scheduler and blocks current execution path
	//s.StartBlocking()
}