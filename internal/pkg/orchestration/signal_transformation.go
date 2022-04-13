package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func TransformSignal(eventHub messaging.EventHub, event *models.SignalEvent) error {
	return nil
}
