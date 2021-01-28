package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//从上往下打印出二叉树的每个节点，同层节点从左至右打印。

func main() {
	testTree := &tnode{Value: 0}
	tmpNode := testTree
	for i := 1; i < 10; i++ {
		if i&1 == 0 {
			if tmpNode.right != nil {
				tmpNode = tmpNode.right
			} else {
				tmpNode.right = &tnode{Value: i}
			}
		} else {
			if tmpNode.left != nil {
				tmpNode = tmpNode.left
			} else {
				tmpNode.left = &tnode{Value: i}
			}
		}
	}
	levelOrderTraversal(testTree, func(t *tnode) {
		fmt.Println(t.Value)
	})
}

type tnode struct {
	left  *tnode
	right *tnode
	Value interface{}
}
type consumer = func(*tnode)

// 层序遍历（利用队列）
func levelOrderTraversal(tree *tnode, con consumer) {
	if tree == nil {
		return
	}
	queue := new(datastructures.Queue)
	queue.Enqueue(tree)
	for !queue.Empty() {
		tmpNode := queue.Dequeue().(*tnode)
		con(tmpNode)
		if tmpNode.left != nil {
			queue.Enqueue(tmpNode.left)
		}
		if tmpNode.right != nil {
			queue.Enqueue(tmpNode.right)
		}
	}
}
