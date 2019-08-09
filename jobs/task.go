package jobs

import (
	"github.com/robfig/cron"
	"log"
)

var task = *cron.New()

// 启动任务调度
func TaskRun() {
	task1()
	task2()
	task.Start()
}

// 任务1
func task1() {
	//spec := "0/60 * * * * *"
	//task.AddFunc(spec, func() {
	//	log.Println("cron1 running")
	//})
}

// 任务2
func task2() {
	spec := "0/10 * * * * *"
	task.AddFunc(spec, func() {
		log.Println("task2 running")
	})
}