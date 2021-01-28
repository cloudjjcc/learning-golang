package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
)

//题目描述
//输入一颗二叉树的根节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
//路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
//(注意: 在返回值的list中，数组长度大的数组靠前)

func main() {
	preorder := []int{0, 1, 3, 5, 2, 4, 6}
	inorder := []int{3, 1, 5, 0, 4, 2, 6}
	testTree := datastructures.BuildTreeFromOrder(preorder, inorder)
	testTree.LevelOrder(func(i int) {
		fmt.Println(i)
	})
	fmt.Println(findPath(testTree, 6))
}

func findPath(tree *datastructures.TestTreeNode, target int) [][]int {
	return dfs(tree, []int{}, target)
}

func dfs(tree *datastructures.TestTreeNode, path []int, target int) [][]int {
	if tree == nil {
		return nil
	}
	var res [][]int
	path = append(path, tree.Value)
	if tree.Left == nil && tree.Right == nil {
		pathSum := 0
		for _, v := range path {
			pathSum += v
		}
		if pathSum == target {
			res = append(res, path)
		}
	} else {
		res = append(res, dfs(tree.Left, path, target)...)
		res = append(res, dfs(tree.Right, path, target)...)
	}
	return res
}
