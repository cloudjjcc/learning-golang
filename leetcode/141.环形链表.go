package leetcode

func hasCycle(head *ListNode) bool {
	dummy := new(ListNode)
	dummy.Next = head
	fast, slow := dummy, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
