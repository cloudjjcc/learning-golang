package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//给定一棵二叉搜索树，请找出其中的第k小的结点。
//例如， （5，3，7，2，4，6，8）    中，按结点数值大小顺序第三小结点的值为4。

func main() {
	tree := datastructures.BuildTreeFromArray([]interface{}{4, 2, 6, 1, 3, 5, 7})
	fmt.Println(kthNode(tree, 7))
}

func kthNode(tree *datastructures.TestTreeNode, k int) *datastructures.TestTreeNode {
	if tree == nil {
		return nil
	}
	dfsFindK := func(tree *datastructures.TestTreeNode, k int) *datastructures.TestTreeNode {
		if tree == nil {
			return nil
		}
		// inOrder traversal
		stack := datastructures.Stack{}
		cur := tree
		count := 0
		for cur != nil || !stack.Empty() {
			for cur != nil {
				stack.Push(cur)
				cur = cur.Left
			}
			cur = stack.Pop().(*datastructures.TestTreeNode)
			count++
			if count == k {
				return cur
			}
			cur = cur.Right
		}
		return nil
	}
	return dfsFindK(tree, k)
}
