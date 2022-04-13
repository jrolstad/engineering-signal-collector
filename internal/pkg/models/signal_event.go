package models

import "time"

type SignalEvent struct {
	SourceName string
	Source     string
	ReceivedAt time.Time
	ObjectType string
	ObjectId   string
	Data       interface{}
}
