package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	appConfig "github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
)

var (
	_eventHub messaging.EventHub
)

func init() {
	config := appConfig.NewAppConfig()
	_eventHub = messaging.NewEventHub(config)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, snsEvent events.SNSEvent) error {

	for _, event := range snsEvent.Records {
		err := ProcessEvent(event)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessEvent(message events.SNSEventRecord) error {
	data := MapToSignalEvent(message)

	saveError := orchestration.MeasurePolicyAdherence(_eventHub, data)

	return saveError
}
