package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
)

func SendSignal(messageHub messaging.MessageHub, toSend *messaging.SignalMessage) error {
	err := messageHub.Send(toSend, messaging.Queue_engineeringsignal_input)

	return err
}

func ProcessSignal(eventHub messaging.EventHub, signal *messaging.SignalMessage) error {

	event := messaging.MapToSignalEvent(signal)
	sendError := eventHub.Send(event, messaging.Topic_engineeringsignal_raw)

	return sendError
}
