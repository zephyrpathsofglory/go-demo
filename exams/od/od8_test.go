package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterviewer(t *testing.T) {
	assert.Equal(t, 3, leastInterviewer([][2]int{
		{1, 3},
		{2, 5},
		{3, 9},
		{2, 8},
	}, 4, 4))

	assert.Equal(t, 4, leastInterviewer([][2]int{
		{1, 7},
		{3, 10},
		{2, 5},
		{4, 6},
	}, 3, 4))
}
