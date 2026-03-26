package leetcode

func averageOfLevels(root *TreeNode) []float64 {
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
	ans := make([]float64, 0)
	for !empty() {
		s := size()
		avg := float64(0)
		for i := 0; i < s; i++ {
			n := pop()
			avg += float64(n.Val)
			if n.Left != nil {
				push(n.Left)
			}
			if n.Right != nil {
				push(n.Right)
			}
		}
		avg /= float64(s)
		ans = append(ans, avg)
	}
	return ans
}
