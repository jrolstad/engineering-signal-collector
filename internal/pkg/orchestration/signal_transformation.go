package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/transformation"
)

func TransformSignal(eventHub messaging.EventHub, event *models.SignalEvent) error {

	result := transformation.TransformSignal(event.ObjectType, event.SourceName, event.Data)

	if result == nil {
		return nil
	}

	standardEvent := MapTransformedToSignalEvent(event, result)
	if eventHub == nil {
		return nil
	}
	sendError := eventHub.Send(standardEvent, messaging.Topic_engineeringsignal_standardized)

	return sendError
}

func MapTransformedToSignalEvent(originalEvent *models.SignalEvent, data interface{}) *models.SignalEvent {
	dataAsJson := core.MapToJson(data)
	event := &models.SignalEvent{
		SourceName: originalEvent.SourceName,
		Source:     originalEvent.Source,
		ReceivedAt: models.GetCurrentTime(),
		ObjectType: originalEvent.ObjectType,
		ObjectId:   originalEvent.ObjectId,
		Data:       dataAsJson,
	}

	return event
}
