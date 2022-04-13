package messaging

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func NewSnsEventHub(appConfig *config.AppConfig) EventHub {
	hub := new(SnsEventHub)

	session := core.GetAwsSession()
	hub.sns = sns.New(session)

	return hub
}

type SnsEventHub struct {
	sns *sns.SNS
}

func (hub *SnsEventHub) Send(toSend *models.SignalEvent, target string) error {
	message, mappingError := MapToSnsPublishMessage(hub.sns, toSend, target)
	if mappingError != nil {
		return mappingError
	}

	_, publishError := hub.sns.Publish(message)

	return publishError
}
