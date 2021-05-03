package main

import (
	"strconv"
	"strings"
	"time"
)

func parseUserAlarm(args []string) time.Time {
	timeElements := strings.Split(args[0], ":")
	hour, _ := strconv.Atoi(timeElements[0])
	minute, _ := strconv.Atoi(timeElements[1])
	second, _ := strconv.Atoi(timeElements[2])
	now := time.Now()
	year, month, day := now.Date()
	userTime := time.Date(year, month, day, hour, minute, second, 0, now.Location())
	return userTime
}
