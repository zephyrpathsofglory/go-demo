package listnode_test

import (
	"testing"

	"github.com/go-demo/listnode"
	"github.com/stretchr/testify/assert"
)

func TestAddAtHead(t *testing.T) {
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

		res := listnode.AddAtHead(head, 5)

		nums := make([]int, 0, 5)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{5, 1, 2, 3, 4})
	}

	{
		var head *listnode.ListNode
		res := listnode.AddAtHead(head, 1)

		nums := make([]int, 0, 1)
		cur := res
		for cur != nil {
			nums = append(nums, cur.Val)
			cur = cur.Next
		}

		assert.Equal(t, nums, []int{1})
	}
}
