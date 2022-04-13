package orchestration

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/models"
	"testing"
)

func Test_MeasurePolicyAdherence_PeerReview_PolicyResultEvent(t *testing.T) {
	data := `{"SourceName":"GitHub","Source":"GitHub","ReceivedAt":"2022-04-13T20:27:51.434149846Z","ObjectType":"peer_review","ObjectId":"7d027e3d-27df-5830-2387-80f16341ad84","Data":"{\"Source\":\"GitHub\",\"Id\":\"909330077\",\"Title\":\"Update README.md\",\"State\":\"open\",\"CreatedAt\":\"2022-04-13T20:27:50Z\",\"ClosedAt\":\"0001-01-01T00:00:00Z\",\"Approvers\":[],\"Repository\":{\"Id\":\"480076219\",\"Name\":\"jrolstad/engineering-signal-collector\"}}"}`
	event := &models.SignalEvent{}
	core.MapFromJson(data, event)

	MeasurePolicyAdherence(nil, event)

}
