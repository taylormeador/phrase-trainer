package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"phrasetrainer.tm.com/internal/aws/s3utils"
	"phrasetrainer.tm.com/internal/constants"
	"phrasetrainer.tm.com/internal/db"
)

func LogUpload(conn *pgx.Conn, userID int, fileName string, fileLabel string, objectKey string) error {
	err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.InsertRow(context.Background(), tx, userID, time.Now(), fileName, fileLabel, objectKey)
	})
	return err
}

// Upload wraps the UploadFile method from AWS and
// uses a UUID for the name of the blob.
func Upload(client *s3.Client, conn *pgx.Conn, args []string) error {
	objectKey := uuid.New().String()
	fileName := args[0]

	b := s3utils.BucketBasics{S3Client: client}
	err := b.UploadFile(constants.BucketName, objectKey, fileName)
	if err != nil {
		return err
	}

	userID := 1 // TODO: user stuff not implemented yet
	fileLabel := constants.USER_MP3_UPLOAD
	err = LogUpload(conn, userID, fileName, fileLabel, objectKey)
	return err
}

func Play(client *s3.Client, conn *pgx.Conn, args []string) error {
	// play <song name>
	// look up blob from user_uploads
	// download file
	// play with beep library
	userID := 1
	songName := args[0]
	err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		blobName, err := db.GetUserUploadData(context.Background(), tx, userID, songName)
		fmt.Println(blobName)
		return err
	})

	return err
}

// List wraps the ListObjects method from AWS and prints out the
// name of each blob on a new line.
func List(client *s3.Client) error {
	basics := s3utils.BucketBasics{S3Client: client}
	objects, err := basics.ListObjects(constants.BucketName)
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range objects {
		fmt.Println(*o.Key)
	}

	return err
}

// DeleteAll deletes all items in the bucket. This is a convenience function
// intended to be run manually. This will save costs by getting rid of objects
// created during development.
func DeleteAll(client *s3.Client) error {
	basics := s3utils.BucketBasics{S3Client: client}
	objects, err := basics.ListObjects(constants.BucketName)
	if err != nil {
		log.Fatal(err)
	}

	var keys []string
	for _, o := range objects {
		keys = append(keys, *o.Key)
	}
	err = basics.DeleteObjects(constants.BucketName, keys)

	return err
}
