package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalSum(t *testing.T) {
	assert.Equal(t, []int{0, 2}, totalSum([]int{1, 1, 3, 2}, 4))
	assert.Equal(t, []int{0, 1}, totalSum([]int{1, 1}, 4))
	assert.Equal(t, []int{0}, totalSum([]int{1}, 4))
	assert.Equal(t, []int{0, 1}, totalSum([]int{1, 3, 2}, 4))
	assert.Equal(t, []int{0, 1}, totalSum([]int{1, 3, 3}, 4))
	assert.Equal(t, []int{0, 2}, totalSum([]int{1, 2, 3}, 4))
	assert.Equal(t, []int{}, totalSum([]int{}, 4))
	assert.Equal(t, []int{1, 2}, totalSum([]int{4, 1, 2}, 3))
}
