package main

import (
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(beep.Take(sr.N(5*time.Second), Noise{}), beep.Callback(func() {
		done <- true
	})))
	<-done
}
