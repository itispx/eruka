package api

import (
	"encoding/json"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
)

type Status struct {
	Code int  `json:"code"`
	Ok   bool `json:"ok"`
}

var Headers = map[string]string{
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Headers": "*",
}

func CreateAPIGatewayProxyResponse(resp interface{}) (events.APIGatewayProxyResponse, error) {
	var statusCode int

	// Use reflection to inspect the type of 'resp'
	val := reflect.ValueOf(resp)

	// Check if the value is a struct or a pointer to a struct
	if val.Kind() == reflect.Ptr {
		// If it's a pointer, get the element it points to
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		statusCode = 500
	}

	// Try to find and access the 'Status' field
	statusField := val.FieldByName("Status")
	if !statusField.IsValid() {
		statusCode = 500
	}

	// Within 'Status', try to access the 'Code' field
	codeField := statusField.FieldByName("Code")
	if !codeField.IsValid() {
		statusCode = 500
	}

	// Check if the 'Code' field is of type int
	if codeField.Kind() != reflect.Int {
		statusCode = 500
	}

	statusCode = int(codeField.Int())

	msg, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    Headers,
		Body:       string(msg),
	}, nil
}
