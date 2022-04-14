package main

import (
	"github.com/hashicorp/go-uuid"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func MapToSignalMessage(objectType string, source string, body string) *models.SignalMessage {
	objectId, _ := uuid.GenerateUUID()

	message := &models.SignalMessage{
		SourceName: "ChangeLog",
		Source:     source,
		ReceivedAt: models.GetCurrentTime(),
		ObjectType: objectType,
		ObjectId:   objectId,
		Data:       body,
	}
	return message
}
