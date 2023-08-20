package leetcode

// 思路：记录上一个值出现的次数
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	nHead := dummy
	pre := head
	duplicateCnt := 1
	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val == pre.Val {
			pre = cur
			duplicateCnt++
			continue
		}
		if duplicateCnt > 1 {
			pre = cur
			duplicateCnt = 1
			continue
		}
		nHead.Next = pre
		nHead = nHead.Next
		pre = cur
	}
	if duplicateCnt == 1 {
		nHead.Next = pre
	} else {
		nHead.Next = nil
	}
	return dummy.Next
}
