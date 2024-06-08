// 25. Reverse Nodes in k-Group

package rnikg

type Node struct {
	val  int
	next *Node
}

func ReverseNodesInKGroup(initNode *Node, k int) *Node {
	if initNode == nil {
		return nil
	}

	// return if len < k
	last := initNode
	for i := 0; i < k-1; i++ {
		if last.next == nil {
			return initNode
		}

		last = last.next
	}

	nextRoundHead := last.next

	reversedFirstNode := reverseNodes(initNode, k)

	n := ReverseNodesInKGroup(nextRoundHead, k)
	initNode.next = n

	return reversedFirstNode
}

func reverseNodes(initNode *Node, k int) *Node {
	if k == 1 {
		return initNode
	}

	next := initNode.next

	n := reverseNodes(next, k-1)
	next.next = initNode

	return n
}
