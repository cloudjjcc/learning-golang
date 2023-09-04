package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	s := make([]*TreeNode, 0)
	push := func(n *TreeNode) {
		s = append(s, n)
	}
	pop := func() *TreeNode {
		tmp := s[len(s)-1]
		s = s[:len(s)-1]
		return tmp
	}
	isEmpty := func() bool {
		return len(s) == 0
	}
	preVal := math.MinInt64
	for !isEmpty() || root != nil {
		for root != nil {
			push(root)
			root = root.Left
		}
		root = pop()
		if root.Val <= preVal {
			return false
		}
		preVal = root.Val
		root = root.Right
	}
	return true
}
