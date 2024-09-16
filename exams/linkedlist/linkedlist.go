/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。
如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。
如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/

package linkedlist

import "github.com/go-demo/exams/linkedlist/listnode"

type MyLinkedList struct {
	dummy *listnode.ListNode
	size  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &listnode.ListNode{},
	}
}

func (l *MyLinkedList) Get(index int) int {
	if l.size == 0 || index < 0 || index > l.size-1 {
		return -1
	}

	cur := l.dummy
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	tmp := l.dummy.Next

	toAdd := &listnode.ListNode{
		Val:  val,
		Next: tmp,
	}

	l.dummy.Next = toAdd
	l.size++

	return
}

func (l *MyLinkedList) AddAtTail(val int) {
	toAdd := &listnode.ListNode{
		Val:  val,
		Next: nil,
	}

	cur := l.dummy
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = toAdd
	l.size++

	return
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		l.AddAtHead(val)
		return
	}

	if index == l.size {
		l.AddAtTail(val)
		return
	}

	if index > l.size {
		return
	}

	toAdd := &listnode.ListNode{
		Val: val,
	}

	cur := l.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	tmp := cur.Next

	cur.Next = toAdd
	toAdd.Next = tmp
	l.size++

	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}

	cur := l.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
		if cur == nil {
			return
		}
	}

	tmp := cur.Next.Next

	cur.Next = tmp
	l.size--

	return
}
