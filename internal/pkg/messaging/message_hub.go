package messaging

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

type MessageHub interface {
	Send(toSend *models.SignalMessage, target string) error
	Receive(receiver func(message *models.SignalMessage), target string) error
}

func NewMessageHub(appConfig *config.AppConfig) MessageHub {
	return NewSqsMessageHub(appConfig)
}
