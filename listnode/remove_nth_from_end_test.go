package listnode_test

import (
	"testing"

	"github.com/go-demo/listnode"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	{
		head := &listnode.ListNode{
			Val: 1,
			Next: &listnode.ListNode{
				Val: 2,
				Next: &listnode.ListNode{
					Val: 3,
					Next: &listnode.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}

		res := listnode.RemoveNthFromEnd(head, 2)

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
			head := &listnode.ListNode{
				Val: 1,
				Next: &listnode.ListNode{
					Val:  2,
					Next: nil,
				},
			}

			res := listnode.RemoveNthFromEnd(head, 1)

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
			head := &listnode.ListNode{
				Val: 1,
				Next: &listnode.ListNode{
					Val:  2,
					Next: nil,
				},
			}

			res := listnode.RemoveNthFromEnd(head, 2)

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
			head := &listnode.ListNode{
				Val:  1,
				Next: nil,
			}

			res := listnode.RemoveNthFromEnd(head, 1)

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
