package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	left, right := 0, 1
	var ans []string
	addAns := func(a, b int) {
		t := ""
		if a == b {
			t = fmt.Sprintf("%d", a)
		} else {
			t = fmt.Sprintf("%d->%d", a, b)
		}
		ans = append(ans, t)
	}
	for right < len(nums) {
		if nums[right] > nums[right-1]+1 {
			addAns(nums[left], nums[right-1])
			left = right
		}
		right++
	}
	addAns(nums[left], nums[right-1])
	return ans
}
