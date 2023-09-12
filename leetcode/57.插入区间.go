package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	ans := make([][]int, 0)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	merged := false
	add := false
	for i := 0; i < len(intervals); i++ {
		t := intervals[i]
		if !add && (newInterval[0] <= t[0] && newInterval[1] >= t[0] ||
			newInterval[0] <= t[1] && newInterval[1] >= t[0] ||
			newInterval[0] >= t[0] && newInterval[1] <= t[1]) {
			newInterval[0] = min(newInterval[0], t[0])
			newInterval[1] = max(newInterval[1], t[1])
			merged = true
		} else {
			if !add && newInterval[1] < t[0] {
				add = true
				ans = append(ans, newInterval)
			}
			if merged && !add {
				add = true
				ans = append(ans, newInterval)
			}
			ans = append(ans, t)
		}
	}
	if !add {
		ans = append(ans, newInterval)
	}
	return ans
}
