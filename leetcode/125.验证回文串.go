package leetcode

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	validCh := func(v byte) bool {
		if (v >= '0' && v <= '9') || v >= 'a' && v <= 'z' {
			return true
		}
		return false
	}
	left, right := 0, len(s)-1
	for left < right {
		if !validCh(s[left]) {
			left++
			continue
		}
		if !validCh(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
