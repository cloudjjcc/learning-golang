package leetcode

import "math"

func pairSum(head *ListNode) int {
	slow, fast := head, head
	// 找到后半部分起点n/2
	for fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半部分链表
	pre, cur, next := (*ListNode)(nil), slow, slow.Next
	for ; cur != nil; cur = next {
		next = cur.Next
		cur.Next = pre
		pre = cur
	}
	fast, slow = pre, head
	max := math.MinInt64
	for fast != nil && slow != nil {
		if tmp := slow.Val + fast.Val; tmp > max {
			max = tmp
		}
		slow = slow.Next
		fast = fast.Next
	}
	return max
}
