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

// dfs
func numIslands2(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	queue := make([][2]int, 0)
	push := func(v [2]int) {
		queue = append(queue, v)
	}
	pop := func() [2]int {
		head := queue[0]
		queue = queue[1:]
		return head
	}
	empty := func() bool {
		return len(queue) == 0
	}
	size := func() int {
		return len(queue)
	}
	ans := 0
	var bfs func(i, j int)
	bfs = func(i, j int) {
		queue = queue[:0]
		push([2]int{i, j})
		for !empty() {
			s := size()
			for z := 0; z < s; z++ {
				xy := pop()
				grid[xy[0]][xy[1]] = '0'
				// 上
				if xy[0]-1 >= 0 &&
					grid[xy[0]-1][xy[1]] == '1' {
					push([2]int{xy[0] - 1, xy[1]})
				}
				// 下
				if xy[0]+1 < m &&
					grid[xy[0]+1][xy[1]] == '1' {
					push([2]int{xy[0] + 1, xy[1]})
				}
				// 左
				if xy[1]-1 >= 0 &&
					grid[xy[0]][xy[1]-1] == '1' {
					push([2]int{xy[0], xy[1] - 1})
				}
				// 右
				if xy[1]+1 < n &&
					grid[xy[0]][xy[1]+1] == '1' {
					push([2]int{xy[0], xy[1] + 1})
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			ans++
			bfs(i, j)
		}
	}
	return ans
}
