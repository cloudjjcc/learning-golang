package leetcode

import "math"

func maxProfit4(k int, prices []int) int {
	buyState := make([][]int, len(prices))  //持有
	sellState := make([][]int, len(prices)) //未持有
	for i := range buyState {
		buyState[i] = make([]int, k+1)
		sellState[i] = make([]int, k+1)
	}
	max := func(a int, b ...int) int {
		max := a
		for _, v := range b {
			if v > max {
				max = v
			}
		}
		return max
	}
	buyState[0][0] = -prices[0]
	for i := 0; i <= k; i++ {
		buyState[0][i] = -prices[0]
		sellState[0][i] = 0
	}
	for i := 1; i < len(prices); i++ {
		buyState[i][0] = max(buyState[i-1][0], -prices[i])
		sellState[i][0] = max(sellState[i-1][0])
		for j := 1; j <= k; j++ {
			buyState[i][j] = max(buyState[i-1][j], sellState[i-1][j]-prices[i])
			sellState[i][j] = max(sellState[i-1][j], buyState[i-1][j]+prices[i])
		}
	}
	return max(math.MinInt64, sellState[len(sellState)-1]...)
}
