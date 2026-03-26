package leetcode

import (
	"strconv"
)

func compressString(str string) string {
	if len(str) == 0 {
		return ""
	}
	res := make([]byte, 0)
	repeat := 0
	for i := 0; i < len(str); i++ {
		if i == 0 {
			repeat++
			res = append(res, str[i])
			continue
		}
		if str[i] == str[i-1] {
			repeat++
		} else {
			res = append(res, strconv.Itoa(repeat)...)
			res = append(res, str[i])
			repeat = 1
		}
	}
	res = append(res, strconv.Itoa(repeat)...)
	if len(res) >= len(str) {
		return str
	}
	return string(res)
}
