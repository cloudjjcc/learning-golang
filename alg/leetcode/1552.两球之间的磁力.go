package leetcode

import "sort"

func maxDistance1(position []int, m int) int {
	sort.Ints(position)
	left, right := 1, position[len(position)-1]-position[0]
	ans := -1
	check := func(limit int) bool {
		pre := position[0]
		cnt := 1
		for i := 1; i < len(position); i++ {
			if position[i]-pre >= limit {
				cnt++
				pre = position[i]
			}
		}
		return cnt >= m
	}
	for left <= right {
		mid := (left + right) >> 1
		if check(mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans

}
