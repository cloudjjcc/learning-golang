package huawei

import (
	"fmt"
	"unsafe"
)

func mainconvert() {
	var a int
	if _, err := fmt.Scan(&a); err != nil {
		return
	}
	s := convert(a)
	fmt.Println(s)
}

func convert(a int) string {
	var ans []byte
	for a != 0 {
		a -= 1
		ans = append(ans, byte(a%26+'a'))
		a /= 26
	}
	left, right := 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return *(*string)(unsafe.Pointer(&ans))
}
