package main

import (
	"log"
	"os"

	"github.com/200sc/klangsynthese/audio/filter"
	kmp3 "github.com/200sc/klangsynthese/mp3"
)

// func createSnippet(s beep.Streamer, f beep.Format, start int, end int, loops int, ratio float64) beep.Streamer {
// 	buffer := beep.NewBuffer(f)
// 	buffer.Append(s)

// 	half := buffer.Streamer(buffer.Len()/start, buffer.Len()/end)
// 	loop := beep.Loop(loops, half)

// 	fast := beep.ResampleRatio(4, ratio, loop)

// 	return fast
// }

func main() {
	f, err := os.Open("./cmd/free_shevacadoo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	// _, format, err := mp3.Decode(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	snippet, err := kmp3.Load(f)
	if err != nil {
		log.Fatal(err)
	}
	snippet.Play()

	ps, err := filter.NewFFTShifter(2, 1)
	if err != nil {
		log.Fatal(err)
	}
	encoding := ps.PitchShift(1.5)

	snippet, err = snippet.Filter(encoding)
	if err != nil {
		log.Fatal(err)
	}

	// snippet.Play()

	select {}

	// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// done := make(chan bool)
	//
	//	speaker.Play(beep.Seq(snippet, beep.Callback(func() {
	//		done <- true
	//	})))
	//
	// <-done
}
