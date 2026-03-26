package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	push := func(n *TreeNode) {
		queue = append(queue, n)
	}
	pop := func() *TreeNode {
		top := queue[0]
		queue = queue[1:]
		return top
	}
	empty := func() bool {
		return len(queue) == 0
	}
	size := func() int {
		return len(queue)
	}
	push(root)
	ans := make([][]int, 0)
	isNormal := true
	for !empty() {
		s := size()
		level := make([]int, s)
		for i := 0; i < s; i++ {
			n := pop()
			if isNormal {
				level[i] = n.Val
			} else {
				level[s-i-1] = n.Val
			}
			if n.Left != nil {
				push(n.Left)
			}
			if n.Right != nil {
				push(n.Right)
			}
		}
		isNormal = !isNormal
		ans = append(ans, level)
	}
	return ans
}
