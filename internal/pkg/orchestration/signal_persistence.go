package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/data"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func SaveSignal(blobService data.BlobService, event *models.SignalEvent) error {
	saveError := blobService.Save(event, data.Blob_engineeringsignal_data, data.Directory_engineeringsignal_raw)
	return saveError
}

func SaveStandardizedModel(blobService data.BlobService, event *models.SignalEvent) error {
	saveError := blobService.Save(event, data.Blob_engineeringsignal_data, data.Directory_engineeringsignal_standardized)
	return saveError
}
