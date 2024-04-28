package helpers

import (
	"github.com/aws/aws-lambda-go/events"
)

// Helper function to create a response
func Response(statusCode int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "OPTIONS,POST,GET,PUT,DELETE,PATCH",
			"Access-Control-Allow-Headers": "Content-Type,Authorization,X-API-Key,Content-Length,X-Amz-Date,X-Amz-Security-Token,X-Amz-User-Agent",
		},
		StatusCode: statusCode,
		Body:       body,
	}
}
