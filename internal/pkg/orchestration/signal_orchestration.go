package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func SignalReceived(hub messaging.MessageHub, toSend *models.SignalMessage) error {
	err := hub.Send(toSend, models.Queue_engineeringsignal_input)

	return err
}
