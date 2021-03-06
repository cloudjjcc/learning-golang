package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，
//重复的结点不保留，返回链表头指针。
//例如，链表1->2->3->3->4->4->5 处理后为 1->2->5

func deleteDuplicate(list *datastructures.ListNode) *datastructures.ListNode {
	if list == nil {
		return nil
	}
	if list.Next == nil {
		return list
	}
	var (
		nHead    = new(datastructures.ListNode)
		pre, cur = nHead, list
	)
	for cur != nil {
		// find repeat
		if cur.Next != nil && cur.Value == cur.Next.Value {
			for cur.Next != nil && cur.Value == cur.Next.Value {
				cur = cur.Next
			}
			pre.Next = cur.Next
		} else { // no repeat
			pre.Next = cur
			pre = pre.Next
		}
		cur = cur.Next
	}
	return nHead.Next
}
