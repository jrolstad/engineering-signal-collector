package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	appConfig "github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
)

var (
	_messageHub messaging.MessageHub
)

func init() {
	config := appConfig.NewAppConfig()
	_messageHub = messaging.NewMessageHub(config)
}

func main() {
	runtime.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	eventType := request.Headers["X-GitHub-Event"]
	ProcessSignal(eventType, request.Body, _messageHub)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: request.Body}, nil
}
