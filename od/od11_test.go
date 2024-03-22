package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeBroadcastCount(t *testing.T) {
	assert.Equal(t, 3, computeBroadcastCount(
		[][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	))

	assert.Equal(t, 1, computeBroadcastCount(
		[][]int{
			{1, 1},
			{1, 1},
		},
	))
}
