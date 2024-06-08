package cwmw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainWithMostWater(t *testing.T) {
	res := ContainWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(t, 49, res)

	res = ContainWithMostWater([]int{1, 1})
	assert.Equal(t, 1, res)

	res = ContainWithMostWater([]int{4, 3, 2, 1, 4})
	assert.Equal(t, 16, res)

	res = ContainWithMostWater([]int{1, 2, 1})
	assert.Equal(t, 2, res)
}
