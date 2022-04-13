package main

import (
	"github.com/hashicorp/go-uuid"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	models_github "github.com/jrolstad/engineering-signal-collector/internal/pkg/models/github"
	"strings"
)

const (
	githubeventPullRequest = "pull_request"
)

func MapGithubEventToType(event string) string {
	if event == "" {
		return "unknown"
	}
	switch strings.ToLower(event) {
	case models_github.ObjectType_PullRequest:
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
		Data:       core.EncodeString(body),
	}
	return message
}
