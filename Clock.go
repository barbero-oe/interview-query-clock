package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
)

func main() {
	streamer := loadAlarm("alarm.mp3")
	defer streamer.Close()

	done := make(chan bool)
	startAlarm := func() {
		speaker.Play(beep.Seq(streamer, beep.Callback(func() { done <- true })))
	}

	printTime := func(t time.Time) { fmt.Println(t) }

	clock := Clock{time.Now().Add(time.Second * 5), printTime, startAlarm}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			clock.Tick(t)
		}
	}
}

func loadAlarm(file string) beep.StreamSeekCloser {
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

type Clock struct {
	alarm      time.Time
	display    func(t time.Time)
	startAlarm func()
}

func (clock Clock) Tick(date time.Time) {
	clock.display(date)
	currentTime := timeElements(date)
	alarmTime := timeElements(clock.alarm)
	if currentTime == alarmTime {
		clock.startAlarm()
	}
}

func timeElements(date time.Time) [3]int {
	return [3]int{date.Hour(), date.Minute(), date.Second()}
}
