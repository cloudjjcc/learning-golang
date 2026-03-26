package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	reverseTree(root.Right)
	return isSameTree(root.Left, root.Right)
}

func reverseTree(root *TreeNode) {
	if root == nil {
		return
	}
	reverseTree(root.Left)
	reverseTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
}
