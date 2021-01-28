package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//从上往下打印出二叉树的每个节点，同层节点从左至右打印。
//
//例如，以下二叉树层次遍历的结果为：1,2,3,4,5,6,7

func main() {
	tree := datastructures.BuildTreeFromArray([]interface{}{1, 2, 3, 4, 5, 6, 7})
	printBinaryTree(tree)
}

// level order
func printBinaryTree(tree *datastructures.TestTreeNode) {
	if tree == nil {
		return
	}
	queue := &datastructures.Queue{}
	queue.Enqueue(tree)
	for !queue.Empty() {
		root := queue.Dequeue().(*datastructures.TestTreeNode)
		fmt.Println(root.Value)
		if root.Left != nil {
			queue.Enqueue(root.Left)
		}
		if root.Right != nil {
			queue.Enqueue(root.Right)
		}
	}
}
