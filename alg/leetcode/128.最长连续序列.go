package leetcode

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}
	maxLen := 0
	for _, v := range nums {
		if _, ok := m[v-1]; !ok {
			i := v
			for {
				i++
				if _, ok := m[i]; !ok {
					break
				}
			}
			if i-v > maxLen {
				maxLen = i - v
			}
		}
	}
	return maxLen
}
