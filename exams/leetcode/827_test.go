package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestIsland(t *testing.T) {
	assert.Equal(t, 4, largestIsland([][]int{
		{1, 1},
		{1, 0},
	}))

	assert.Equal(t, 3, largestIsland([][]int{
		{1, 0},
		{0, 1},
	}))

	assert.Equal(t, 4, largestIsland([][]int{
		{1, 1},
		{1, 1},
	}))
}
