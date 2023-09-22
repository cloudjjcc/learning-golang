package leetcode

type DoublyLinkedListNode struct {
	pre, next *DoublyLinkedListNode
	key       int
	val       int
}

type LRUCache1 struct {
	size       int
	cap        int
	m          map[int]*DoublyLinkedListNode
	head, tail *DoublyLinkedListNode
}

func NewLRU1(capacity int) *LRUCache1 {
	return &LRUCache1{
		cap: capacity,
		m:   make(map[int]*DoublyLinkedListNode),
	}
}

func (c *LRUCache1) Get(key int) int {
	node := c.m[key]
	if node == nil {
		return -1
	}
	c.moveToHead(node)
	return node.val
}
func (c *LRUCache1) moveToHead(n *DoublyLinkedListNode) {
	if n == c.head {
		return
	}
	// 更新c.tail
	if n == c.tail {
		c.tail = n.pre
	}
	// 更新n.pre
	n.pre.next = n.next
	// 更新n.next
	if n.next != nil {
		n.next.pre = n.pre
	}
	// 更新n
	n.next = c.head
	n.pre = nil
	// 更新c.head
	c.head.pre = n
	c.head = n
	return
}
func (c *LRUCache1) deleteTail() {
	if c.tail == nil {
		return
	}
	delete(c.m, c.tail.key)
	if c.tail.pre != nil {
		c.tail.pre.next = nil
	}
	c.tail = c.tail.pre
	if c.tail == nil {
		c.head = nil
	}
	c.size--
}
func (c *LRUCache1) Put(key int, value int) {
	// key 存在
	node := c.m[key]
	if node != nil {
		node.val = value
		c.moveToHead(node)
		return
	}
	// key 不存在
	if c.size+1 > c.cap {
		c.deleteTail()
	}
	node = &DoublyLinkedListNode{key: key, val: value}
	node.next = c.head
	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		c.head.pre = node
		c.head = node
	}
	c.m[key] = node
	c.size++
}
