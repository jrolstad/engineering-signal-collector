package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func init() {
	fmt.Println("Initializing")
}

func main() {
	runtime.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//see https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}
