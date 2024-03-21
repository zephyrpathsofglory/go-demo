package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestNodePath(t *testing.T) {
	assert.Equal(t, []int{3, 7, 2}, smallestNodePath([]int{-1, 3, 5, 7, -1, -1, 2, 4}))
	assert.Equal(t, []int{5, 8, 7, 6}, smallestNodePath([]int{-1, 5, 9, 8, -1, -1, 7, -1, -1, -1, -1, -1, 6}))
}
