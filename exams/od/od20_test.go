package od

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRentingSystem(t *testing.T) {
	{
		rs := New()

		assert.True(t, rs.addRoom(3, 24, 200, 2, []int{0, 1}))
		assert.True(t, rs.addRoom(1, 10, 400, 2, []int{1, 0}))
		assert.Equal(t, []int{1, 3}, rs.queryRoom(1, 400, 2, []int{1, 1}, [][]int{
			{3, 1},
			{2, -1},
		}))
		assert.True(t, rs.deleteRoom(3))
	}

	{
		rs := New()
		assert.False(t, rs.deleteRoom(10))
		assert.True(t, rs.addRoom(3, 24, 200, 2, []int{0, 1}))
		assert.False(t, rs.addRoom(3, 24, 500, 2, []int{0, 1}))
		assert.False(t, rs.addRoom(3, 27, 500, 4, []int{1, 1}))
		assert.True(t, rs.addRoom(1, 27, 500, 4, []int{20, 43}))
		assert.True(t, rs.addRoom(6, 35, 227, 4, []int{2, 4}))
		assert.True(t, rs.addRoom(9, 20, 3540, 4, []int{4, 321}))
		assert.Equal(t, []int{3, 1, 6}, rs.queryRoom(25, 900, 4, []int{10, 1}, [][]int{
			{1, 1},
			{2, -1},
			{3, 1},
		}))
	}
}
