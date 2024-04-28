package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	// A zero Queue is an empty Queue.
	var queue Queue
	speaker.Play(&queue)

	for {
		var name string
		fmt.Print("Type an MP3 file name: ")
		fmt.Scanln(&name)

		// Open the file on the disk.
		f, err := os.Open(name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Decode it.
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// The speaker's sample rate is fixed at 44100. Therefore, we need to
		// resample the file in case it's in a different sample rate.
		resampled := beep.Resample(4, format.SampleRate, sr, streamer)

		// And finally, we add the song to the queue.
		speaker.Lock()
		queue.Add(resampled)
		speaker.Unlock()
	}
}
