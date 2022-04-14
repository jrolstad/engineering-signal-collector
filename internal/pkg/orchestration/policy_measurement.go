package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/policy"
	"time"
)

func MeasurePolicyAdherence(eventHub messaging.EventHub, event *models.SignalEvent) error {
	policies := policy.GetPolicyDefinitions(event.ObjectType)

	for _, policyDefinition := range policies {
		result := policy.MeasurePolicy(event.ObjectType, event.ObjectId, event.Data, policyDefinition)
		policyResultEvent := MapPolicyResultToEvent(event, result)

		if eventHub != nil {
			eventHub.Send(policyResultEvent, messaging.Topic_engineeringsignal_policymeasured)
		}
	}

	return nil
}

func MapPolicyResultToEvent(originalEvent *models.SignalEvent, result *policy.PolicyResult) *models.SignalEvent {
	data := core.MapToJson(result)
	return &models.SignalEvent{
		SourceName: originalEvent.SourceName,
		Source:     originalEvent.Source,
		ReceivedAt: time.Time{},
		ObjectType: originalEvent.ObjectType,
		ObjectId:   originalEvent.ObjectId,
		Data:       data,
	}
}
