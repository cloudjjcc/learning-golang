package leetcode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = int(matrix[i][0] - '0')
	}
	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
	}
	maxEdge := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
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
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 {
				dp[0][j] = 1
				maxEdge = max(maxEdge, 1)
				continue
			}
			if j == 0 {
				dp[0][j] = 1
				maxEdge = max(maxEdge, 1)
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			maxEdge = max(maxEdge, dp[i][j])
		}
	}
	return maxEdge * maxEdge
}
