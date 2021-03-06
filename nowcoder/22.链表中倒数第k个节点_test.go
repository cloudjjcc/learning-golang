package nowcoder

import (
	"fmt"
	"testing"
)

func Test_findKthToTail(t *testing.T) {
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
