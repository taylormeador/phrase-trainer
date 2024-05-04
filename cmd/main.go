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

// Want to flesh out functionality of creating a snippet. Basically everything a user has to do before we send the audio to Python
// 1. User uploads their own mp3
// 2. User plays their mp3 back
// 3. User seeks the mp3 and finds a section they want to create a snippet from
// 4. User runs command `createsnippet <song name> <start timestamp> <end timestamp>`
// 5. Go calls Python script that slices the audio, creates copies at different speeds, and writes them to blob? local?

func main() {
	// Open file
	f, err := os.Open("./cmd/free_shevacadoo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Decode audio to Streamer
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Create the snippet with parameters
	start := 4
	end := 2
	loops := 2
	ratio := 1.1
	snippet := createSnippet(streamer, format, start, end, loops, ratio)

	// Play
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(snippet, beep.Callback(func() {
		done <- true
	})))
	<-done
}
