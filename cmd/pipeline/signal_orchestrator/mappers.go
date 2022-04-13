package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func MapToSignalMessage(message events.SQSMessage) *models.SignalMessage {
	data := &models.SignalMessage{}
	json.Unmarshal([]byte(message.Body), data)
	return data
}
