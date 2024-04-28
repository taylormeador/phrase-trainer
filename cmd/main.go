package main

import (
	"log"
	"os"

	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/mp3"
)

func main() {
	f, err := os.Open("./cmd/free_shevacadoo.mp3")
	if err != nil {
		log.Fatal(err)
	}

	snippet, err := mp3.Load(f)
	if err != nil {
		log.Fatal(err)
	}

	ps, err := filter.NewFFTShifter(2048, 32)
	if err != nil {
		log.Fatal(err)
	}

	encoding := ps.PitchShift(1.5)

	snippet, err = snippet.Filter(encoding)
	if err != nil {
		log.Fatal(err)
	}

	loop := filter.LoopOn()

	snippet, err = loop.Apply(snippet)
	if err != nil {
		log.Fatal(err)
	}

	snippet.Play()
	select {}

}
