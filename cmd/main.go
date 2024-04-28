package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func createSnippet(s beep.Streamer, f beep.Format, start int, end int, loops int) beep.Streamer {
	buffer := beep.NewBuffer(f)
	buffer.Append(s)

	half := buffer.Streamer(buffer.Len()/start, buffer.Len()/end)
	loop := beep.Loop(loops, half)

	return loop
}

func main() {
	f, err := os.Open("./cmd/free_shevacadoo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	snippet := createSnippet(streamer, format, 4, 2, 3)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(snippet, beep.Callback(func() {
		done <- true
	})))
	<-done
}
