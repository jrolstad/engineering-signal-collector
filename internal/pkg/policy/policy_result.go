package policy

import "time"

type PolicyResult struct {
	PolicyId       string
	PolicyName     string
	PolicyVersion  string
	MeasuredAt     time.Time
	Result         bool
	ResultReason   string
	ObjectType     string
	ObjectId       string
	ObjectMeasured interface{}
}
