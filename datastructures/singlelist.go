package datastructures

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return "{}"
	}
	cur := l
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	for cur != nil {
		buf.WriteString(fmt.Sprintf("->%v", cur.Value))
		cur = cur.Next
	}
	return buf.String()
}

func BuildList(arr []int) *ListNode {
	list := new(ListNode)
	if len(arr) == 0 {
		return list
	}
	list.Value = arr[0]
	if len(arr) == 1 {
		return list
	}
	cur := list
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Value: arr[i]}
		cur = cur.Next
	}
	return list
}

type SingleList struct {
	head ListNode
	tail ListNode
	len  int
}

func NewSingleList() *SingleList {
	return new(SingleList).Init()
}
func (l *SingleList) Front() *ListNode {
	if l.len == 0 {
		return nil
	}
	return l.head.Next
}
func (l *SingleList) Back() *ListNode {
	if l.len == 0 {
		return nil
	}
	return &l.tail
}
func (l *SingleList) PushBack(data interface{}) {
	l.insertValue(data, &l.tail)
}
func (l *SingleList) PushFront(data interface{}) {
	l.insertValue(data, l.head.Next)
}
func (l *SingleList) PopBack() error {
	return nil
}

func (l *SingleList) Len() int {
	return l.len
}

func (l *SingleList) Init() *SingleList {
	l.head.Next = &l.tail
	return l
}

// insert value after at
func (l *SingleList) insertValue(v interface{}, at *ListNode) {
	node := &ListNode{Value: v}
	at.Next = node
	l.len++
}
