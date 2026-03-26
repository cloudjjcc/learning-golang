package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Next.Val < head.Val {
			next := head.Next
			head.Next = nil
			next.Next = head
			return next
		}
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)
	ansHead := new(ListNode)
	cur := ansHead
	for {
		if left == nil {
			cur.Next = right
			break
		}
		if right == nil {
			cur.Next = left
			break
		}
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	return ansHead.Next
}
