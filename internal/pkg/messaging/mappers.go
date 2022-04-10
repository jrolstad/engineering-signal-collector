package messaging

import (
	"encoding/json"
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
