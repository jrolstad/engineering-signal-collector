package main

import (
	"github.com/hashicorp/go-uuid"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"strings"
)

const (
	githubeventPullRequest = "pull_request"
)

func MapGithubEventToType(event string) string {
	if event == "" {
		return "unknown"
	}
	switch strings.ToLower(githubeventPullRequest) {
	case models.ObjectType_PeerReview:
		return models.ObjectType_PeerReview
	default:
		return event
	}
}

func MapToSignalMessage(objectType string, body string) *models.SignalMessage {
	objectId, _ := uuid.GenerateUUID()

	message := &models.SignalMessage{
		SourceName: "GitHub",
		Source:     "GitHub",
		ReceivedAt: models.GetCurrentTime(),
		ObjectType: objectType,
		ObjectId:   objectId,
		Data:       body,
	}
	return message
}
