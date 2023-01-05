package listnode_test

import (
	"testing"

	"github.com/go-demo/linkedlist/listnode"
	"github.com/stretchr/testify/assert"
)

func TestRevert(t *testing.T) {
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

		res := listnode.Revert(head)

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
			head := &listnode.ListNode{
				Val:  1,
				Next: nil,
			}

			res := listnode.Revert(head)

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
				var head *listnode.ListNode

				res := listnode.Revert(head)

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
