package nowcoder

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

type node struct {
	data interface{}
	next *node
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
