package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
