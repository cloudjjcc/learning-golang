package huawei

import (
	"fmt"
	"strconv"
	"strings"
)

// 例子：
//2,10,-3,-8,40,5
//4
//39

func mainparseIntSlice() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	arr := parseIntSlice(s, ",")
	var k int
	_, err = fmt.Scanln(&k)
	if err != nil {
		return
	}
	sum := findMaxSum(arr, k)
	fmt.Println(sum)
}
func parseIntSlice(s string, split string) []int {
	ss := strings.Split(s, split)
	ans := make([]int, 0, len(ss))
	for _, v := range ss {
		vv, _ := strconv.Atoi(v)
		ans = append(ans, vv)
	}
	return ans
}
func findMaxSum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	startIdx := 0
	sum := 0
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, v := range nums {
		sum += v
		if i-startIdx+1 == k {
			ans = max(ans, sum)
			sum -= nums[startIdx]
			startIdx++
		}
	}
	return ans
}
