package main

import (
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(Noise{})
	select {}
}
