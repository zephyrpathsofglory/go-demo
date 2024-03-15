package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	{
		i := shortestPathLength([][]int{{1, 2, 3}, {0}, {0}, {0}})
		assert.Equal(t, 4, i)
	}

	{
		i := shortestPathLength([][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}})
		assert.Equal(t, 4, i)
	}

	// {
	// 	i := shortestPathLength2([][]int{{1, 2, 3}, {0}, {0}, {0}})
	// 	assert.Equal(t, 4, i)
	// }

	// {
	// 	i := shortestPathLength2([][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}})
	// 	assert.Equal(t, 4, i)
	// }
}

func TestN(t *testing.T) {
	// {
	// 	i := shortestPathLength2([][]int{{2, 3}, {7}, {0, 6}, {0, 4, 7}, {3, 8}, {7}, {2}, {5, 3, 1}, {4}})
	// 	assert.Equal(t, 11, i)
	// }

	{
		i := shortestPathLength([][]int{{2, 3}, {7}, {0, 6}, {0, 4, 7}, {3, 8}, {7}, {2}, {5, 3, 1}, {4}})
		assert.Equal(t, 11, i)
	}
}

// func TestP(t *testing.T) {
// 	{
// 		i := shortestPathLength2(
// 			[][]int{
// 				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
// 				{0, 2, 5, 6, 8},
// 				{0, 1, 4, 5, 6, 9, 10, 11},
// 				{0, 4, 5, 6, 8, 9, 10, 11},
// 				{0, 2, 3, 5, 6, 8, 10},
// 				{0, 1, 2, 3, 4, 6, 8, 9, 10, 11},
// 				{0, 1, 2, 3, 4, 5, 8, 10, 11},
// 				{0, 8},
// 				{0, 1, 3, 4, 5, 6, 7, 9, 10, 11},
// 				{0, 2, 3, 5, 8, 10},
// 				{0, 2, 3, 4, 5, 6, 8, 9},
// 				{0, 2, 3, 5, 6, 8},
// 			},
// 		)
// 		assert.Equal(t, 11, i)
// 	}
// }

func TestK(t *testing.T) {
	{
		i := shortestPathLength([][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			{0, 2, 5, 6, 8},
			{0, 1, 4, 5, 6, 9, 10, 11},
			{0, 4, 5, 6, 8, 9, 10, 11},
			{0, 2, 3, 5, 6, 8, 10},
			{0, 1, 2, 3, 4, 6, 8, 9, 10, 11},
			{0, 1, 2, 3, 4, 5, 8, 10, 11},
			{0, 8},
			{0, 1, 3, 4, 5, 6, 7, 9, 10, 11},
			{0, 2, 3, 5, 8, 10},
			{0, 2, 3, 4, 5, 6, 8, 9},
			{0, 2, 3, 5, 6, 8},
		},
		)
		assert.Equal(t, 11, i)
	}
}
