package messaging

import (
	"github.com/Shopify/sarama"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

type EventHub interface {
	Send(toSend *models.SignalEvent, target string) error
}

func NewEventHub(appConfig *config.AppConfig) EventHub {
	hub := new(KafkaEventHub)

	config := sarama.NewConfig()

	producer, newError := sarama.NewSyncProducer(appConfig.KafkaBrokers, config)
	if newError != nil {
		return nil
	}

	hub.producer = producer

	return hub
}

type KafkaEventHub struct {
	producer sarama.SyncProducer
}

func (hub *KafkaEventHub) Send(toSend *models.SignalEvent, target string) error {
	message := MapToKafkaMessage(toSend)

	_, _, sendError := hub.producer.SendMessage(message)

	return sendError
}
