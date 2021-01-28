package main

import (
	"fmt"
	"strings"
	"unicode"
)

//字符串替换问题
//问题描述
//请编写一个方法，将字符串中的空格全部替换为“%20”。
//假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
//给定一个string为原始的串，返回替换后的string。

func ReplaceStr(str string) (string, bool) {
	for _, r := range str {
		if r != ' ' && !unicode.IsLetter(r) {
			return str, false
		}
	}
	return strings.ReplaceAll(str, " ", "%20"), false
}
func main() {
	testStr := "absd sd"
	r, b := ReplaceStr(testStr)
	fmt.Printf("%s:%s:%t", testStr, r, b)
}
