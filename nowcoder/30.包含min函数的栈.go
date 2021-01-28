package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//定义栈的数据结构，请在该类型中实现一个能够得到栈最小元素的 min 函数。

func main() {
	s := NewMyStack()
	s.Push(10)
	s.Push(2)
	s.Push(3)
	s.Push(3)
	s.Push(1)
	s.Push(9)
	fmt.Println(s.Min())
	s.Pop()
	s.Pop()
	fmt.Println(s.Min())
}

// 利用辅助栈
type MyStack struct {
	s *datastructures.Stack
	m *datastructures.Stack
}

func NewMyStack() *MyStack {
	return &MyStack{
		s: new(datastructures.Stack),
		m: new(datastructures.Stack),
	}
}
func (s *MyStack) Min() int {
	return s.m.Peek().(int)
}
func (s *MyStack) Push(data int) {
	s.s.Push(data)
	if s.m.Empty() {
		s.m.Push(data)
	} else if data < s.m.Peek().(int) {
		s.m.Push(data)
	}
}
func (s *MyStack) Pop() int {
	if s.s.Len() <= 0 {
		return 0
	}
	i := s.s.Pop().(int)
	if i == s.m.Peek() {
		s.m.Pop()
	}
	return i
}
