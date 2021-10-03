package utils

import "github.com/robfig/cron/v3"

var Cron *cron.Cron

func LoadCron() {
	Cron = cron.New(cron.WithSeconds())
}
func StartCron() {
	Cron.Start()
}
