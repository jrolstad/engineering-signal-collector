package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/data"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
)

func SaveSignal(blobService data.BlobService, event *messaging.SignalEvent) error {
	saveError := blobService.Save(event, data.Blob_engineeringsignal_data, data.Directory_engineeringsignal_raw)
	return saveError
}
