package leetcode

func longestPalindromeSubseq(s string) int {
	m := len(s)
	dp := make([][]int, m) //dp[i][j]表示s[i:j+1]的最大回文子串
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := m - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < m; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][m-1]
}
