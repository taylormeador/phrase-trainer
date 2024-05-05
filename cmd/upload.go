package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func upload(client *s3.Client, args []string) error {
	// TODO: Need to log the upload
	// TODO: Need to map upload to user
	bucketName := "phrase-trainer"
	objectKey := uuid.New().String()
	fileName := args[0]

	b := BucketBasics{S3Client: client}
	err := b.UploadFile(bucketName, objectKey, fileName)
	return err
}
