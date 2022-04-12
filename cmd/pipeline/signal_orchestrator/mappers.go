package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
)

func MapToSignalMessage(message events.SQSMessage) *messaging.SignalMessage {
	data := &messaging.SignalMessage{}
	json.Unmarshal([]byte(message.Body), data)
	return data
}
