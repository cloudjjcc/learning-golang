package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//请实现一个函数，用来判断一颗二叉树是不是对称的。
//注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的。

func main() {
	tree := datastructures.BuildTreeFromArray([]interface{}{1, 1, 1})
	fmt.Println(isSymmetrical(tree))
}

func isSymmetrical(tree *datastructures.TestTreeNode) bool {
	if tree == nil {
		return true
	}
	return checkCore(tree.Left, tree.Right)
}
func checkCore(tree1, tree2 *datastructures.TestTreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Value != tree2.Value {
		return false
	}
	return checkCore(tree1.Left, tree2.Right) && checkCore(tree2.Left, tree1.Right)
}
