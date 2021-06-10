package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	idx := BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9)
	assert.Equal(t, 4, idx)

	idx = BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 2)
	assert.Equal(t, -1, idx)

	idx = BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 12)
	assert.Equal(t, 5, idx)

	idx = BinarySearch([]int{-1, 0, 3, 5, 9, 12}, -1)
	assert.Equal(t, 0, idx)
}
