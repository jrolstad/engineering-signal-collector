package messaging

import (
	"github.com/Shopify/sarama"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
)

func MapToSqsSendMessage(sqsInstance *sqs.SQS, toMap interface{}, queueName string) (*sqs.SendMessageInput, error) {
	urlResult, queueUrlErr := sqsInstance.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if queueUrlErr != nil {
		return nil, queueUrlErr
	}

	input := new(sqs.SendMessageInput)
	input.MessageBody = aws.String(core.MapToJson(toMap))
	input.QueueUrl = urlResult.QueueUrl

	return input, nil
}

func MapToSnsPublishMessage(snsInstance *sns.SNS, toMap interface{}, topicName string) (*sns.PublishInput, error) {
	topicInput := &sns.CreateTopicInput{
		Name: aws.String(topicName),
	}
	topicResult, topicError := snsInstance.CreateTopic(topicInput)
	if topicError != nil {
		return nil, topicError
	}

	input := new(sns.PublishInput)
	input.Message = aws.String(core.MapToJson(toMap))
	input.TopicArn = topicResult.TopicArn

	return input, nil
}

func MapToKafkaMessage(toMap interface{}, topicName string) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(core.MapToJson(toMap)),
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
		ObjectId:   toMap.ObjectId,
	}
}
