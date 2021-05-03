package clock

import (
	"time"
)

type Clock struct {
	AlarmAt    time.Time
	Display    func(t time.Time)
	StartAlarm func()
}

func (clock Clock) Tick(date time.Time) {
	clock.Display(date)
	currentTime := timeElements(date)
	alarmTime := timeElements(clock.AlarmAt)
	if currentTime == alarmTime {
		clock.StartAlarm()
	}
}

func timeElements(date time.Time) [3]int {
	return [3]int{date.Hour(), date.Minute(), date.Second()}
}
