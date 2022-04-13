package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
)

func MapToSignalEvent(message events.SNSEventRecord) *messaging.SignalEvent {
	data := &messaging.SignalEvent{}
	json.Unmarshal([]byte(message.SNS.Message), data)
	return data
}
