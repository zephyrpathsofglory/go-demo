package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	heights := []int{1, 3, 4, 5, 2, 7, 4}
	assert.Equal(t, 12, LargestRectangleArea(heights))
}
