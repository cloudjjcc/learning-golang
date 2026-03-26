package leetcode

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	tmp := intervals[0]
	ans := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > tmp[1] { //不能合并
			ans = append(ans, tmp)
			tmp = intervals[i]
			continue
		}
		if intervals[i][1] > tmp[1] {
			tmp = []int{tmp[0], intervals[i][1]}
		}
	}
	ans = append(ans, tmp)
	return ans
}
