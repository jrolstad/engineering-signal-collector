package main

import (
	"context"
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

func handleRequest(ctx context.Context) error {
	pollError := PollForLatestChangeLogs(_messageHub)
	return pollError
}
