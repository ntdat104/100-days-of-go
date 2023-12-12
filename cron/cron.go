package cron

import (
	"fmt"
	"time"

	"gopkg.in/robfig/cron.v2"
)

func Main() {
	c := cron.New()

	now := time.Now()

	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("[Info]:", now.String())
	})
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("TZ=Asia/Bangkok 30 04 * * * *", func() { fmt.Println("Runs at 04:30 Bangkok time every day") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.Start()
	c.AddFunc("@every 0h0m1s", func() { fmt.Println(now.Hour(), now.Minute(), now.Second()) })

	// Funcs are invoked in their own goroutine, asynchronously.

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	time.Sleep(10 * time.Second)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
