package orchestration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetApplicationHealth_ReturnsValues(t *testing.T) {
	result := GetApplicationHealth()

	assert.NotNil(t, result, "No result returned")

}
