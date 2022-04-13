package messaging

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

type EventHub interface {
	Send(toSend *models.SignalEvent, target string) error
}

func NewEventHub(appConfig *config.AppConfig) EventHub {
	return NewSnsEventHub(appConfig)
}
