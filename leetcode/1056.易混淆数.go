package leetcode

import (
	"strconv"
)

// 给定一个数字 N，当它满足以下条件的时候返回 true：
//
// 原数字旋转 180° 以后可以得到新的数字。
//
// 如 0, 1, 6, 8, 9 旋转 180° 以后，得到了新的数字 0, 1, 9, 8, 6 。
//
// 2, 3, 4, 5, 7 旋转 180° 后，得到的不是数字。
//
// 易混淆数 (confusing number) 在旋转180°以后，可以得到和原来不同的数，且新数字的每一位都是有效的。
func confusingNumber(n int) bool {
	nStr := strconv.Itoa(n)
	m := []int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6}
	hasC := false
	for i := 0; i < len(nStr); i++ {
		t1 := nStr[i] - '0'
		t2 := nStr[len(nStr)-1-i] - '0'

		if m[t1] == -1 || m[t2] == -1 {
			return false
		}
		if int(t1) != m[t2] {
			hasC = true
		}
	}
	return hasC
}
