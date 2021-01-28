package main

import (
	"fmt"
)

//题目描述
//给定一个无序数组，包含正数、负数和0，要求从中找出3个数的乘积，使得乘积最大，
//要求时间复杂度：O(n)，空间复杂度：O(1)
//输入描述:
//输入共2行，第一行包括一个整数n，表示数组长度
//第二行为n个以空格隔开的整数，分别为A1,A2, … ,An
//输出描述:
//满足条件的最大乘积

func main() {
	testArr := []int{0, -1, 2, 3, -3, 9}
	fmt.Println(maxProduct(testArr))
}

// 最大的乘积：
// 最大的三个非零数
// 最小的两个非零数*最大非零数
func maxProduct(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	if len(arr) == 3 {
		return arr[0] * arr[1] * arr[2]
	}
	// 最大的三个非零数
	for i := 0; i < 3; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] != 0 && (arr[j-1] == 0 || arr[j] > arr[j-1]) {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	// 最小的两个非零数
	for i := 0; i < 2; i++ {
		for j := 4; j < len(arr)-1-i; j++ {
			if arr[j] != 0 && (arr[j+1] == 0 || arr[j] < arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	max1 := arr[0] * arr[1] * arr[2]
	max2 := arr[0] * arr[len(arr)-1] * arr[len(arr)-2]
	if max1 > max2 {
		return max1
	}
	return max2
}
