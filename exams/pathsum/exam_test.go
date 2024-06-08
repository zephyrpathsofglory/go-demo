package pathsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExam(t *testing.T) {
	{
		root := &Node{
			val: 1,
			left: &Node{
				val: 2,
				left: &Node{
					val: 3,
				},
				right: &Node{
					val: 4,
				},
			},
			right: &Node{
				val: 3,
				left: &Node{
					val: 2,
				},
				right: &Node{
					val: 2,
				},
			},
		}

		assert.Equal(t, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{1, 3, 2},
		}, pathSum(root, 6))
	}

	{
		root := &Node{
			val: 1,
			left: &Node{
				val: 2,
				left: &Node{
					val: 0,
				},
			},
			right: &Node{
				val: 4,
				left: &Node{
					val: -2,
				},
				right: &Node{
					val: 2,
					right: &Node{
						val: -4,
					},
				},
			},
		}

		assert.Equal(t, [][]int{
			{1, 2, 0},
			{1, 4, -2},
			{1, 4, 2, -4},
		}, pathSum(root, 3))
	}
}
