package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeightII(t *testing.T) {
	lw := LastStoneWeightII([]int{2, 7, 4, 1, 8, 1})
	assert.Equal(t, 1, lw)

	lw = LastStoneWeightII([]int{31, 26, 33, 21, 40})
	assert.Equal(t, 5, lw)

	lw = LastStoneWeightII([]int{1, 2})
	assert.Equal(t, 1, lw)
}
