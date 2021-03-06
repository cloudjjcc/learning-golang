package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//输入一棵二叉树，求该树的深度。
//从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

// h(root)=max(h(root.left),h(root.right))+1
func getTreeDepth(tree *datastructures.TestTreeNode) int {
	if tree == nil {
		return 0
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	left := getTreeDepth(tree.Left)
	right := getTreeDepth(tree.Right)
	return max(left, right) + 1
}
