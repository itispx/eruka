package cloudwatch

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/itispx/eruka/aws/session"
)

var (
	CloudWatchSession = cloudwatch.New(session.New())
)

func MetricDataQuery(id *string, namespace *string, metricName *string, dimensionName *string, dimensionValue *string) *cloudwatch.MetricDataQuery {
	query := &cloudwatch.MetricDataQuery{
		Id: id,
		MetricStat: &cloudwatch.MetricStat{
			Metric: &cloudwatch.Metric{
				Namespace:  namespace,
				MetricName: metricName,
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  dimensionName,
						Value: dimensionValue,
					},
				},
			},
			Period: aws.Int64(60), // In seconds
			Stat:   aws.String("Average"),
		},
	}

	return query
}

func GetMetricData(startTime *time.Time, endTime *time.Time, query *cloudwatch.MetricDataQuery) (*cloudwatch.GetMetricDataOutput, error) {
	resp, err := CloudWatchSession.GetMetricData(&cloudwatch.GetMetricDataInput{
		StartTime:         aws.Time(*startTime),
		EndTime:           aws.Time(*endTime),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{query},
	})

	return resp, err
}
