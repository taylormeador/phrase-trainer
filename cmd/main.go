package main

import (
	"log"
	"os"

	"github.com/200sc/klangsynthese/audio/filter"
	kmp3 "github.com/200sc/klangsynthese/mp3"
)

func main() {
	f, err := os.Open("./cmd/free_shevacadoo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	snippet, err := kmp3.Load(f)
	if err != nil {
		log.Fatal(err)
	}

	ps, err := filter.NewFFTShifter(2048, 32)
	if err != nil {
		log.Fatal(err)
	}
	encoding := ps.PitchShift(0.5)

	// encoding := filter.Speed(2, ps)

	snippet, err = snippet.Filter(encoding)
	if err != nil {
		log.Fatal(err)
	}

	snippet.Play()
	select {}

}
