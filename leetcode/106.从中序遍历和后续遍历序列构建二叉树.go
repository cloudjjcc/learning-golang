package leetcode

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return root
	}
	rootIndex := 0
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == root.Val {
			break
		}
	}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-2])
	return root
}
