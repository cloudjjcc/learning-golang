package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，
//第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。

func main() {
	tree := datastructures.BuildTreeFromArray([]interface{}{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(printTree(tree))
}

func printTree(tree *datastructures.TestTreeNode) [][]int {
	if tree == nil {
		return make([][]int, 0)
	}
	res := make([][]int, 0)
	// level traversal
	queue := datastructures.Queue{}
	queue.Enqueue(tree)
	// direction
	rev := false
	for !queue.Empty() {
		leng := queue.Len()
		tmp := make([]int, leng)
		// per level
		for i := 0; i < leng; i++ {
			root := queue.Dequeue().(*datastructures.TestTreeNode)
			j := i
			if rev {
				j = leng - i - 1
			}
			tmp[j] = root.Value
			if root.Left != nil {
				queue.Enqueue(root.Left)
			}
			if root.Right != nil {
				queue.Enqueue(root.Right)
			}
		}
		res = append(res, tmp)
		rev = !rev
	}
	return res
}
