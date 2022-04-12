package models

import "time"

type ApplicationHealth struct {
	Success     bool
	CurrentTime time.Time
}
