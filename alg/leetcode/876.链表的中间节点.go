package leetcode

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。

func middleNode(list *datastructures.ListNode) *datastructures.ListNode {
	if list == nil {
		return nil
	}
	p1, p2 := list, list
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}
