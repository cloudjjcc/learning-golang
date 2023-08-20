package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	leftPre := dummy
	idx := 1
	pre, cur, next := dummy, head, head.Next
	for cur != nil && idx <= right {
		if idx == left-1 {
			leftPre = cur
		}
		if idx >= left {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		} else {
			cur = cur.Next
		}
		idx++
	}
	leftPre.Next.Next = cur
	leftPre.Next = pre
	return dummy.Next
}
