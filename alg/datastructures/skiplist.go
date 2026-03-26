package datastructures

import "math/rand"

const (
	maxLevel = 32
	p        = 0.25
)

type skipListNode struct {
	prev *skipListNode
	next []*skipListNode
	key  int
	val  interface{}
}

type SkipList struct {
	head, tail *skipListNode
	length     int
	level      int
}

func NewSkipList() *SkipList {
	list := &SkipList{}
	list.level = 1
	list.head = list.createNode(maxLevel, 0, nil)
	return list
}

// 获得一个随机的层数
func (sl *SkipList) randLevel() int {
	level := 1
	for level < 32 && float64(rand.Int()&0xFFFF) < 0xFFFF*0.25 {
		level++
	}
	return level
}
func (sl *SkipList) createNode(level int, key int, val interface{}) *skipListNode {
	return &skipListNode{
		key:  key,
		val:  val,
		next: make([]*skipListNode, level),
	}
}
func (sl *SkipList) Put(key int, val interface{}) {
	update := make([]*skipListNode, maxLevel)
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key <= key {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	// key 已存在
	if update[0].key == key {
		update[0].val = val
		return
	}
	// 获取新节点的最大层数
	level := sl.randLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}
	// 插入新节点
	node := sl.createNode(level, key, val)
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
	// 更新prev指针
	node.prev = update[0]
	if node.next[0] != nil {
		node.next[0].prev = node
	} else {
		sl.tail = node //更新tail节点
	}
	sl.length++
}
func (sl *SkipList) Get(key int) interface{} {
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key <= key {
			cur = cur.next[i]
		}
		if cur.key == key {
			return cur.val
		}
	}
	return nil
}
func (sl *SkipList) Del(key int) bool {
	cur := sl.head
	update := make([]*skipListNode, sl.level)
	for i := sl.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key < key {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	// 待删除的节点
	x := cur.next[0]
	if x == nil || x.key != key {
		return false
	}
	for i := 0; i < sl.level; i++ {
		if update[i].next[i] != x {
			break
		}
		update[i].next[i] = x.next[i]
	}
	if x.next[0] != nil {
		x.next[0].prev = x.prev
	} else {
		sl.tail = x.prev
	}
	for sl.level > 1 && sl.head.next[sl.level] == nil {
		sl.level--
	}
	sl.length--
	return true
}
