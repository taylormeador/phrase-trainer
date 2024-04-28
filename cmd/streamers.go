package main

import (
	"math/rand"

	"github.com/gopxl/beep"
)

type SnippetStreamer struct {
	beep.Streamer
}

func (no SnippetStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0] = rand.Float64()*2 - 1
		samples[i][1] = rand.Float64()*2 - 1
	}
	return len(samples), true
}

func (no SnippetStreamer) Err() error {
	return nil
}

type SnippetStreamerSeeker struct {
	SnippetStreamer
}

func (sss SnippetStreamerSeeker) Len() (n int) {
	return 200
}

func (sss SnippetStreamerSeeker) Position() (pos int) {
	return 100
}

func (sss SnippetStreamerSeeker) Seek(p int) error {
	return nil
}
