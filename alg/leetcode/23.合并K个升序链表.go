package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	h := make([]*ListNode, 0, len(lists))
	dummyHead := new(ListNode)
	cur := dummyHead
	k := len(lists)
	heapDown := func(i int) {
		for {
			leftCh := 2*i + 1
			if leftCh == len(h) {
				break
			}
			rightCh := leftCh + 1
			maxCh := leftCh
			if rightCh < len(h) && h[rightCh].Val < h[leftCh].Val {
				maxCh = rightCh
			}
			if h[i].Val <= h[maxCh].Val {
				break
			}
			h[i], h[maxCh] = h[maxCh], h[i]
			i = maxCh
		}
	}
	heapUp := func(i int) {
		for {
			p := (i - 1) / 2
			if p < 0 {
				break
			}
			if h[p].Val <= h[i].Val {
				break
			}
			h[p], h[i] = h[i], h[p]
			i = p
		}
	}
	push := func(n *ListNode) {
		h = append(h, n)
		heapUp(len(h) - 1)
	}
	pop := func() *ListNode {
		if len(h) == 0 {
			return nil
		}
		head := h[0]
		h[0] = h[len(h)-1]
		h = h[:len(h)-1]
		heapDown(0)
		return head
	}
	// build min heap
	for i := 0; i < k; i++ {
		if lists[i] != nil {
			push(lists[i])
		}
	}
	for {
		node := pop()
		if node == nil {
			break
		}
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			push(node.Next)
		}
	}
	return dummyHead.Next
}
