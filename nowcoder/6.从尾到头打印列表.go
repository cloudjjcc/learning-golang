package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

type node struct {
	data interface{}
	next *node
}

func main() {
	list := &node{
		data: 0,
		next: nil,
	}
	// 构建链表
	cur := list
	for i := 1; i < 10; i++ {
		cur.next = &node{
			data: i,
			next: nil,
		}
		cur = cur.next
	}
	//reversePrintList(list)
	reversePrintList2(list)
}

// 递归实现
func reversePrintList(list *node) {
	if list != nil {
		reversePrintList(list.next)
		fmt.Println(list.data)
	}
}

// stack 实现
func reversePrintList2(list *node) {
	var stack datastructures.Stack
	for cur := list; cur != nil; cur = cur.next {
		stack.Push(cur.data)
	}
	for top := stack.Pop(); top != nil; top = stack.Pop() {
		fmt.Println(top)
	}
}
