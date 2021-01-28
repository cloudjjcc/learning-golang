package main

import "fmt"

type linkedNode struct {
	data interface{}
	next *linkedNode
}

func main() {
	list := &linkedNode{
		data: 0,
		next: nil,
	}
	// 构建链表
	cur := list
	for i := 1; i < 10; i++ {
		cur.next = &linkedNode{
			data: i,
			next: nil,
		}
		cur = cur.next
	}
	fmt.Println(findKthToTail(list, 3).data)
}

// 查找链表中的倒数第k个节点
func findKthToTail(list *linkedNode, k int) *linkedNode {
	if list == nil || k <= 0 {
		return nil
	}
	var (
		p1, p2 *linkedNode
		i      int
	)
	for p1, p2, i = list, list, 0; p1 != nil; p1, i = p1.next, i+1 {
		if i >= k {
			p2 = p2.next
		}
	}
	return p2
}
