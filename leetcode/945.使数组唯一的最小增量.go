package main

import "fmt"

//给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。
//返回使 A 中的每个值都是唯一的最少操作次数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}

func minIncrementForUnique(arr []int) int {
	count := make([]int, 80000)
	for _, v := range arr {
		count[v]++
	}
	ans, token := 0, 0
	for i := 0; i < 80000; i++ {
		if count[i] >= 2 {
			token += count[i] - 1
			ans -= i * (count[i] - 1)
		} else if token > 0 && count[i] == 0 {
			token--
			ans += i
		}
	}
	return ans
}
