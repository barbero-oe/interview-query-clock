package alarm

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func CreateAlarm(file string) (beep.StreamSeekCloser, chan bool, func()) {
	streamer := loadAlarmSound("alarm.mp3")

	done := make(chan bool)
	startAlarm := func() {
		speaker.Play(beep.Seq(streamer, beep.Callback(func() { done <- true })))
	}
	return streamer, done, startAlarm
}

func loadAlarmSound(file string) beep.StreamSeekCloser {
	sound, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(sound)
	if err != nil {
		log.Fatal(err)
	}
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
	return streamer
}
