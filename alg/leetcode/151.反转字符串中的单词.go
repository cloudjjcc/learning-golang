package leetcode

import "unsafe"

func reverseWords1(s string) string {
	var words []string
	// get words
	start := -1
	targetLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && start == -1 {
			start = i
		}
		if s[i] == ' ' && start != -1 {
			words = append(words, s[start:i])
			targetLen += i - start
			start = -1
		}
	}
	if start != -1 && len(s)-start > 0 {
		words = append(words, s[start:])
		targetLen += len(s) - start
	}
	if len(words) > 1 {
		targetLen += len(words) - 1
	}
	// revers words
	ans := make([]byte, 0, targetLen)
	for i := len(words) - 1; i >= 0; i-- {
		ans = append(ans, words[i]...)
		if i != 0 {
			ans = append(ans, ' ')
		}
	}
	return *(*string)(unsafe.Pointer(&ans))
}
