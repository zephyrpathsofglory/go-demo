package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastCommitter(t *testing.T) {
	assert.Equal(t, 2, LeastCommitter([][]int{
		{0},
		{5, 2, 3},
		{0, 5},
	}))

	assert.Equal(t, 3, LeastCommitter([][]int{
		{1, 5, 2, 4},
		{2, 9, 1},
		{4},
		{9},
		{3, 1},
	}))
}
