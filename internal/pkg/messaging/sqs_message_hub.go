package messaging

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewSqsMessageHub() MessageHub {
	hub := new(SqsMessageHub)

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
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
