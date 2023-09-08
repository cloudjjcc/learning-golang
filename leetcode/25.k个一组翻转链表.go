package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	i := 0
	cur := dummy
	startPre := dummy //翻转起始点的前一个节点  startPre->start
	for ; cur != nil; i++ {
		if i > 0 && i%k == 0 {
			endNext := cur.Next //翻转结束点的后一个节点 end->endNext
			start := startPre.Next
			startPre.Next = nil
			cur.Next = nil                     // 形成 start->..->end 链表
			startPre.Next = reverseList(start) //翻转  start->...->end  ,startPre->end->...start->endNext
			startPre = start
			start.Next = endNext
			cur = endNext
			continue
		}
		cur = cur.Next
	}
	return dummy.Next
}
