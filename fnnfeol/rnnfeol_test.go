package rnnfeol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	n := constructNodes()

	n1 := RemoveNthNodeFromEndOfList(n, 2)

	assert.Equal(t, 1, n1.val)
	assert.Equal(t, 2, n1.next.val)
	assert.Equal(t, 3, n1.next.next.val)
	assert.Equal(t, 5, n1.next.next.next.val)
	assert.Nil(t, n1.next.next.next.next)

	n2 := RemoveNthNodeFromEndOfList(n1, 1)
	assert.Equal(t, 1, n2.val)
	assert.Equal(t, 2, n2.next.val)
	assert.Equal(t, 3, n2.next.next.val)
	assert.Nil(t, n2.next.next.next)

	n3 := RemoveNthNodeFromEndOfList(n2, 3)
	assert.Equal(t, 2, n3.val)
	assert.Equal(t, 3, n3.next.val)
	assert.Nil(t, n3.next.next)
}

func constructNodes() *Node {
	node5 := &Node{
		val:  5,
		next: nil,
	}

	node4 := &Node{
		val:  4,
		next: node5,
	}

	node3 := &Node{
		val:  3,
		next: node4,
	}

	node2 := &Node{
		val:  2,
		next: node3,
	}

	node1 := &Node{
		val:  1,
		next: node2,
	}

	return node1
}
