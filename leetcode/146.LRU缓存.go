package leetcode

import (
	"fmt"
	"strings"
)

type LRUCache struct {
	cache map[int]*meta
	cap   int
	lru   *metaList
}
type metaList struct {
	head, tail *meta
}
type meta struct {
	pre, next *meta
	key       int
	val       int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*meta, capacity),
		cap:   capacity,
		lru:   &metaList{},
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.lru.moveToTail(v)
	return v.val
}
func (list *metaList) remove(v *meta) {
	if v.pre != nil {
		v.pre.next = v.next
	} else {
		list.head = v.next
	}
	v.next.pre = v.pre
	v.pre = nil
	v.next = nil
}

// 移动到队尾
func (list *metaList) moveToTail(v *meta) {
	if list.tail == v {
		return
	}
	list.remove(v)
	list.addToTail(v)
}

// 添加到队尾
func (list *metaList) addToTail(ele *meta) {
	if list.head == nil {
		list.head = ele
		list.tail = ele
		return
	}
	list.tail.next = ele
	ele.pre = list.tail
	list.tail = ele
	ele.next = nil
}
func (list *metaList) removeHead() {
	if list.head == nil {
		return
	}
	list.head = list.head.next
	if list.head != nil {
		list.head.pre = nil
	}
}
func (c *LRUCache) Print() string {
	sb := strings.Builder{}
	for cur := c.lru.head; cur != nil; cur = cur.next {
		sb.WriteString(fmt.Sprintf("%d->", cur.key))
	}
	return sb.String()
}
func (c *LRUCache) Put(key int, value int) {
	// 更新
	v, ok := c.cache[key]
	if ok {
		v.val = value
		c.lru.moveToTail(v)
		return
	}
	if len(c.cache) == c.cap {
		delete(c.cache, c.lru.head.key)
		c.lru.removeHead()
	}
	//新增
	ele := &meta{
		key: key,
		val: value,
	}
	c.lru.addToTail(ele)
	c.cache[key] = ele
}
