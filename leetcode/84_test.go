package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	{
		heights := []int{1, 3, 4, 5, 2, 7, 4}
		assert.Equal(t, 12, LargestRectangleArea(heights))
	}
	{
		heights := []int{1, 3, 4, 500, 2, 7, 4}
		assert.Equal(t, 500, LargestRectangleArea(heights))
	}
	{
		heights := []int{2, 4}
		assert.Equal(t, 4, LargestRectangleArea(heights))
	}
	{
		heights := []int{2, 1, 5, 6, 2, 3}
		assert.Equal(t, 10, LargestRectangleArea(heights))
	}
}
