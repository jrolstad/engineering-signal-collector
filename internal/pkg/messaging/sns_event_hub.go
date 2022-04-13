package messaging

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/config"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
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

func (hub *SnsEventHub) Send(toSend *SignalEvent, target string) error {
	fmt.Println("Mapping event...")
	message, mappingError := MapToSnsPublishMessage(hub.sns, toSend, target)
	if mappingError != nil {
		fmt.Println("Event Mapping Error: " + mappingError.Error())
		return mappingError
	}
	fmt.Println("Event mapped")

	fmt.Println("Publishing Event...")
	_, publishError := hub.sns.Publish(message)
	if publishError != nil {
		fmt.Println("Event Publishing Error: " + publishError.Error())
	}
	fmt.Println("Event published")
	return publishError
}
