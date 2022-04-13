package data

import "github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"

type BlobService interface {
	Save(toSend *messaging.SignalEvent, target string, path string) error
}

func NewBlobService() BlobService {
	return NewS3BlobService()
}
