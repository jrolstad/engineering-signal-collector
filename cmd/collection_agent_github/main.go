package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	appConfig "github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
)

var (
	messageHub messaging.MessageHub
)

func init() {
	config := appConfig.NewAppConfig()
	messageHub = messaging.NewMessageHub(config)
}

func main() {
	runtime.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	eventType := request.Headers["X-GitHub-Event"]

	switch eventType {
	case githubEvent_PullRequest:
		{
			ProcessPullRequest(request.Body)
			break
		}
	}
	//see https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func ProcessPullRequest(body string) error {
	message := MapPullRequestToSignalMessage(body)
	sendError := orchestration.SendSignal(messageHub, message)

	return sendError
}
