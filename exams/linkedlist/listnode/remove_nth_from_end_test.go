package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
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

		res := RemoveNthFromEnd(head, 2)

		nums := make([]int, 0, 5)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{1, 2, 4})
	}

	{
		{
			head := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			}

			res := RemoveNthFromEnd(head, 1)

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
			head := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			}

			res := RemoveNthFromEnd(head, 2)

			nums := make([]int, 0, 5)
			cur := res
			for cur != nil {
				nums = append(nums, cur.Val)
				cur = cur.Next
			}

			assert.Equal(t, nums, []int{2})
		}
	}

	{
		{
			head := &ListNode{
				Val:  1,
				Next: nil,
			}

			res := RemoveNthFromEnd(head, 1)

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
