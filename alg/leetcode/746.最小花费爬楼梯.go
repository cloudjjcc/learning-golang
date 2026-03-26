package leetcode

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	minFn := func(a int, b ...int) int {
		min := a
		for _, v := range b {
			if v < min {
				min = v
			}
		}
		return min
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}
	if n == 2 {
		return minFn(cost[0], cost[1])
	}
	// F(i)=min(F(i-1)+cost[i-1],F(i-2)+cost[i-2]
	n1, n2, cur := cost[0], cost[1], 0
	for i := 2; i <= n; i++ {
		cur = minFn(n1+cost[i-1], n2+cost[i-2])
		n1, n2 = cur, n1
	}
	return cur
}
