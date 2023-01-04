package listnode_test

import (
	"testing"

	"github.com/go-demo/listnode"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
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

		assert.Equal(t, listnode.Get(head, 0), 1)
		assert.Equal(t, listnode.Get(head, 1), 2)
		assert.Equal(t, listnode.Get(head, 2), 3)
		assert.Equal(t, listnode.Get(head, 3), 4)
		assert.Equal(t, listnode.Get(head, 4), -1)
		assert.Equal(t, listnode.Get(head, 5), -1)
		assert.Equal(t, listnode.Get(head, 6), -1)
	}

	{
		var head *listnode.ListNode

		assert.Equal(t, listnode.Get(head, 0), -1)
		assert.Equal(t, listnode.Get(head, 1), -1)
		assert.Equal(t, listnode.Get(head, 2), -1)
		assert.Equal(t, listnode.Get(head, 3), -1)
		assert.Equal(t, listnode.Get(head, 4), -1)
		assert.Equal(t, listnode.Get(head, 5), -1)
		assert.Equal(t, listnode.Get(head, 6), -1)

	}
}
