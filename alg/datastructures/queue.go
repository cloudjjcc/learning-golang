package datastructures

import "container/list"

type IQueue interface {
	Len() int
	Front() interface{}
	Enqueue(data interface{})
	Dequeue() interface{}
	Empty() bool
}
type Queue struct {
	list list.List
}

func (q *Queue) Len() int {
	return q.list.Len()
}
func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}
func (q *Queue) Enqueue(data interface{}) {
	q.list.PushBack(data)
}

func (q *Queue) Dequeue() interface{} {
	if q.Len() > 0 {
		return q.list.Remove(q.list.Front())
	}
	return nil
}
func (q *Queue) Front() interface{} {
	return q.list.Front()
}
