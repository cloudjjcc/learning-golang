package leetcode

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	ans := 0

	for i := len(citations) - 1; i >= 0 && citations[i] > ans; i-- {
		ans++
	}
	return ans
}
