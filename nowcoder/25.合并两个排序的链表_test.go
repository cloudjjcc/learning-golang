package nowcoder

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	list1 := new(anode)
	list2 := new(anode)
	list2.data = 10
	// 构建链表1
	cur := list1
	for i := 1; i < 10; i++ {
		cur.next = &anode{
			data: i,
			next: nil,
		}
		cur = cur.next
	}
	// 构建链表2
	cur = list2
	for i := 11; i < 15; i++ {
		cur.next = &anode{
			data: i,
			next: nil,
		}
		cur = cur.next
	}
	list3 := Merge(list1, list2)
	fmt.Println(list3)
}
