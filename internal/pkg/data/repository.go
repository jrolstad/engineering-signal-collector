package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/policy"
)

type Repository interface {
	SavePolicyResult(id string, data *policy.PolicyResult) error
	SaveStandardDataModel(id string, data *map[string]interface{}) error
}

func NewDataRepository() Repository {
	repository := new(DynamoDbRepository)

	session := core.GetAwsSession()
	repository.client = dynamodb.New(session)

	return repository
}

const (
	dynamodbTable_PolicyResult = "engineering_signal_collector_prd_policyresult"
	dynamodbTable_StandardData = "engineering_signal_collector_prd_standarddata"
)

type DynamoDbRepository struct {
	client *dynamodb.DynamoDB
}

func (repository *DynamoDbRepository) SavePolicyResult(id string, data *policy.PolicyResult) error {
	attributes, _ := dynamodbattribute.MarshalMap(data)

	idValue, _ := dynamodbattribute.Marshal(id)
	attributes["id"] = idValue

	input := &dynamodb.PutItemInput{
		Item:      attributes,
		TableName: aws.String(dynamodbTable_PolicyResult),
	}

	_, putError := repository.client.PutItem(input)

	return putError
}

func (repository *DynamoDbRepository) SaveStandardDataModel(id string, data *map[string]interface{}) error {
	attributes, _ := dynamodbattribute.MarshalMap(data)

	idValue, _ := dynamodbattribute.Marshal(id)
	attributes["id"] = idValue

	input := &dynamodb.PutItemInput{
		Item:      attributes,
		TableName: aws.String(dynamodbTable_StandardData),
	}

	_, putError := repository.client.PutItem(input)

	return putError
}
