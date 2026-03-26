package leetcode

import "unsafe"

func dynamicPassword(password string, target int) string {
	n := len(password)
	arr := []byte(password)
	// 翻转区间
	reverse := func(a, b int) {
		left, right := a, b
		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	reverse(0, target-1)
	reverse(target, n-1)
	reverse(0, n-1)
	return *(*string)(unsafe.Pointer(&arr))
}
