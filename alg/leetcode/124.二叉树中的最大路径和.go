package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(r *TreeNode) int
	maxSum := math.MinInt32
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := max(0, dfs(r.Left))
		right := max(0, dfs(r.Right))
		p := root.Val + left + right
		maxSum = max(maxSum, p)
		return root.Val + max(left, right)
	}
	dfs(root)
	return maxSum
}
