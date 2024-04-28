package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
	f, err := os.Open("./cmd/free_shevacadoo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	half := buffer.Streamer(buffer.Len()/4, buffer.Len()/2)
	loop := beep.Loop(3, half)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))
	<-done
}
