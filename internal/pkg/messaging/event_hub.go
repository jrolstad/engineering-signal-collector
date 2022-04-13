package messaging

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
)

type EventHub interface {
	Send(toSend *SignalEvent, target string) error
}

func NewEventHub(appConfig *config.AppConfig) EventHub {
	return NewSnsEventHub(appConfig)
}
