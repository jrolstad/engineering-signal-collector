package messaging

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
)

func NewSqsMessageHub(appConfig *config.AppConfig) MessageHub {
	hub := new(SqsMessageHub)

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-west-1"),
		},
	}))
	hub.sqs = sqs.New(session)

	return hub
}

type SqsMessageHub struct {
	sqs *sqs.SQS
}

func (hub *SqsMessageHub) Send(toSend *SignalMessage, target string) error {
	fmt.Println("Mapping Message...")
	message, mapError := MapToSqsSendMessage(hub.sqs, toSend, target)
	fmt.Println("Message Mapped")
	if mapError != nil {
		return mapError
	}

	fmt.Println("Sending Message")
	_, sendError := hub.sqs.SendMessage(message)
	if sendError != nil {
		return sendError
	}
	fmt.Println("Message Sent")
	return nil
}

func (hub *SqsMessageHub) Receive(receiver func(message *SignalMessage), target string) error {
	return errors.New("not yet implemented")
}
