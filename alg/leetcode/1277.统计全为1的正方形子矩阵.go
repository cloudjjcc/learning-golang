package leetcode

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	min := func(a int, b ...int) int {
		t := a
		for _, v := range b {
			if v < t {
				t = v
			}
		}
		return t
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
				ans += 1
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			ans += dp[i][j]
		}
	}
	return ans
}
