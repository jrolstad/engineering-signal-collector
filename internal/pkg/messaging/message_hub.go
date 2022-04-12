package messaging

import (
	"errors"
	"fmt"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
)

type MessageHub interface {
	Send(toSend *SignalMessage, target string) error
	Receive(receiver func(message *SignalMessage), target string) error
}

func NewMessageHub(appConfig *config.AppConfig) MessageHub {
	return NewLoggingMessageHub(appConfig)
}

func NewLoggingMessageHub(appConfig *config.AppConfig) MessageHub {
	hub := new(LoggingMessageHub)

	return hub
}

type LoggingMessageHub struct {
}

func (hub *LoggingMessageHub) Send(toSend *SignalMessage, target string) error {
	data := core.MapToJson(toSend)

	fmt.Println("target:" + target + "|" + data)

	return nil
}

func (hub *LoggingMessageHub) Receive(receiver func(message *SignalMessage), target string) error {
	return errors.New("not yet implemented")
}
