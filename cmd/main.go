package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	"phrasetrainer.tm.com/internal/aws/s3utils"
	"phrasetrainer.tm.com/internal/db"
)

// Want to flesh out functionality of creating a phrase. Basically everything a user has to do before we send the audio to Python
// 1. User uploads their own mp3 ***DONE***
// 2. User plays their mp3 back
// 3. User seeks the mp3 and finds a section they want to create a phrase from
// 4. User runs command `createphrase <song name> <start timestamp> <end timestamp>`
// 5. Go calls Python script that slices the audio, creates copies at different speeds, and writes them to blob? local?

func main() {

	// Connect to database and create table
	conn := db.Connect()
	defer conn.Close(context.Background())

	err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.InitTable(context.Background(), tx)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Loop and get commands from user
	r := bufio.NewReader(os.Stdin)

	for {
		// Parse user input into variables
		fmt.Print("-> ")
		command, args := parseInput(r)

		client := s3utils.GetS3Client()

		switch command {
		case "upload":
			err := Upload(client, conn, args)
			if err != nil {
				log.Fatal(err)
			}
		case "play":
			err := Play(client, conn, args)
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
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Command not recognized")
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
