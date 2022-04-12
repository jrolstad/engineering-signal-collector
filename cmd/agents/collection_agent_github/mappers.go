package main

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"strings"
)

const (
	githubeventPullRequest = "pull_request"
)

func MapGithubEventToType(event string) string {
	switch strings.ToLower(githubeventPullRequest) {
	case models.ObjectType_PeerReview:
		return models.ObjectType_PeerReview
	default:
		return event
	}
}

func MapPullRequestToSignalMessage(objectType string, body string) *messaging.SignalMessage {
	message := &messaging.SignalMessage{
		SourceName: "GitHub",
		Source:     "GitHub",
		ReceivedAt: models.GetCurrentTime(),
		ObjectType: objectType,
		Data:       body,
	}
	return message
}
