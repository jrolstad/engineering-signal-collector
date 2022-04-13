package messaging

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
)

type MessageHub interface {
	Send(toSend *SignalMessage, target string) error
	Receive(receiver func(message *SignalMessage), target string) error
}

func NewMessageHub(appConfig *config.AppConfig) MessageHub {
	return NewSqsMessageHub(appConfig)
}
