package data

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

type BlobService interface {
	Save(toSend *models.SignalEvent, target string, path string) error
}

func NewBlobService() BlobService {
	return NewS3BlobService()
}
