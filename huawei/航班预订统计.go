package main

import "fmt"

//这里有 n 个航班，它们分别从 1 到 n 进行编号。
//我们这儿有一份航班预订表，表中第 i 条预订记录 bookings[i] = [i, j, k] 意味着我们在从 i 到 j 的每个航班上预订了 k 个座位。
//请你返回一个长度为 n 的数组 answer，按航班编号顺序返回每个航班上预订的座位数。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/corporate-flight-bookings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	testArr := [][]int{{1, 2, 3}, {2, 3, 5}}
	testN := 3
	fmt.Println(corpFlightBookings(testArr, testN))
	fmt.Println(corpFlightBookings2(testArr, testN))
}

// 暴力法
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, v := range bookings {
		for i := v[0] - 1; i < v[1]; i++ {
			res[i] += v[2]
		}
	}
	return res
}

// 利用查分序列
func corpFlightBookings2(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, v := range bookings {
		res[v[0]-1] += v[2]
		if v[1] < n {
			res[v[1]] -= v[2]
		}
	}
	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}
	return res
}
