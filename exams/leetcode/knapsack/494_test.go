package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	ways := FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	assert.Equal(t, 5, ways)

	ways = FindTargetSumWays([]int{1}, 1)
	assert.Equal(t, 1, ways)

	ways = FindTargetSumWays([]int{1, 2, 2, 3, 5}, 3)
	assert.Equal(t, 4, ways)
}
