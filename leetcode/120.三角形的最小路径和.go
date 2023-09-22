package leetcode

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, i+1)
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
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j]+triangle[i][j], dp[i-1][j-1]+triangle[i][j])
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	return min(dp[m-1][0], dp[m-1][1:]...)
}
