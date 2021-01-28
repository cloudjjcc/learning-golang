package main

import "fmt"

//题目描述
//输入一个整数，输出该数二进制表示中 1 的个数。
func main() {
	fmt.Println(numsOf1(0x1010111))
}

// 返回1的个数
func numsOf1(i int) int {
	var count int
	for i != 0 {
		i &= i - 1
		count++
	}
	return count
}
