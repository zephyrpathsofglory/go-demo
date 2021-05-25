package lfu

type node struct {
	pre  *node
	next *node
	key  int
	val  int
	freq int
}

func newNode() *node {
	return &node{
		freq: 1,
	}
}

type Cache struct {
	min      int
	capacity int
	size     int
	freqMap  map[int]*doublyLinkedList
	nodeMap  map[int]*node
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		freqMap:  make(map[int]*doublyLinkedList),
		nodeMap:  make(map[int]*node),
	}
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

	dll.head.next = node
	dll.head.next.pre = node
}

func (c *Cache) incrFreq(n *node) {
	freq := n.freq
	freqDll := c.freqMap[freq]

	freqDll.removeNode(n)

	if freq == c.min && freqDll.head.next == freqDll.tail.pre {
		c.min = freq + 1
	}

	n.freq = freq + 1
	nextFreqDll := c.freqMap[freq+1]
	if nextFreqDll == nil {
		nextFreqDll = newDLL()
	}

	c.freqMap[freq+1] = nextFreqDll
	nextFreqDll.addNode(n)
}

func (c *Cache) Get(key int) (val int) {
	node := c.nodeMap[key]
	if node == nil {
		return -1
	}

	c.incrFreq(node)

	val = node.val

	return
}

func (c *Cache) Put(key, val int) {
	if c.capacity == 0 {
		return
	}

	node := c.nodeMap[key]
	if node != nil {
		node.val = val
		c.incrFreq(node)
	}

	if c.size == c.capacity {
		dll := c.freqMap[c.min]
		dll.removeNode(dll.tail.pre)
		delete(c.nodeMap, dll.tail.pre.key)
		c.size = c.size - 1
	}
	node = newNode()
	node.key = key
	node.val = val
	c.nodeMap[node.key] = node
	c.min = 1
	c.size = c.size + 1
	freqDll := c.freqMap[1]
	if freqDll == nil {
		freqDll = newDLL()
	}

	freqDll.addNode(node)
}
