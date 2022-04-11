package collection_agent_github

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"time"
)

func MapPullRequestToSignalMessage(body string) *messaging.SignalMessage {
	message := &messaging.SignalMessage{
		SourceName: "GitHub",
		Source:     "GitHub",
		ReceivedAt: time.Now(),
		ObjectType: models.ObjectType_PeerReview,
		Data:       body,
	}
	return message
}
