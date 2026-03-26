package leetcode

// dp
func maxProfit333(prices []int) int {
	buy1 := -prices[0] //完成一次买入
	sell1 := 0         //完成一次交易
	buy2 := -prices[0] //完成一次交易一次买入
	sell2 := 0         //完成两次交易
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
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return max(sell1, sell2)
}
