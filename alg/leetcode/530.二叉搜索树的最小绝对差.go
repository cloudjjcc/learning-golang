package leetcode

import "math"

func getMinimumDifference(root *TreeNode) int {
	var inOrder func(t *TreeNode)
	var pre *TreeNode
	min := math.MaxInt64
	inOrder = func(t *TreeNode) {
		if t == nil {
			return
		}
		inOrder(t.Left)
		if pre != nil {
			if v := t.Val - pre.Val; v < min {
				min = v
			}
		}
		pre = t
		inOrder(t.Right)
	}
	inOrder(root)
	return min
}
