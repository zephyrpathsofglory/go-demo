package permute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermuteUnique(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}, permuteUnique([]int{1, 1, 2}))

	assert.Equal(t, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}, permuteUnique([]int{1, 2, 3}))
}
