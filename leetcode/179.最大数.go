package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		sa, sb := 10, 10
		for sa <= a {
			sa *= 10
		}
		for sb <= b {
			sb *= 10
		}
		return a*sb+b > b*sa+a
	})
	var sb strings.Builder
	for _, v := range nums {
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}
