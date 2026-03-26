package nowcoder

import "testing"

func Test_reversePrintList(t *testing.T) {
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
	reversePrintList(list)
	reversePrintList2(list)
}
