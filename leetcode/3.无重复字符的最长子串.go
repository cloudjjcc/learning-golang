package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	left, right := 0, 0
	max := 0
	m := make(map[byte]int)
	for ; right < len(s); right++ {
		if i, ok := m[s[right]]; ok && i >= left {
			left = i + 1
		}
		m[s[right]] = right
		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}
