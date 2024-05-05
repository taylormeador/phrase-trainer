package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Want to flesh out functionality of creating a phrase. Basically everything a user has to do before we send the audio to Python
// 1. User uploads their own mp3 ***DONE***
// 2. User plays their mp3 back
// 3. User seeks the mp3 and finds a section they want to create a phrase from
// 4. User runs command `createphrase <song name> <start timestamp> <end timestamp>`
// 5. Go calls Python script that slices the audio, creates copies at different speeds, and writes them to blob? local?

func main() {
	// Loop and get commands from user
	r := bufio.NewReader(os.Stdin)

	for {
		// Parse user input into variables
		fmt.Print("-> ")

		var input string

		input, _ = r.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		words := strings.Split(input, " ")

		var command string
		var args []string

		command = words[0]
		args = words[1:]

		// Load the Shared AWS Configuration (~/.aws/config)
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		// Create an Amazon S3 service client
		client := s3.NewFromConfig(cfg)

		switch command {
		case "upload":
			bucketName := "phrase-trainer"
			objectKey := uuid.New().String()
			fileName := args[0]
			b := BucketBasics{S3Client: client}
			b.UploadFile(bucketName, objectKey, fileName)
		}
	}
}
