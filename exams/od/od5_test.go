package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskScheduler(t *testing.T) {
	assert.Equal(t, 4, taskScheduler([][2]int{
		{2, 2},
	}))

	assert.Equal(t, 4, taskScheduler([][2]int{
		{1, 1},
		{2, 2},
	}))

	assert.Equal(t, 7, taskScheduler([][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}))

	assert.Equal(t, 11, taskScheduler([][2]int{
		{1, 10},
		{2, 6},
		{3, 2},
	}))
}
