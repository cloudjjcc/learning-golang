package nowcoder

//题目描述
//地上有一个 m 行和 n 列的方格。
//一个机器人从坐标 (0, 0) 的格子开始移动，每一次只能向左右上下四个方向移动一格，
//但是不能进入行坐标和列坐标的数位之和大于 k 的格子。
//例如，当 k 为 18 时，机器人能够进入方格 (35,37)，因为 3+5+3+7=18。
//但是，它不能进入方格 (35,38)，因为 3+5+3+8=19。
//请问该机器人能够达到多少个格子？

func moveCount(k int, rows int, cols int) int {
	if k == 0 {
		return 1
	}
	count := 0
	visit := make([]bool, rows*cols)
	dfsCore(visit, k, rows, cols, 0, 0, &count)
	return count
}

func dfsCore(visit []bool, k int, rows int, cols int, r int, c int, count *int) {
	// check index range
	if r < 0 || r > rows-1 || c < 0 || c > cols-1 {
		return
	}
	cur := r*cols + c
	// get a and b bit sum
	bitSum := func(a, b int) int {
		sum := 0
		for a != 0 {
			sum += a % 10
			a /= 10
		}
		for b != 0 {
			sum += b % 10
			b /= 10
		}
		return sum
	}
	if visit[cur] {
		return
	}
	visit[cur] = true
	if bitSum(r, c) <= k {
		*count++
		dfsCore(visit, k, rows, cols, r-1, c, count)
		dfsCore(visit, k, rows, cols, r+1, c, count)
		dfsCore(visit, k, rows, cols, r, c-1, count)
		dfsCore(visit, k, rows, cols, r, c+1, count)
	}
}
