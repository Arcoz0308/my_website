package commands

import (
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"strconv"
	"time"
)

type Uptime struct{}

func (*Uptime) Run(args []string) {
	logger.AppInfof("command:uptime", "uptime : %s", durationToString(time.Since(utils.Start)))
}

func durationToString(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) - (days * 24)
	minutes := int(d.Minutes()) - (days * 24) - (hours * 60)
	seconds := int(d.Seconds()) - (days * 24) - (hours * 60) - (minutes * 60)
	s := ""
	if seconds > 0 {
		s = strconv.Itoa(seconds) + "s" + s
	}
	if minutes > 0 {
		s = strconv.Itoa(minutes) + "m" + s
	}
	if hours > 0 {
		s = strconv.Itoa(hours) + "h" + s
	}
	if days > 0 {
		s = strconv.Itoa(days) + "d" + s
	}
	return s
}
