package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回 。
//注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。

func main() {
	// 构造二叉树
	tree := datastructures.BuildTreeFromArray([]interface{}{5, 4, nil, 3, nil, 2})
	nextNode := findNextNode(tree)
	fmt.Printf("%v\n", nextNode)
}
func findNextNode(node *datastructures.TestTreeNode) *datastructures.TestTreeNode {
	// have right child
	if node.Right != nil {
		cur := node.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}
	// no right child ,but it is parent's left child
	if node.Parent != nil && node.Parent.Left == node {
		return node.Parent
	}
	// no right child , but it is parent's right child
	if node.Parent != nil && node.Parent.Right == node {
		cur := node
		for cur.Parent != nil && cur.Parent.Left != cur {
			cur = cur.Parent
		}
		return cur.Parent
	}
	return nil
}
