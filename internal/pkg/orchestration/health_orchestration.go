package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"time"
)

func GetApplicationHealth() *models.ApplicationHealth {
	health := &models.ApplicationHealth{
		Success:     true,
		CurrentTime: time.Now(),
	}

	return health
}
