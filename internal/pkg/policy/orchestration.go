package policy

import (
	"fmt"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"strings"
)

func GetPolicyDefinitions(objectType string) []*PolicyDefinition {
	if objectType == models.ObjectType_PeerReview {
		policies := make([]*PolicyDefinition, 1)

		// Stub out 1 policy for peer reviews
		policies[0] = getSamplePRPolicy()

		return policies
	}

	return make([]*PolicyDefinition, 0)
}

func getSamplePRPolicy() *PolicyDefinition {
	return &PolicyDefinition{
		Id:      "1",
		Name:    "Peer Review Has Description",
		Version: "v1",
		IsValid: func(toMeasure interface{}) (bool, string) {
			data := &models.PeerReview{}
			mappingError := core.MapFromJson(fmt.Sprint(toMeasure), data)

			if mappingError != nil {
				return false, "Unable to parse"
			}

			if strings.Contains(data.Title, "Pass") {
				return true, ""
			}

			return false, "Pass must be in the PR title"

		},
	}
}

func MeasurePolicy(objectType string, objectId string, item interface{}, policy *PolicyDefinition) *PolicyResult {
	isValid, reason := policy.IsValid(item)

	return &PolicyResult{
		PolicyId:       policy.Id,
		PolicyName:     policy.Name,
		PolicyVersion:  policy.Version,
		MeasuredAt:     models.GetCurrentTime(),
		Result:         isValid,
		ResultReason:   reason,
		ObjectType:     objectType,
		ObjectId:       objectId,
		ObjectMeasured: item,
	}
}
