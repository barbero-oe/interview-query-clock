package clock

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testDate = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
var atHour = time.Date(2020, time.January, 1, 0, 30, 0, 0, time.UTC)

func Test_displays_value_on_each_tick(t *testing.T) {
	anotherTime := testDate.Add(time.Second * 2)
	var timesSaved []time.Time
	alarm := Clock{atHour, func(t time.Time) { timesSaved = append(timesSaved, t) }, func() {}}

	alarm.Tick(testDate)
	alarm.Tick(anotherTime)

	assert.Equal(t, testDate, timesSaved[0])
	assert.Equal(t, anotherTime, timesSaved[1])
}

func Test_sounds_the_alarm_on_the_specified_moment(t *testing.T) {
	wasCalled := false
	alarm := Clock{atHour, func(_ time.Time) {}, func() { wasCalled = true }}

	alarm.Tick(testDate)
	alarm.Tick(atHour)

	assert.True(t, wasCalled)
}
