// 146. LRU Cache

package lru

type node struct {
	pre  *node
	next *node
	key  int
	val  int
}

func newNode() *node {
	return &node{}
}

type doublyLinkedList struct {
	head *node
	tail *node
}

func newDLL() (dll *doublyLinkedList) {
	head := newNode()
	tail := newNode()

	head.next = tail
	tail.pre = head

	dll = &doublyLinkedList{
		head: head,
		tail: tail,
	}

	return
}

func (dll *doublyLinkedList) removeNode(node *node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (dll *doublyLinkedList) addNode(node *node) {
	node.next = dll.head.next
	node.pre = dll.head

	dll.head.next.pre = node
	dll.head.next = node
}

type Cache struct {
	capacity int
	size     int
	rlist    *doublyLinkedList
	nodeMap  map[int]*node
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		rlist:    newDLL(),
		nodeMap:  make(map[int]*node),
	}
}

func (c *Cache) incrRecent(node *node) {
	c.rlist.removeNode(node)
	c.rlist.addNode(node)
}

func (c *Cache) Get(key int) (val int) {
	n := c.nodeMap[key]
	if n == nil {
		return -1
	}

	c.incrRecent(n)

	return n.val
}

func (c *Cache) Put(key, val int) {
	n := c.nodeMap[key]
	if n != nil {
		n.val = val
		c.incrRecent(n)
	}

	if c.size == c.capacity {
		toRemove := c.rlist.tail.pre
		c.rlist.removeNode(toRemove)
		delete(c.nodeMap, toRemove.key)
		c.size--
	}

	n = &node{
		key: key,
		val: val,
	}

	c.nodeMap[key] = n
	c.size++

	c.rlist.addNode(n)
}
