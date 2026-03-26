package skiplist

import "math/rand/v2"

const (
	MaxLevel = 32
)

type Node struct {
	Key  int
	Val  int
	Next []*Node
	Prev *Node
}

type SkipList struct {
	randSource *rand.Rand
	Level      int   //当前最大层数
	Length     int   //当前元素个数
	Tail       *Node //最后一个元素节点
	Head       *Node
}

func NewSkipList() *SkipList {
	head := &Node{
		Key:  0,
		Val:  0,
		Next: make([]*Node, MaxLevel),
		Prev: nil,
	}
	return &SkipList{
		Level:      1,
		Length:     0,
		Head:       head,
		randSource: rand.New(rand.NewPCG(1, 2)),
	}
}

// 获取一个随机层数 每层概率：25%
func (l *SkipList) randLevel() int {
	level := 1
	for level < MaxLevel && l.randSource.Uint32N(100) < 25 {
		level++
	}
	return level
}
func (l *SkipList) insert(key int, val int) {
	update := make([]*Node, MaxLevel)
	cur := l.Head
	for i := l.Level - 1; i >= 0; i-- {
		for cur.Next[i] != nil &&
			(cur.Next[i].Key < key ||
				(cur.Next[i].Key == key && cur.Next[i].Val < val)) {
			cur = cur.Next[i]
		}
		update[i] = cur
	}
	level := l.randLevel()
	if level > l.Level {
		for i := l.Level; i < level; i++ {
			update[i] = l.Head
		}
		l.Level = level
	}

	x := &Node{
		Key:  key,
		Val:  val,
		Next: make([]*Node, level),
	}
	for i := 0; i < level; i++ {
		x.Next[i] = update[i].Next[i]
		update[i].Next[i] = x
	}
	if update[0] == l.Head {
		x.Prev = nil
	} else {
		x.Prev = update[0]
	}
	if x.Next[0] != nil {
		x.Next[0].Prev = x
	} else {
		l.Tail = x
	}
	l.Length++
}
func (l *SkipList) Put(key int, val int) {
	l.insert(key, val)
}

func (l *SkipList) Get(key int) *Node {
	cur := l.Head
	for i := l.Level - 1; i >= 0; i-- {
		for cur.Next[i] != nil && cur.Next[i].Key < key {
			cur = cur.Next[i]
		}
	}
	cur = cur.Next[0]
	if cur != nil && cur.Key == key {
		return cur
	}
	return nil
}
