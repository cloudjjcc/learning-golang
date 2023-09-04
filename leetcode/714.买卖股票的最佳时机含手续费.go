package leetcode

// 动态规划
func maxProfit1(prices []int, fee int) int {
	buy, sell := -prices[0], 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(prices); i++ {
		sell = max(sell, buy+prices[i]-fee)
		buy = max(buy, sell-prices[i])
	}
	return sell
}
