package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
)

func GetApplicationHealth() *models.ApplicationHealth {
	health := &models.ApplicationHealth{
		Success:     true,
		CurrentTime: models.GetCurrentTime(),
	}

	return health
}
