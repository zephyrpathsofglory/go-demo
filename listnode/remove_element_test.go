package listnode_test

import (
	"testing"

	"github.com/go-demo/listnode"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
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

		res := listnode.RemoveElements(head, 1)

		nums := make([]int, 0, 4)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{2, 3, 4})
	}

	{
		head := &listnode.ListNode{
			Val: 1,
			Next: &listnode.ListNode{
				Val: 2,
				Next: &listnode.ListNode{
					Val: 1,
					Next: &listnode.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}

		res := listnode.RemoveElements(head, 1)

		nums := make([]int, 0, 4)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{2, 4})
	}

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

		res := listnode.RemoveElements(head, 2)

		nums := make([]int, 0, 4)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{1, 3, 4})
	}

	{
		head := &listnode.ListNode{
			Val: 1,
			Next: &listnode.ListNode{
				Val: 2,
				Next: &listnode.ListNode{
					Val: 3,
					Next: &listnode.ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		}

		res := listnode.RemoveElements(head, 2)

		nums := make([]int, 0, 4)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{1, 3})
	}
}
