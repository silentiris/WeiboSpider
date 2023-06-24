package main

import (
	"WeiboSpider/func"
	"github.com/robfig/cron/v3"
)

func main() {
	var flag string
	crontab := cron.New(cron.WithSeconds())
	task := _func.TaskFunc(flag)
	spec := "0 */1 * * * ?"
	_, _ = crontab.AddFunc(spec, task)
	crontab.Start()
	select {}
}
