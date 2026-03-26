package leetcode

import "math"

func minWindow(s string, t string) string {
	tChars := make(map[byte]int)
	for _, v := range t {
		tChars[byte(v)]++
	}
	minLen := math.MaxInt64
	minL, minR := 0, 0
	left, right := 0, 0
	countM := make(map[byte]int)
	containsFn := func() bool {
		if len(countM) < len(tChars) {
			return false
		}
		for k, v := range tChars {
			if countM[k] < v {
				return false
			}
		}
		return true
	}
	for left <= right && right < len(s) {
		if tChars[s[right]] > 0 {
			countM[s[right]]++
		}
		for containsFn() {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minL, minR = left, right
			}
			if countM[s[left]] > 0 {
				countM[s[left]]--
			}
			left++
		}
		right++
	}
	if minLen == math.MaxInt64 {
		return ""
	}
	return s[minL : minR+1]
}
