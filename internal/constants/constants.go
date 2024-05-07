package constants

import "os"

var BucketName = os.Getenv("BUCKET_NAME")

// Set of valid values for `fileLabel` field in `user_uploads` table.
const (
	USER_MP3_UPLOAD string = "User MP3 Upload"
)
