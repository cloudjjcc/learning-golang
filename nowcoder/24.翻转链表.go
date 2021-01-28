package main

import "fmt"

type lnode struct {
	data interface{}
	next *lnode
}

func main() {
	list := &lnode{
		data: 0,
		next: nil,
	}
	// 构建链表
	cur := list
	for i := 1; i < 10; i++ {
		cur.next = &lnode{
			data: i,
			next: nil,
		}
		cur = cur.next
	}
	newlist := ReverseList(list)
	fmt.Println(newlist)
}

func ReverseList(list *lnode) *lnode {
	if list == nil {
		return nil
	}
	var (
		prev *lnode
		cur  = list
		next *lnode
	)
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}
