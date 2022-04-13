package models

import "time"

type PeerReview struct {
	Source    string
	Id        string
	Title     string
	State     string
	CreatedAt time.Time
	ClosedAt  time.Time
	Approvers []string
}
