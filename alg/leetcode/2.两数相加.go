package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	overflow := 0
	l3 := new(ListNode)
	l3Head := l3
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + overflow
		overflow = val / 10
		val = val % 10
		l1 = l1.Next
		l2 = l2.Next
		l3.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		l3 = l3.Next
	}
	if l1 != nil {
		for l1 != nil {
			val := l1.Val + overflow
			overflow = val / 10
			val = val % 10
			l3.Next = &ListNode{Val: val}
			l1 = l1.Next
			l3 = l3.Next
		}
	}
	if l2 != nil {
		for l2 != nil {
			val := l2.Val + overflow
			overflow = val / 10
			val = val % 10
			l3.Next = &ListNode{Val: val}
			l2 = l2.Next
			l3 = l3.Next
		}
	}
	if overflow == 1 {
		l3.Next = &ListNode{Val: 1}
	}
	return l3Head.Next
}
