package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"strconv"
)

//题目描述
//请实现两个函数，分别用来序列化和反序列化二叉树
//二叉树的序列化是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，
//从而使得内存中建立起来的二叉树可以持久保存。
//序列化可以基于先序、中序、后序、层序的二叉树遍历方式来进行修改，
//序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#），
//以 ！ 表示一个结点值的结束（value!）。
//二叉树的反序列化是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。

func main() {
	testTreeStr := "0!#1!#2!#3!#4!#"
	tree := deserializable([]byte(testTreeStr))
	fmt.Println(string(serializable(tree)))
}

// 序列化二叉树
func serializable(tree *datastructures.TestTreeNode) []byte {
	buf := make([]byte, 0)
	if tree == nil {
		return append(buf, '#')
	}
	queue := &datastructures.Queue{}
	queue.Enqueue(tree)
	buf = append(buf, fmt.Sprintf("%d!", tree.Value)...)
	for !queue.Empty() {
		root := queue.Dequeue().(*datastructures.TestTreeNode)
		if root.Left != nil {
			queue.Enqueue(root.Left)
			buf = append(buf, fmt.Sprintf("%d!", root.Left.Value)...)
		} else {
			return append(buf, '#')
		}
		if root.Right != nil {
			queue.Enqueue(root.Right)
			buf = append(buf, fmt.Sprintf("%d!", root.Right.Value)...)
		} else {
			return append(buf, '#')
		}
	}
	return buf
}

// 反序列化二叉树
func deserializable(str []byte) *datastructures.TestTreeNode {
	if len(str) == 0 {
		return new(datastructures.TestTreeNode)
	}
	word := make([]byte, 0)
	nodes := make(map[int]*datastructures.TestTreeNode)
	nodeIdx := 0
	for _, v := range str {
		switch v {
		case '!':
			value, err := strconv.Atoi(string(word))
			if err != nil {
				continue
			}
			nodes[nodeIdx] = &datastructures.TestTreeNode{Value: value}
			fallthrough
		case '#':
			nodeIdx++
			word = word[:0]
		default:
			word = append(word, v)
		}
	}
	for i := 0; i < nodeIdx; i++ {
		if v, ok := nodes[i]; ok {
			v.Left = nodes[2*i+1]
			v.Right = nodes[2*i+2]
		}
	}
	return nodes[0]
}

// TODO
