package data

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"io"
	"strings"
)

func NewS3BlobService() BlobService {
	service := new(S3BlobService)

	session := core.GetAwsSession()
	service.uploader = s3manager.NewUploader(session)
	return service
}

type S3BlobService struct {
	uploader *s3manager.Uploader
}

func (service *S3BlobService) Save(event *models.SignalEvent, target string, path string) error {
	key := getBlobKey(event)
	body := getBody(event)

	fileName := strings.Join([]string{path, key}, "/")
	content := &s3manager.UploadInput{
		Bucket: aws.String(target),
		Key:    aws.String(fileName),
		Body:   body,
	}
	_, uploadError := service.uploader.Upload(content)

	return uploadError
}

func getBlobKey(event *models.SignalEvent) string {
	key := fmt.Sprintf("%v_%v_%v", event.Source, event.ObjectType, event.ObjectId)

	return key
}

func getBody(event *models.SignalEvent) io.Reader {
	eventData := core.MapToJson(event)
	reader := strings.NewReader(eventData)

	return reader
}
