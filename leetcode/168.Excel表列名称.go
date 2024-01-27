package leetcode

import "unsafe"

func convertToTitle(columnNumber int) string {
	var buf []byte
	for columnNumber > 0 {
		columnNumber--
		buf = append(buf, byte(columnNumber%26+'A'))
		columnNumber /= 26
	}
	// reverse
	left, right := 0, len(buf)-1
	for left < right {
		buf[left], buf[right] = buf[right], buf[left]
		left++
		right--
	}
	return *(*string)(unsafe.Pointer(&buf))
}
