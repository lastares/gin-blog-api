package crontab

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func CronLaunch() {
	log.Println("Starting......")

	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run my cron start ...")
		fmt.Println("定时任务进行中")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
