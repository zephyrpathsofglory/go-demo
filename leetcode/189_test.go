package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test189(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(input, k)

	assert.Equal(t, input, []int{5, 6, 7, 1, 2, 3, 4})
}
