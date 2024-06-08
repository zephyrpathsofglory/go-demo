package rnikg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseNodesInKGroup(t *testing.T) {
	node := constructNodes()

	n := ReverseNodesInKGroup(node, 2)
	assert.Equal(t, 2, n.val)
}

func TestReverseNodes(t *testing.T) {
	node := constructNodes()

	n := reverseNodes(node, 5)
	assert.Equal(t, 5, n.val)
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
