package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSolutions(t *testing.T) {
	assert.Equal(t, 6, countSolutions([]int{12, 13, 20}, []int{10, 8, 7}))
	assert.Equal(t, 1, countSolutions([]int{11, 8, 20}, []int{10, 13, 7}))
	assert.Equal(t, 2, countSolutions([]int{11, 12, 20}, []int{10, 13, 7}))
	assert.Equal(t, 4, countSolutions([]int{19, 12, 20}, []int{10, 13, 7}))
	assert.Equal(t, 6, countSolutions([]int{11, 12, 20}, []int{30, 43, 57}))
	assert.Equal(t, 2, countSolutions([]int{11, 32, 20}, []int{30, 43, 57}))
}
