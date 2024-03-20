package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMazeWalker(t *testing.T) {
	{
		tc, uc := mazeWalker(6, 4, [][2]int{
			{0, 2},
			{1, 2},
			{2, 2},
			{4, 1},
			{5, 1},
		})
		assert.Equal(t, 2, tc)
		assert.Equal(t, 3, uc)
	}

	{
		{
			tc, uc := mazeWalker(6, 4, [][2]int{
				{2, 0},
				{2, 1},
				{3, 0},
				{3, 1},
			})
			assert.Equal(t, 0, tc)
			assert.Equal(t, 4, uc)
		}
	}
}
