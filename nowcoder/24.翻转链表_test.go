package nowcoder

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
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
