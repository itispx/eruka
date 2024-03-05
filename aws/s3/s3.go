package s3

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	cfg, _    = config.LoadDefaultConfig(context.TODO())
	S3Session = s3.New(s3.Options{
		Credentials: cfg.Credentials,
		Region:      os.Getenv("AWS_REGION"),
	})
	Bucket = aws.String(os.Getenv("AWS_BUCKET"))
	S3URL  = fmt.Sprintf("https://%s.s3.%s.amazonaws.com", os.Getenv("AWS_BUCKET"), os.Getenv("AWS_REGION"))
)

func UploadS3(file *[]byte, filePath *string, bucket *string) (*s3.PutObjectOutput, string, error) {
	if bucket == nil {
		bucket = Bucket
	}

	body := bytes.NewReader(*file)

	result, err := S3Session.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: bucket,
		Key:    filePath,
		Body:   body,
	})

	return result, fmt.Sprintf("%s/%s", S3URL, *filePath), err
}

func GetS3File(key *string, bucket *string) (*s3.GetObjectOutput, error) {
	if bucket == nil {
		bucket = Bucket
	}

	return S3Session.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
}

func GetKey(filePath *string, s3URL *string) string {
	if s3URL == nil {
		s3URL = &S3URL
	}

	s := strings.Split(*filePath, *s3URL)[1]

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

func PresignedGet(key *string, bucket *string, dur time.Duration) (*v4.PresignedHTTPRequest, error) {
	if bucket == nil {
		bucket = Bucket
	}

	presignClient := s3.NewPresignClient(S3Session)

	return presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		},
		s3.WithPresignExpires(dur),
	)
}

func PresignedPut(key *string, bucket *string, dur time.Duration) (*v4.PresignedHTTPRequest, error) {
	if bucket == nil {
		bucket = Bucket
	}

	presignClient := s3.NewPresignClient(S3Session)
	return presignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
		},
		s3.WithPresignExpires(dur))
}
