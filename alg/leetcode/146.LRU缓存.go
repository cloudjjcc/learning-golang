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

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.lru.moveToTail(v)
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
func (this *LRUCache) Print() string {
	sb := strings.Builder{}
	for cur := this.lru.head; cur != nil; cur = cur.next {
		sb.WriteString(fmt.Sprintf("%d->", cur.key))
	}
	return sb.String()
}
func (this *LRUCache) Put(key int, value int) {
	// 更新
	v, ok := this.cache[key]
	if ok {
		v.val = value
		this.lru.moveToTail(v)
		return
	}
	if len(this.cache) == this.cap {
		delete(this.cache, this.lru.head.key)
		this.lru.removeHead()
	}
	//新增
	ele := &meta{
		key: key,
		val: value,
	}
	this.lru.addToTail(ele)
	this.cache[key] = ele
}
