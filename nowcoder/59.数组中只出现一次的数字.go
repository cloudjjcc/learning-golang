package main

import "fmt"

//题目描述
//一个整型数组里除了两个数字之外，其他的数字都出现了两次。
//请写程序找出这两个只出现一次的数字。

func main() {
	fmt.Println(findOnceNum([]int{1, 2, 2, 3, 4, 4}))
}

func findOnceNum(arr []int) []int {
	if len(arr) < 2 {
		return []int{}
	}
	// xor all ele
	xor := 0
	for i := 0; i < len(arr); i++ {
		xor ^= arr[i]
	}
	// get first bit 1
	index := 1
	for xor&index == 0 {
		index <<= 1
	}
	// split arr
	res1 := 0
	res2 := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&index == 0 {
			res1 ^= arr[i]
		} else {
			res2 ^= arr[i]
		}
	}
	return []int{res1, res2}
}
