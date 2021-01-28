package main

import (
	"fmt"
	"strings"
)

//判断字符串中字符是否全都不同
//问题描述
//请实现一个算法，确定一个字符串的所有字符【是否全都不同】。
//这里我们要求【不允许使用额外的存储结构】。
//给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
//保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】

// 利用strings.Count 函数
func IsAllLetterDifferent(str string) bool {
	for _, chars := range str {
		if strings.Count(str, string(chars)) > 1 {
			return false
		}
	}
	return true
}

// 利用strings.Index 和 strings.LastIndex
func IsAllLetterDifferent2(str string) bool {
	for _, chars := range str {
		if strings.Index(str, string(chars)) != strings.LastIndex(str, string(chars)) {
			return false
		}
	}
	return true
}
func main() {
	testStr := "abcadef"
	fmt.Printf("%s : %t \n", testStr, IsAllLetterDifferent(testStr))
	fmt.Printf("%s : %t \n", testStr, IsAllLetterDifferent2(testStr))
}
