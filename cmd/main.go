package main

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func createSnippet(s beep.Streamer, f beep.Format, start int, end int, loops int, ratio float64) beep.Streamer {
	buffer := beep.NewBuffer(f)
	buffer.Append(s)

	half := buffer.Streamer(buffer.Len()/start, buffer.Len()/end)
	loop := beep.Loop(loops, half)

	fast := beep.ResampleRatio(4, ratio, loop)

	// TODO correct the pitch here.

	return fast
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

	start := 4
	end := 2
	loops := 2
	ratio := 1.1
	snippet := createSnippet(streamer, format, start, end, loops, ratio)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(snippet, beep.Callback(func() {
		done <- true
	})))
	<-done
}
