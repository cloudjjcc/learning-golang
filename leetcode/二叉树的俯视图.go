package leetcode

import "sort"

func topView(root *TreeNode) []int {
	ans := make([]int, 0)
	m := make(map[int][2]int) // key:pos val[0]:val val[1]:level
	posArr := make([]int, 0)
	var dfs func(n *TreeNode, pos int, level int)
	dfs = func(n *TreeNode, pos int, level int) {
		if n == nil {
			return
		}
		if v, ok := m[pos]; !ok || level < v[1] {
			m[pos] = [2]int{n.Val, level}
			posArr = append(posArr, pos)
		}
		dfs(n.Left, pos-1, level+1)
		dfs(n.Right, pos+1, level+1)
	}
	dfs(root, 0, 0)
	sort.Ints(posArr)
	for _, v := range posArr {
		ans = append(ans, m[v][0])
	}
	return ans
}
