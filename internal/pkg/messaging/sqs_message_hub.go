package messaging

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
)

func NewSqsMessageHub(appConfig *config.AppConfig) MessageHub {
	hub := new(SqsMessageHub)

	session := core.GetAwsSession()
	hub.sqs = sqs.New(session)

	return hub
}

type SqsMessageHub struct {
	sqs *sqs.SQS
}

func (hub *SqsMessageHub) Send(toSend *SignalMessage, target string) error {
	message, mapError := MapToSqsSendMessage(hub.sqs, toSend, target)
	if mapError != nil {
		return mapError
	}

	_, sendError := hub.sqs.SendMessage(message)
	if sendError != nil {
		return sendError
	}

	return nil
}

func (hub *SqsMessageHub) Receive(receiver func(message *SignalMessage), target string) error {
	return errors.New("not yet implemented")
}
