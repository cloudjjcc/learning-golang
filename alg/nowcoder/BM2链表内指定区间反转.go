package nowcoder

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	i := 0
	var (
		pre, cur, next     *ListNode = nil, head, head.Next
		preM, mNode, nNext *ListNode = nil, nil, nil
	)
	for {
		if i >= m {
			if i == m {
				preM = pre
				pre = nil
				mNode = cur
			}
			cur.Next = pre
			pre = cur
			cur = next
			next = next.Next
			if i == n {
				nNext = cur
				break
			}
			i++
			continue
		}
		pre = cur
		cur = next
		next = next.Next
		i++
	}
	preM.Next = pre
	mNode.Next = nNext
	return head
}

// todo
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	pre, cur, next := dummy, head, head.Next
	for i := 0; i < k; i++ {
		pre = cur
		cur.Next = pre
		cur = next
		if cur == nil {
			break
		}
		next = next.Next
	}
	dummy.Next.Next = reverseKGroup(next, k)
	return pre
}
