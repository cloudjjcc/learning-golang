package leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := maxDepth(root.Left)
	if t := maxDepth(root.Right); t > max {
		max = t
	}
	return max + 1
}
