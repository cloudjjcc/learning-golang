package main

import "fmt"

//题目描述
//输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
//假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

type TreeNode struct {
	data       interface{}
	leftChild  *TreeNode
	rightChild *TreeNode
}

func rebuildBinaryTree(pre []int, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	root := &TreeNode{data: pre[0]}
	for i, v := range in {
		if v == root.data {
			root.leftChild = rebuildBinaryTree(pre[1:i+1], in[0:i])
			root.rightChild = rebuildBinaryTree(pre[i+1:], in[i+1:])
			break
		}
	}
	return root
}
func main() {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 8, 6}
	tree := rebuildBinaryTree(pre, in)
	fmt.Println(tree)
}
