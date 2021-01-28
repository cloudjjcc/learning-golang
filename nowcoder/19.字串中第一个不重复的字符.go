package main

import "fmt"

//题目描述
//请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符"go"时，第一个只出现一次的字符是"g"。当从该字符流中读出前六个字符“google"时，第一个只出现一次的字符是"l"。
//输出描述:
//如果当前字符流没有存在出现一次的字符，返回#字符。

func main() {
	fmt.Println(string(FirstAppearingOnce("google")))
}

func FirstAppearingOnce(str string) byte {
	if len(str) == 0 {
		return '#'
	}
	if len(str) == 1 {
		return str[0]
	}
	tmp := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		tmp[str[i]]++
	}
	for i := 0; i < len(str); i++ {
		if tmp[str[i]] == 1 {
			return str[i]
		}
	}
	return '#'
}
