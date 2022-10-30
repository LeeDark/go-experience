package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })
	c.AddFunc("@every 1m", func() { fmt.Println("Every minute") })
	c.AddFunc("@every 30s", func() { fmt.Println("Every 30 seconds") })

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
