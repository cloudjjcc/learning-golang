package leetcode

import "unsafe"

func reverseStr(s string, k int) string {
	ans := make([]byte, len(s))
	copy(ans, s)
	reverseFn := func(start, end int) {
		for start < end {
			ans[start], ans[end] = ans[end], ans[start]
			start++
			end--
		}
	}
	i := 0
	for ; i < len(s); i += 2 * k {
		if (len(s) - i) >= k {
			reverseFn(i, i+k-1)
		} else {
			reverseFn(i, len(s)-1)
		}
	}
	return *(*string)(unsafe.Pointer(&ans))
}
