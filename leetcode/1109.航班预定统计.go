package leetcode

// 暴力
func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, v := range bookings {
		for i := v[0]; i <= v[1]; i++ {
			ans[i] += v[2]
		}
	}
	return ans
}

// 差分
func corpFlightBookings2(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, v := range bookings {
		ans[v[0]-1] += v[2]
		if v[1] < n {
			ans[v[1]] -= v[2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}
