package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// queue
	queue := make([]*TreeNode, 0)
	queuePush := func(val *TreeNode) {
		queue = append(queue, val)
	}
	queuePop := func() *TreeNode {
		if len(queue) == 0 {
			return nil
		}
		top := queue[0]
		queue = queue[1:]
		return top
	}
	queueSize := func() int {
		return len(queue)
	}
	queueIsEmpty := func() bool {
		return len(queue) == 0
	}
	cur := root
	queuePush(cur)
	ans := make([][]int, 0)
	for !queueIsEmpty() {
		size := queueSize()
		var tmp []int
		for i := 0; i < size; i++ {
			cur = queuePop()
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queuePush(cur.Left)
			}
			if cur.Right != nil {
				queuePush(cur.Right)
			}
		}
		ans = append(ans, tmp)
	}
	// reverse
	left, right := 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return ans
}
