package leetcode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 找到链表中点
	findMiddle := func(h *ListNode) *ListNode {
		slow, fast := h, h
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}
	// 反转链表
	reverse := func(h *ListNode) *ListNode {
		var pre *ListNode
		for h != nil {
			next := h.Next
			h.Next = pre
			pre = h
			h = next
		}
		return pre
	}
	// 合并列表
	merge := func(a, b *ListNode) {
		var aNext, bNext *ListNode
		for a != nil && b != nil {
			aNext = a.Next
			bNext = b.Next
			a.Next = b
			b.Next = aNext

			a = aNext
			b = bNext
		}
	}
	middle := findMiddle(head)
	right := middle.Next
	middle.Next = nil
	right = reverse(right)
	merge(head, right)
}
