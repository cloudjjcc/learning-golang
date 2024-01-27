package leetcode

import (
	"fmt"
	"strings"
)

type MyLinkedList struct {
	size       int
	head, tail *ListNode
}

func NewMyLinkedList() *MyLinkedList {
	return &MyLinkedList{}
}
func (l *MyLinkedList) String() string {
	cur := l.head
	sb := strings.Builder{}
	sb.WriteString("[")
	for cur != nil {
		sb.WriteString(fmt.Sprintf("%d->", cur.Val))
		cur = cur.Next
	}
	sb.WriteString("]")
	return sb.String()
}
func (l *MyLinkedList) Get(index int) int {
	node := l.getNode(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.size++
	node := &ListNode{Val: val}
	node.Next = l.head
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.size++
	node := &ListNode{Val: val}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
}
func (l *MyLinkedList) getNode(index int) *ListNode {
	if index < 0 || index >= l.size {
		return nil
	}
	cur := l.head
	idx := 0
	for cur != nil && idx < index {
		cur = cur.Next
		idx++
	}
	return cur
}
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.size {
		return
	}
	l.size++
	if index == 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.size {
		l.AddAtTail(val)
	}
	preNode := l.getNode(index - 1)
	next := preNode.Next
	node := &ListNode{Val: val}
	preNode.Next = node
	node.Next = next
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	l.size--
	if index == 0 {
		next := l.head.Next
		l.head.Next = nil
		l.head = next
		if l.head == nil {
			l.tail = nil
		}
		return
	}
	preNode := l.getNode(index - 1)
	next := preNode.Next
	if next != nil {
		next = next.Next
	}
	preNode.Next = next
}
