package messaging

import (
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsEventHub struct {
	sns *sns.SNS
}

func (hub *SnsEventHub) Send(toSend *SignalEvent, target string) error {
	message, mappingError := MapToSnsPublishMessage(hub.sns, toSend, target)
	if mappingError != nil {
		return mappingError
	}

	_, publishError := hub.sns.Publish(message)

	return publishError
}
