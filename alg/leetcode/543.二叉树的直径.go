package leetcode

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(r *TreeNode) int
	ans := math.MinInt64
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(r *TreeNode) int { //返回经过r的最长路径
		if r == nil || r.Left == nil && r.Right == nil {
			return 0
		}
		left := dfs(r.Left)
		right := dfs(r.Right)
		p := 0
		if r.Left != nil {
			p += left
			p += 1
		}
		if r.Right != nil {
			p += right
			p += 1
		}
		ans = max(ans, p)
		return 1 + max(left, right)
	}
	ans = max(ans, dfs(root))
	return ans
}
