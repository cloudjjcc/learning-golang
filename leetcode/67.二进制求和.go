package leetcode

import "unsafe"

func addBinary(a string, b string) string {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	sumBytes := make([]byte, maxLen+1)
	carry := byte(0)
	for i := 0; i < len(sumBytes); i++ {
		sum := carry
		if i < len(a) {
			sum += a[len(a)-1-i] - byte('0')
		}
		if i < len(b) {
			sum += b[len(b)-1-i] - byte('0')
		}
		carry = sum >> 1
		sumBytes[len(sumBytes)-1-i] = sum&1 + '0'
	}
	if sumBytes[0] == '0' {
		sumBytes = sumBytes[1:]
	}
	return *(*string)(unsafe.Pointer(&sumBytes))
}
