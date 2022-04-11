package messaging

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func MapToSqsSendMessage(sqsInstance *sqs.SQS, toMap interface{}, queueName string) (*sqs.SendMessageInput, error) {
	urlResult, queueUrlErr := sqsInstance.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if queueUrlErr != nil {
		return nil, queueUrlErr
	}

	input := new(sqs.SendMessageInput)
	input.MessageBody = aws.String(MapToJson(toMap))
	input.QueueUrl = urlResult.QueueUrl

	return input, nil
}

func MapToJson(toMap interface{}) string {
	result, err := json.Marshal(toMap)
	if err != nil {
		return "{}"
	}

	return string(result)
}

func MapToKafkaMessage(toMap interface{}, topicName string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(MapToJson(toMap)),
	}
}

func MapToSignalEvent(toMap *SignalMessage) *SignalEvent {
	if toMap == nil {
		return nil
	}

	return &SignalEvent{
		SourceName: toMap.SourceName,
		Source:     toMap.Source,
		ReceivedAt: toMap.ReceivedAt,
		ObjectType: toMap.ObjectType,
		Data:       toMap.Data,
	}
}
