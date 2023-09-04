package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := new(ListNode)
	dummyHead.Next = head
	pre, cur, next := dummyHead, head, head.Next
	for pre != nil && cur != nil && next != nil {
		tmp := next.Next
		// 交换cur和next
		pre.Next = next
		next.Next = cur
		cur.Next = tmp
		// 更新 pre,cur,next
		pre = cur
		cur = tmp
		if tmp != nil {
			next = tmp.Next
		} else {
			break
		}
	}
	return dummyHead.Next
}
