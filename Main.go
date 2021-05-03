package main

import (
	"github.com/barbero-oe/interview-query-clock/alarm"
	"github.com/barbero-oe/interview-query-clock/display"
	"os"
	"time"

	"github.com/barbero-oe/interview-query-clock/clock"
)

func main() {
	alarmAt := parseUserAlarm(os.Args[1:])

	streamer, done, startAlarm := alarm.CreateAlarm("alarm.mp3")
	defer streamer.Close()

	clockAlarm := clock.Clock{
		AlarmAt:    alarmAt,
		Display:    display.SimpleConsolePrinter,
		StartAlarm: startAlarm}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			clockAlarm.Tick(t)
		}
	}
}

