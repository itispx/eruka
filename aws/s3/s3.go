package s3

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	cfg, _    = config.LoadDefaultConfig(context.TODO())
	s3session = s3.New(s3.Options{
		Credentials: cfg.Credentials,
		Region:      os.Getenv("AWS_REGION"),
	})
	Bucket = aws.String(os.Getenv("AWS_BUCKET"))
	S3URL  = fmt.Sprintf("https://%s.s3.%s.amazonaws.com", os.Getenv("AWS_BUCKET"), os.Getenv("AWS_REGION"))
)

func UploadS3(file *[]byte, filePath *string) (*s3.PutObjectOutput, string, error) {
	body := bytes.NewReader(*file)

	result, err := s3session.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: Bucket,
		Key:    filePath,
		Body:   body,
	})

	return result, fmt.Sprintf("%s/%s", S3URL, *filePath), err
}

func GetS3File(filePath *string) (*s3.GetObjectOutput, error) {
	return s3session.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: Bucket,
		Key:    filePath,
	})
}

func GetKey(filePath *string) string {
	s := strings.Split(*filePath, S3URL)[1]

	for i := range s {
		if i > 0 {
			// The value i is the index in s of the second
			// rune.  Slice to remove the first rune.
			return s[i:]
		}
	}
	// There are 0 or 1 runes in the string.
	return ""
}
