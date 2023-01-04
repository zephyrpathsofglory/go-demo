package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	current := dummy

	for current != nil && dummy.Next != nil {
		if current.Next == nil {
			break
		}

		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}

func Get(head *ListNode, index int) int {
	if head == nil {
		return -1
	}

	if index == 0 {
		return head.Val
	}

	return Get(head.Next, index-1)
}

func AddAtHead(head *ListNode, val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: head,
	}
}

func AddAtTail(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &ListNode{
		Val: val,
	}

	return dummy.Next
}

func AddAtIndex(head *ListNode, index, val int) *ListNode {
	if index < 0 {
		return AddAtHead(head, val)
	}

	dummy := &ListNode{
		Next: head,
	}

	cur := dummy
	for i := 0; i <= index; i++ {
		if i == index {
			if cur == nil {
				return dummy.Next
			}

			var temp *ListNode
			if cur != nil {
				temp = cur.Next
			}

			cur.Next = &ListNode{
				Val:  val,
				Next: temp,
			}
		} else {
			if cur == nil {
				return dummy.Next
			}

			cur = cur.Next
		}
	}

	return dummy.Next
}

func DeleteAtIndex(head *ListNode, index int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy
	for i := 0; i <= index; i++ {
		if cur.Next == nil {
			return dummy.Next
		}

		if i == index {
			cur.Next = cur.Next.Next
			return dummy.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}

func Revert(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	var tmp *ListNode

	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

func SwitchPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next.Next
		tmp2 := cur.Next.Next
		cur.Next.Next.Next = cur.Next
		cur.Next.Next = tmp
		cur.Next = tmp2
		cur = cur.Next.Next
	}

	return dummy.Next
}
