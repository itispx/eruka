package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/itispx/eruka/aws/cloudwatch"
	"github.com/itispx/eruka/aws/rds"
)

func handler() {
	id := "dbConnections"
	namespace := "AWS/RDS"
	metricName := "DatabaseConnections"
	dimensionName := "DBInstanceIdentifier"
	instanceIdentifier := os.Getenv("INSTANCE_IDENTIFIER")

	query := cloudwatch.MetricDataQuery(&id, &namespace, &metricName, &dimensionName, &instanceIdentifier)

	// Retrieve the metric
	now := time.Now()
	startTime := now.Add(-time.Minute * 1)

	resp, err := cloudwatch.GetMetricData(&startTime, &now, query)
	if err != nil {
		log.Println("Error getting metric:", err)
		return
	}

	shouldStopInstance := true

	// Print the results
	for _, result := range resp.MetricDataResults {
		log.Println("result:", result)
		if *result.Id == "dbConnections" {
			for _, value := range result.Values {
				if *value > 0 {
					shouldStopInstance = false
					break
				}
			}
		}
	}

	if shouldStopInstance {
		_, err := rds.StopDBInstance(&instanceIdentifier)
		if err != nil && !strings.Contains(err.Error(), "InvalidDBInstanceState") {
			log.Println("Error stopping database instance:", err)
		}
	}
}

func main() {
	lambda.Start(handler)
}
