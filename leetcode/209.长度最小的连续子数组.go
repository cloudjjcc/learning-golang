package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	slow, fast := 0, 0
	sum := 0
	minLen := math.MaxInt64
	for ; fast < len(nums); fast++ {
		sum += nums[fast]
		if sum < target {
			continue
		}
		for sum > target {
			sum -= nums[slow]
			slow++
		}
		if sum < target {
			slow--
			sum += nums[slow]
		}
		if ll := fast - slow + 1; ll < minLen {
			minLen = ll
		}
	}
	if minLen == math.MaxInt64 {
		return 0
	}
	return minLen
}
