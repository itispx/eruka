package rds

import (
	"github.com/itispx/eruka/aws/session"

	"github.com/aws/aws-sdk-go/service/rds"
)

var (
	RDSSession = rds.New(session.New())
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
