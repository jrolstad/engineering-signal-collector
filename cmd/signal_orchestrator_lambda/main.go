package signal_orchestrator_lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	appConfig "github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/logging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
)

var (
	eventHub messaging.EventHub
)

func init() {
	config := appConfig.NewAppConfig()
	eventHub = messaging.NewEventHub(config)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	for _, message := range sqsEvent.Records {
		err := ProcessMessage(message)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessMessage(message events.SQSMessage) error {
	LogMessageReceived(message)

	data := MapToSignalMessage(message)

	processError := orchestration.ProcessSignal(eventHub, data)
	if processError != nil {
		return processError
	}

	return nil
}

func LogMessageReceived(message events.SQSMessage) {
	logging.LogMessagef("Message %s received from %s", message.MessageId, message.EventSource)
}
