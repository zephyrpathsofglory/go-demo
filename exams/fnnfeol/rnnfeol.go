// 19. Remove Nth Node From End of List

package rnnfeol

type Node struct {
	val  int
	next *Node
}

func RemoveNthNodeFromEndOfList(node *Node, n int) *Node {
	head := &Node{}
	mov := node
	head.next = node

	headMoved := 0
	step := 0
	for {
		mov = mov.next
		step++
		if step > n {
			headMoved++
			head = head.next
		}

		if mov == nil {
			break
		}
	}

	// the first one removed
	if headMoved == 0 {
		return head.next.next
	}

	toRemove := head.next
	nextOfToRemove := toRemove.next
	head.next = nextOfToRemove
	toRemove.next = nil

	return node
}
