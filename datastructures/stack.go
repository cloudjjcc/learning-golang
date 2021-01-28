package datastructures

import "container/list"

// go 语言实现stack 可用container/list 或slice
type Stack struct {
	list list.List
}

func (s *Stack) Pop() interface{} {
	if s.list.Len() > 0 {
		return s.list.Remove(s.list.Back())
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Len() > 0 {
		return s.list.Back().Value
	}
	return nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Push(ele interface{}) {
	s.list.PushBack(ele)
}
func (s *Stack) Empty() bool {
	return !(s.list.Len() > 0)
}
