package app

import (
	"fmt"
	"os"
	"github.com/robfig/cron/v3"
)

func MakeAndRunScheduler(s *Server){
	cron_str := os.Getenv("CRON_STR_FOR_SCHED")
	
	fmt.Println("\ncron string:", cron_str)

	c := cron.New()
	c.AddFunc(cron_str, func() {
		fmt.Println("SCHEDULED TASK INITIATED AND RUNNING")
		FakeFinder()
	})
	c.Start()
	fmt.Println("Scheduler Started Successfully")
}