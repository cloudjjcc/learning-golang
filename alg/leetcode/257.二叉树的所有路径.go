package leetcode

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		if path == "" {
			path = strconv.Itoa(node.Val)
		} else {
			path = path + "->" + strconv.Itoa(node.Val)
		}
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path)
			return
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return ans
}
