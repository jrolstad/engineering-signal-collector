package main

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
)

func ProcessSignal(eventType string, body string, messageHub messaging.MessageHub) error {
	objectType := MapGithubEventToType(eventType)
	message := MapPullRequestToSignalMessage(objectType, body)

	sendError := orchestration.SendSignal(messageHub, message)

	return sendError
}
