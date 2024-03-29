package textract

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/itispx/eruka/aws/session"
)

var (
	TextractSession = textract.New(session.New())
	RAW             = "RAW"
	QUERIES         = textract.FeatureTypeQueries
	FORMS           = textract.FeatureTypeForms
	TABLES          = textract.FeatureTypeTables
	LAYOUT          = textract.FeatureTypeLayout
	SIGNATURES      = textract.FeatureTypeSignatures
)

type Block struct {
	textract.Block
}

func StartDocumentAnalysis(bucket *string, name *string, s3OutputPrefix *string, featureTypes *[]string, topicArn *string, roleArn *string) (*textract.StartDocumentAnalysisOutput, error) {
	if os.Getenv("ENV") == "dev" {
		return TextractSession.StartDocumentAnalysis(&textract.StartDocumentAnalysisInput{
			DocumentLocation: &textract.DocumentLocation{
				S3Object: &textract.S3Object{
					Bucket: bucket,
					Name:   name,
				},
			},
			FeatureTypes: aws.StringSlice(*featureTypes),
			OutputConfig: &textract.OutputConfig{
				S3Bucket: bucket,
				S3Prefix: s3OutputPrefix,
			},
		})
	} else {
		return TextractSession.StartDocumentAnalysis(&textract.StartDocumentAnalysisInput{
			DocumentLocation: &textract.DocumentLocation{
				S3Object: &textract.S3Object{
					Bucket: bucket,
					Name:   name,
				},
			},
			FeatureTypes: aws.StringSlice(*featureTypes),
			OutputConfig: &textract.OutputConfig{
				S3Bucket: bucket,
				S3Prefix: s3OutputPrefix,
			},
			NotificationChannel: &textract.NotificationChannel{
				SNSTopicArn: topicArn,
				RoleArn:     roleArn,
			},
		})
	}
}

func GetDocumentAnalysis(jobId *string, paginationToken *string) (*textract.GetDocumentAnalysisOutput, error) {
	analysisInput := &textract.GetDocumentAnalysisInput{
		JobId: jobId,
	}

	if *paginationToken != "" {
		analysisInput.NextToken = paginationToken
	}

	analysisOutput, err := TextractSession.GetDocumentAnalysis(analysisInput)
	if err != nil {
		return nil, err
	}

	return analysisOutput, nil
}

func StartDocumentTextDetection(bucket *string, name *string, s3OutputPrefix *string, topicArn *string, roleArn *string) (*textract.StartDocumentTextDetectionOutput, error) {
	if os.Getenv("ENV") == "dev" {
		return TextractSession.StartDocumentTextDetection(&textract.StartDocumentTextDetectionInput{
			DocumentLocation: &textract.DocumentLocation{
				S3Object: &textract.S3Object{
					Bucket: bucket,
					Name:   name,
				},
			},
			OutputConfig: &textract.OutputConfig{
				S3Bucket: bucket,
				S3Prefix: s3OutputPrefix,
			},
		})
	} else {
		return TextractSession.StartDocumentTextDetection(&textract.StartDocumentTextDetectionInput{
			DocumentLocation: &textract.DocumentLocation{
				S3Object: &textract.S3Object{
					Bucket: bucket,
					Name:   name,
				},
			},
			OutputConfig: &textract.OutputConfig{
				S3Bucket: bucket,
				S3Prefix: s3OutputPrefix,
			},
			NotificationChannel: &textract.NotificationChannel{
				SNSTopicArn: topicArn,
				RoleArn:     roleArn,
			},
		})
	}
}

func GetDocumentTextDetection(jobId *string, paginationToken *string) (*textract.GetDocumentTextDetectionOutput, error) {
	analysisInput := &textract.GetDocumentTextDetectionInput{
		JobId: jobId,
	}

	if *paginationToken != "" {
		analysisInput.NextToken = paginationToken
	}

	analysisOutput, err := TextractSession.GetDocumentTextDetection(analysisInput)
	if err != nil {
		return nil, err
	}

	return analysisOutput, nil
}
