package leetcode

// 动态规划（优化）
func maxProfit33(prices []int) int {
	preNo := 0
	preHave := -prices[0]
	max := func(a int, b ...int) int {
		max := a
		for _, v := range b {
			if v > max {
				max = v
			}
		}
		return max
	}
	for i := 1; i < len(prices); i++ {
		preNo = max(preNo, preHave+prices[i])
		preHave = max(preHave, preNo-prices[i])
	}
	return preNo
}

func maxProfit3(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	max := func(a int, b ...int) int {
		max := a
		for _, v := range b {
			if v > max {
				max = v
			}
		}
		return max
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
