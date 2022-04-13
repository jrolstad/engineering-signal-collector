package transformation

import (
	"fmt"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	models_github "github.com/jrolstad/engineering-signal-collector/internal/pkg/models/github"
)

func TransformSignal(objectType string, objectSource string, data string) interface{} {
	switch objectType {
	case models.ObjectType_PeerReview:
		return MapPeerReview(objectSource, data)
	default:
		return nil
	}
}

func MapPeerReview(objectSource string, toMap string) models.PeerReview {
	data := &models_github.PeerReviewEvent{}

	core.MapFromJson(toMap, data)

	result := models.PeerReview{
		Source:    objectSource,
		Id:        fmt.Sprint(data.PullRequest.ID),
		Title:     data.PullRequest.Title,
		CreatedAt: data.PullRequest.CreatedAt,
		ClosedAt:  data.PullRequest.ClosedAt,
		Approvers: []string{},
	}

	return result
}
