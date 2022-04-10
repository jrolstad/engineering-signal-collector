package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func SendSignal(messageHub messaging.MessageHub, toSend *models.SignalMessage) error {
	err := messageHub.Send(toSend, messaging.Queue_engineeringsignal_input)

	return err
}

func ProcessSignal(eventHub messaging.EventHub, signal *models.SignalMessage) error {
	event := MapToSignalEvent(signal)

	sendError := eventHub.Send(event, messaging.Topic_engineeringsignal_raw)

	return sendError
}
