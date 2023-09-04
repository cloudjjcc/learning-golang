package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	if p.Left == nil && p.Right == nil && q.Left == nil && q.Right == nil {
		return true
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
