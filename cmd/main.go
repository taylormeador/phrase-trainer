package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Want to flesh out functionality of creating a snippet. Basically everything a user has to do before we send the audio to Python
// 1. User uploads their own mp3
// 2. User plays their mp3 back
// 3. User seeks the mp3 and finds a section they want to create a snippet from
// 4. User runs command `createsnippet <song name> <start timestamp> <end timestamp>`
// 5. Go calls Python script that slices the audio, creates copies at different speeds, and writes them to blob? local?

func main() {
	// Loop and get commands from user
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")

		var input string

		input, _ = r.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		words := strings.Split(input, " ")

		var command string
		var args []string

		command = words[0]
		args = words[1:]

		fmt.Println(command)
		for i, a := range args {
			fmt.Println(i, a)
		}
	}
}
