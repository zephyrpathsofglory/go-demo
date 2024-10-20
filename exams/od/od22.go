/*
浏览器历史记录
*/
package od

type BrowserHistorySys struct {
	head *node
	tail *node
	cur  *node

	curIdx   int // 从 0 开始， cur 指向结点的 idx, cur 指向结点变化，更新
	length   int
	capacity int // > 0, 超出时清除最早的记录
}

type node struct {
	prev *node
	next *node

	url string
}

func NewBrowserHistorySys(homepage string, maxCount int) *BrowserHistorySys {
	n := &node{
		url: homepage,
	}

	return &BrowserHistorySys{
		cur:      n,
		head:     n,
		tail:     n,
		length:   1,
		capacity: maxCount,
	}
}

// 返回浏览历史中的缓存页面数量
// 如果入参 url 不是当前页，则跳转到此 url，并把此 url 作为当前页，同时清除浏览历史中原当前页的前进记录，再将此 url 缓存到浏览历史中。
// 如果入参 url 是当前页，则不变
func (s *BrowserHistorySys) visit(url string) int {
	// 跳转页面是当前页
	if s.cur.url == url {
		return s.length
	}

	// 跳转页面不是当前页
	nxt := s.cur.next

	n := &node{
		url:  url,
		prev: s.cur,
	}

	s.cur.next = n

	if nxt != nil {
		nxt.prev = nil
	}

	s.cur = n
	s.curIdx += 1
	s.length = s.curIdx + 1
	s.tail = n

	// 历史记录超过容量
	if s.curIdx == s.capacity {
		preHead := s.head
		s.head = s.head.next
		preHead.next = nil
		s.head.prev = nil
		s.curIdx -= 1
		s.length -= 1
	}

	return s.length
}

// 在浏览历史中从当前页后退一部，返回停留页面的url，并作为当前页, 如果已经退无可退，则不再后退，继续停留在当前页
func (s *BrowserHistorySys) back() string {
	if s.cur.prev != nil {
		s.cur = s.cur.prev
		s.curIdx -= 1
	}

	return s.cur.url
}

// 在浏览历史中从当前页面前进一部，返回停留页面的 url，并作为当前页，若进无可进，则不再前进， 继续停留在当前页。
func (s *BrowserHistorySys) forward() string {
	if s.cur.next != nil {
		s.cur = s.cur.next
		s.curIdx += 1
	}

	return s.cur.url
}
