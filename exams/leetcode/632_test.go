package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRange(t *testing.T) {
	assert.Equal(t, []int{20, 24}, smallestRange([][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
		{5, 18, 22, 30},
	}))

	assert.Equal(t, []int{1, 1}, smallestRange([][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}))
}
