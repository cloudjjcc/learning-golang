package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//用两个栈来实现一个队列，完成队列的 Push 和 Pop 操作。

func main() {
	push(1)
	push(2)
	push(3)
	pop()
	pop()
	push(4)
	fmt.Println(pop())
}

var (
	stack1 = &datastructures.Stack{}
	stack2 = &datastructures.Stack{}
)

func push(node int) {
	stack1.Push(node)
}
func pop() int {
	if stack2.Empty() {
		for !stack1.Empty() {
			stack2.Push(stack1.Pop())
		}
	}
	return stack2.Pop().(int)
}
