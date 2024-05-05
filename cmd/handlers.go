package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"

	"phrasetrainer.tm.com/internal"
)

func Upload(client *s3.Client, args []string) error {
	// TODO: Need to log the upload
	// TODO: Need to map upload to user
	objectKey := uuid.New().String()
	fileName := args[0]

	b := internal.BucketBasics{S3Client: client}
	err := b.UploadFile(bucketName, objectKey, fileName)
	return err
}

func List(client *s3.Client) error {
	basics := internal.BucketBasics{S3Client: client}
	objects, err := basics.ListObjects(bucketName)
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
	basics := internal.BucketBasics{S3Client: client}
	objects, err := basics.ListObjects(bucketName)
	if err != nil {
		log.Fatal(err)
	}

	var keys []string
	for _, o := range objects {
		keys = append(keys, *o.Key)
	}
	err = basics.DeleteObjects(bucketName, keys)

	return err
}
