package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTravelAll(t *testing.T) {
	assert.True(t, canTravelAll([][2]int{
		{1, 2},
		{1, 0},
	}))

	assert.True(t, canTravelAll([][2]int{
		{0, 2},
		{0, 1},
	}))

	assert.False(t, canTravelAll([][2]int{
		{0, 1},
		{1, 0},
	}))

	assert.True(t, canTravelAll([][2]int{
		{4, 3},
		{0, 4},
		{2, 1},
		{3, 2},
	}))

	assert.True(t, canTravelAll([][2]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
	}))

	assert.False(t, canTravelAll([][2]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 4},
		{0, 5},
		{4, 0},
	}))
}
