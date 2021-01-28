package main

import (
	"fmt"
	ds "github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。
//要求不能创建任何新的结点，只能调整树中结点指针的指向。

func main() {
	tree := ds.BuildTreeFromArray([]interface{}{10, 8, 12, 7, 9, 11, 13})
	dlinkedlist := convert(tree)
	fmt.Println(dlinkedlist.Value)
}
func convert(tree *ds.TestTreeNode) *ds.TestTreeNode {
	if tree == nil {
		return nil
	}
	// inOrder traversal  tree
	nodes := make([]*ds.TestTreeNode, 0)
	nodes = append(nodes, nil)
	tree.InOrder(func(node *ds.TestTreeNode) {
		nodes = append(nodes, node)
	})
	nodes = append(nodes, nil)
	// modify pointer left->prev right -> next
	for i := 1; i < len(nodes)-1; i++ {
		nodes[i].Left = nodes[i-1]
		nodes[i].Right = nodes[i+1]
	}
	return nodes[1]
}
