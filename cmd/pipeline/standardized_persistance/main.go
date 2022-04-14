package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/data"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
)

var (
	_blobService    data.BlobService
	_dataRepository data.Repository
)

func init() {
	_blobService = data.NewBlobService()
	_dataRepository = data.NewDataRepository()
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, snsEvent events.SNSEvent) error {

	for _, event := range snsEvent.Records {
		err := ProcessEvent(event)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessEvent(message events.SNSEventRecord) error {
	data := MapToSignalEvent(message)

	saveError := orchestration.SaveStandardizedModel(_blobService, _dataRepository, data)

	return saveError
}
