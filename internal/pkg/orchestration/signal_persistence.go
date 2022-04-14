package orchestration

import (
	"github.com/hashicorp/go-uuid"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/data"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/policy"
)

func SaveSignal(blobService data.BlobService, event *models.SignalEvent) error {
	saveError := blobService.Save(event, data.Blob_engineeringsignal_data, data.Directory_engineeringsignal_raw)
	return saveError
}

func SaveStandardizedModel(blobService data.BlobService, dataRepository data.Repository, event *models.SignalEvent) error {
	saveError := blobService.Save(event, data.Blob_engineeringsignal_data, data.Directory_engineeringsignal_standardized)

	modelData := make(map[string]interface{})
	core.MapFromJson(event.Data, &modelData)

	dataSaveError := dataRepository.SaveStandardDataModel(event.ObjectId, event.ObjectType, &modelData)

	if saveError != nil {
		return saveError
	}
	return dataSaveError
}

func SavePolicyResult(blobService data.BlobService, dataRepository data.Repository, event *models.SignalEvent) error {
	saveError := blobService.Save(event, data.Blob_engineeringsignal_data, data.Directory_engineeringpolicy_result)

	policyResult := &policy.PolicyResult{}
	core.MapFromJson(event.Data, policyResult)

	uniqueId, _ := uuid.GenerateUUID()
	dataSaveError := dataRepository.SavePolicyResult(uniqueId, policyResult)

	if saveError != nil {
		return saveError
	}
	return dataSaveError
}
