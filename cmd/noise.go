package main

import "math/rand"

type Noise struct{}

func (no Noise) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		samples[i][0] = rand.Float64()*2 - 1
		samples[i][1] = rand.Float64()*2 - 1
	}
	return len(samples), true
}

func (no Noise) Err() error {
	return nil
}
