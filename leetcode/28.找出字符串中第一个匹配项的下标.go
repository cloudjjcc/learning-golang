package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	getNext := func(s string) []int {
		next := make([]int, len(s))
		next[0] = 0
		j := 0
		for i := 1; i < len(s); i++ {
			for j > 0 && s[i] != s[j] {
				j = next[j-1]
			}
			if s[i] == s[j] {
				j++
			}
			next[i] = j
		}
		return next
	}
	next := getNext(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
