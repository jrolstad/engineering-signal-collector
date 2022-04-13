package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func MapToSignalEvent(message events.SNSEventRecord) *models.SignalEvent {
	data := &models.SignalEvent{}
	json.Unmarshal([]byte(message.SNS.Message), data)
	return data
}
