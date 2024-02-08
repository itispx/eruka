package rds

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

var (
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	RDSSession = rds.New(sess)
)

func StopDBInstance(dbInstanceIdentifier *string) (*rds.StopDBInstanceOutput, error) {
	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: dbInstanceIdentifier,
	}

	result, err := RDSSession.StopDBInstance(input)

	return result, err

}

func StartDBInstance(dbInstanceIdentifier *string) (*rds.StartDBInstanceOutput, error) {
	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: dbInstanceIdentifier,
	}

	result, err := RDSSession.StartDBInstance(input)

	return result, err
}
