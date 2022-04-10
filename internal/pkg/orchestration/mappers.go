package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func MapToSignalEvent(toMap *models.SignalMessage) *models.SignalEvent {
	if toMap == nil {
		return nil
	}
	
	return &models.SignalEvent{
		SourceName: toMap.SourceName,
		Source:     toMap.Source,
		ReceivedAt: toMap.ReceivedAt,
		ObjectType: toMap.ObjectType,
		Data:       toMap.Data,
	}
}
