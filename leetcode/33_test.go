package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, -1, search([]int{1}, 0))
	assert.Equal(t, -1, search([]int{1, 3}, 4))
	assert.Equal(t, -1, search([]int{3, 1}, 0))
	assert.Equal(t, 2, search([]int{3, 5, 1}, 1))
}
