package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevert(t *testing.T) {
	{
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}

		res := Revert(head)

		nums := make([]int, 0, 5)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{4, 3, 2, 1})
	}

	{
		{
			head := &ListNode{
				Val:  1,
				Next: nil,
			}

			res := Revert(head)

			nums := make([]int, 0, 5)
			cur := res
			for cur != nil {
				nums = append(nums, cur.Val)
				cur = cur.Next
			}

			assert.Equal(t, nums, []int{1})
		}
	}

	{
		{
			{
				var head *ListNode

				res := Revert(head)

				nums := make([]int, 0, 5)
				cur := res
				for cur != nil {
					nums = append(nums, cur.Val)
					cur = cur.Next
				}

				assert.Equal(t, nums, []int{})
			}
		}
	}
}
