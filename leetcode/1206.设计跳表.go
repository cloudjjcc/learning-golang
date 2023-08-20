package leetcode

import (
	"math"
	"math/rand"
	"time"
)

const (
	MaxLevel = 32
	P        = 0.25
)

type skiplistNode struct {
	key  int
	next []*skiplistNode
}

type Skiplist struct {
	head  *skiplistNode
	len   int
	level int
}

func NewSkiplist() *Skiplist {
	rand.Seed(time.Now().Unix())
	sl := &Skiplist{}
	sl.head = &skiplistNode{key: math.MinInt64}
	sl.head.next = make([]*skiplistNode, MaxLevel)
	return sl
}

func (s *Skiplist) Search(target int) bool {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for ; cur.next[i] != nil && cur.next[i].key < target; cur = cur.next[i] {
		}
	}
	cur = cur.next[0]
	return cur != nil && cur.key == target
}
func (s *Skiplist) Add(num int) {
	randLevel := func() int {
		level := 1
		for rand.Int63n(100) < 100*P && level < MaxLevel {
			level++
		}
		return level
	}
	level := randLevel()
	if level > s.level {
		s.level = level
	}
	node := &skiplistNode{key: num, next: make([]*skiplistNode, level)}
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for ; cur.next[i] != nil && cur.next[i].key < node.key; cur = cur.next[i] {
		}
		if i < level {
			node.next[i] = cur.next[i]
			cur.next[i] = node
		}
	}
	s.len++
}

func (s *Skiplist) Erase(target int) bool {
	cur := s.head
	update := make([]*skiplistNode, s.level)
	for i := s.level - 1; i >= 0; i-- {
		for ; cur.next[i] != nil && cur.next[i].key < target; cur = cur.next[i] {

		}
		update[i] = cur
	}
	cur = cur.next[0]
	if cur == nil || cur.key != target {
		return false
	}
	for i, v := range update {
		if v.next[i] == cur {
			v.next[i] = cur.next[i]
		}
	}
	for i := s.level - 1; i >= 0; i-- {
		if s.head.next[i] == nil {
			s.level--
		} else {
			break
		}
	}
	return true
}
