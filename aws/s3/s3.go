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

func UploadS3(file *[]byte, filePath *string, bucket *string, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, string, error) {
	if bucket == nil {
		bucket = Bucket
	}

	body := bytes.NewReader(*file)

	result, err := S3Session.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: bucket,
		Key:    filePath,
		Body:   body,
	}, optFns...)

	return result, fmt.Sprintf("%s/%s", S3URL, *filePath), err
}

func GetS3File(key *string, bucket *string, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	if bucket == nil {
		bucket = Bucket
	}

	return S3Session.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	}, optFns...)
}

func GetKey(filePath *string, s3URL *string) string {
	if s3URL == nil {
		s3URL = &S3URL
	}

	s := strings.Split(*filePath, *s3URL)[1]

	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}

	return ""
}

func PresignedGet(key *string, bucket *string, dur time.Duration, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if bucket == nil {
		bucket = Bucket
	}

	presignClient := s3.NewPresignClient(S3Session)

	options := make([]func(*s3.PresignOptions), 0, len(optFns)+1)
	options = append(options, s3.WithPresignExpires(dur))
	options = append(options, optFns...)

	return presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		},
		options...,
	)
}

func PresignedPut(key *string, bucket *string, dur time.Duration, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if bucket == nil {
		bucket = Bucket
	}

	presignClient := s3.NewPresignClient(S3Session)

	options := make([]func(*s3.PresignOptions), 0, len(optFns)+1)
	options = append(options, s3.WithPresignExpires(dur))
	options = append(options, optFns...)

	return presignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: bucket,
			Key:    key,
		},
		options...,
	)
}

// 1. Setting Up and Configuration
// Initialize the SDK: Import the AWS package and configure credentials (e.g., via IAM roles, environment variables, or specifying them directly in the code) and region to start interacting with S3.
// Session Creation: Create a session using the SDK, which is a necessary step before executing any operations against S3 services.

// 2. Bucket Operations
// Create/Delete Buckets: Programmatically create or delete S3 buckets.
// List Buckets: Retrieve a list of all S3 buckets in your account.
// Bucket Policies: Manage policies to control access to S3 buckets, including setting bucket policies for permissions and public access.
// Bucket CORS: Configure Cross-Origin Resource Sharing (CORS) settings for your buckets.

// 3. Object Operations
// Upload Files: Upload files to S3 buckets, with support for large files via multipart uploads.
// Download Files: Download files from S3 buckets to your local system or application.
// Delete Files: Remove files from buckets.
// List Objects: List the objects stored in a specific S3 bucket, with support for pagination.
// Copy Objects: Copy objects from one bucket to another or within the same bucket.

// 4. Advanced Features
// Pre-signed URLs: Generate pre-signed URLs for secure and temporary access to S3 objects without requiring AWS credentials.
// Versioning: Manage versioning of objects within S3 buckets, allowing you to preserve, retrieve, and restore every version of every object stored in your buckets.
// Lifecycle Policies: Automate moving objects between different storage tiers or deleting old objects with lifecycle policies.
// Encryption: Manage server-side encryption (SSE) settings for objects to secure data at rest.

// 5. Security and Access Control
// Access Control Lists (ACLs): Set up ACLs for buckets and objects to manage read and write permissions.
// IAM Integration: Utilize AWS Identity and Access Management (IAM) for fine-grained access control to your S3 resources.

// 6. Performance Optimization
// Transfer Acceleration: Enable and configure Amazon S3 Transfer Acceleration for faster upload and download speeds by utilizing Amazon CloudFront's globally distributed edge locations.
// Multipart Uploads: Optimize the upload of large files by breaking them into smaller parts and uploading them in parallel, improving throughput and reducing the impact of network issues.

// 7. Monitoring and Logging
// Logging: Enable access logging on S3 buckets for auditing and monitoring requests.
// Event Notifications: Set up event notifications to trigger actions or notifications based on specific events in S3 (e.g., object created, deleted).
// Developer Experience Enhancements
// Error Handling: Robust error handling mechanisms to manage and troubleshoot issues during operations.
// SDK Documentation and Examples: Comprehensive documentation and examples are provided by AWS to help developers understand and implement various operations.
