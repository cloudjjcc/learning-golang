package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。

func main() {
	testNode := &datastructures.ListNode{Value: 0}
	a := &datastructures.ListNode{Value: 1}
	b := &datastructures.ListNode{Value: 2}
	c := &datastructures.ListNode{Value: 3}
	d := &datastructures.ListNode{Value: 4}
	testNode.Next = a
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = a
	fmt.Println(entryNodeOfLoop(testNode).Value)
	fmt.Println(entryNodeOfLoop2(testNode).Value)
}

// 辅助空间
func entryNodeOfLoop(node *datastructures.ListNode) *datastructures.ListNode {
	if node == nil {
		return nil
	}
	tmp := make(map[*datastructures.ListNode]bool)
	for node != nil {
		if _, ok := tmp[node]; ok {
			return node
		}
		tmp[node] = true
		node = node.Next
	}
	return nil
}

// 快慢指针
func entryNodeOfLoop2(node *datastructures.ListNode) *datastructures.ListNode {
	if node == nil {
		return nil
	}
	// 找到相遇点
	findMeetingNode := func(node *datastructures.ListNode) *datastructures.ListNode {
		var (
			p1, p2 = node, node
		)
		for p2 != nil && p2.Next != nil {
			p1 = p1.Next
			p2 = p2.Next.Next
			if p1 == p2 {
				return p1
			}
		}
		return nil
	}
	meetingNode := findMeetingNode(node)
	if meetingNode == nil {
		return nil
	}
	p1 := node
	p2 := meetingNode
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
