package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchPairs(t *testing.T) {
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

		res := SwitchPairs(head)

		nums := make([]int, 0, 5)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{2, 1, 4, 3})
	}

	{
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}

		res := SwitchPairs(head)

		nums := make([]int, 0, 5)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{2, 1, 3})
	}

	{
		{
			head := &ListNode{
				Val:  1,
				Next: nil,
			}

			res := SwitchPairs(head)

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

				res := SwitchPairs(head)

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
