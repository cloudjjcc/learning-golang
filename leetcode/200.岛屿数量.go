package leetcode

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	ans := 0
	dfs = func(ti, tj int) {
		if grid[ti][tj] == '0' {
			return
		}
		grid[ti][tj] = '0' //遍历过的位置设置为0
		//向上遍历
		if ii := ti - 1; ii >= 0 {
			dfs(ii, tj)
		}
		//向下遍历
		if ii := ti + 1; ii < m {
			dfs(ii, tj)
		}
		//向左遍历
		if jj := tj - 1; jj >= 0 {
			dfs(ti, jj)
		}
		//向右遍历
		if jj := tj + 1; jj < n {
			dfs(ti, jj)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
