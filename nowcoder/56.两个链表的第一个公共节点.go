package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//输入两个链表，找出它们的第一个公共结点。
//（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

func firstCommonNode(list1 *datastructures.ListNode, list2 *datastructures.ListNode) *datastructures.ListNode {
	if list1 == nil || list2 == nil {
		return nil
	}
	p1, p2 := list1, list2
	for p1 != p2 {
		if p1 == nil {
			p1 = list2
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = list1
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
