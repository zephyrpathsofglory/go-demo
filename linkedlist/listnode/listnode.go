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

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	leftMost := dummy
	leftCur := dummy
	var rightCur, rightMost *ListNode

	for i := 0; i < left; i++ {
		if i == left-1 {
			leftCur = leftMost
		}
		leftMost = leftMost.Next
	}

	tmp1 := leftMost
	tmp2 := tmp1.Next
	for i := 0; i < right-left; i++ {
		tmp := tmp2.Next
		tmp2.Next = tmp1
		tmp1 = tmp2
		tmp2 = tmp
	}

	rightMost = tmp1
	rightCur = tmp2

	leftCur.Next = rightMost
	leftMost.Next = rightCur

	return dummy.Next
}
