package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			if nums[j]+nums[k] == (-nums[i]) {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				continue
			}
			if nums[j]+nums[k] > (-nums[i]) {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
