package leetcode

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) <= 2 {
		return len(s)
	}
	chMap := make(map[byte]int)
	left, right := 0, 0
	maxLen := 2
	for right < len(s) {
		chMap[s[right]] = right
		if len(chMap) > 2 {
			minIdx := len(s)
			for _, v := range chMap {
				if v < minIdx {
					minIdx = v
				}
			}
			delete(chMap, s[minIdx])
			left = minIdx + 1
		}
		if ll := right - left + 1; ll > maxLen {
			maxLen = ll
		}
		right++
	}
	return maxLen
}
