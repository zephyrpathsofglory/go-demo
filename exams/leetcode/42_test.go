package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrappingRainWater(t *testing.T) {
	v := TrappingRainWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	assert.Equal(t, 6, v)

	v = TrappingRainWater([]int{4, 2, 0, 3, 2, 5})
	assert.Equal(t, 9, v)
}
