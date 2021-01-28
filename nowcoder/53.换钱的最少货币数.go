package main

import "fmt"

//题目描述
//给定数组arr，arr中所有的值都为正整数且不重复。
//每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，
//求组成aim的最少货币数。
//输入描述:
//输入包括两行，第一行两个整数n（0<=n<=1000）代表数组长度和aim（0<=aim<=5000），
//第二行n个不重复的正整数，代表arr
//输出描述:
//输出一个整数，表示组成aim的最小货币数，无解时输出-1.
func main() {
	testCoins := []int{2, 5, 10}
	fmt.Println(minCoinNum(testCoins, 8))
	fmt.Println(minCoinNum(testCoins, 9))
	fmt.Println(minCoinNum(testCoins, 10))
	fmt.Println(minCoinNum(testCoins, 99))
}

//f(x)=min{f{x-i}+1},i ∈ coinFace
func minCoinNum(coins []int, aim int) int {
	coinNums := make([]int, aim+1)
	coinNums[0] = 0
	for i := 1; i <= aim; i++ {
		// get min{f(x-i)}
		preMin := -1
		for j := 0; j < len(coins); j++ {
			if preAim := i - coins[j]; preAim >= 0 && coinNums[preAim] != -1 {
				if preMin == -1 {
					preMin = coinNums[preAim]
				} else {
					preMin = min(preMin, coinNums[preAim])
				}
			}
		}
		if preMin == -1 {
			coinNums[i] = -1
		} else {
			coinNums[i] = preMin + 1
		}
	}
	return coinNums[aim]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
