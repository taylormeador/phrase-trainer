package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"phrasetrainer.tm.com/internal"
)

// Want to flesh out functionality of creating a phrase. Basically everything a user has to do before we send the audio to Python
// 1. User uploads their own mp3 ***DONE***
// 2. User plays their mp3 back
// 3. User seeks the mp3 and finds a section they want to create a phrase from
// 4. User runs command `createphrase <song name> <start timestamp> <end timestamp>`
// 5. Go calls Python script that slices the audio, creates copies at different speeds, and writes them to blob? local?

var bucketName = os.Getenv("BUCKET_NAME")

func main() {
	// Loop and get commands from user
	r := bufio.NewReader(os.Stdin)

	for {
		// Parse user input into variables
		fmt.Print("-> ")
		command, args := parseInput(r)

		client := internal.GetS3Client()

		switch command {
		case "upload":
			err := Upload(client, args)
			if err != nil {
				log.Fatal(err)
			}
		case "list":
			err := List(client)
			if err != nil {
				log.Fatal(err)
			}
		case "delete-all-s3-objects":
			err := DeleteAll(client)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func parseInput(r *bufio.Reader) (string, []string) {
	var input string
	var command string
	var args []string

	input, _ = r.ReadString('\n')

	// Make CR LF into LF
	input = strings.Replace(input, "\n", "", -1)

	words := strings.Split(input, " ")

	// Commands will always be one word (the first word)
	// followed by any number of arguments (if any)
	command = words[0]
	args = words[1:]

	// TODO: validate `command` and `args`

	return command, args
}
