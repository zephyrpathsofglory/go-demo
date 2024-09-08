package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandPerimeter(t *testing.T) {
	assert.Equal(t, 16, islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))

	assert.Equal(t, 4, islandPerimeter([][]int{
		{0, 1},
	}))

	assert.Equal(t, 8, islandPerimeter([][]int{
		{1, 1},
		{1, 1},
	}))
}
