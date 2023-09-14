package leetcode

func kthSmallest(root *TreeNode, k int) int {
	var inOrder func(t *TreeNode)
	idx := 0
	ans := 0
	inOrder = func(t *TreeNode) {
		if t == nil {
			return
		}
		inOrder(t.Left)
		idx++
		if idx == k {
			ans = t.Val
			return
		}
		inOrder(t.Right)
	}
	inOrder(root)
	return ans
}
