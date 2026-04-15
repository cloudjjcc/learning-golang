package leetcode

import "math/rand"

func sortColors(nums []int) {
	partition := func(left, right int) (int, int) {
		if left == right {
			return left, left
		}
		randIdx := left + rand.Intn(right-left+1)
		pivot := nums[randIdx]
		nums[left], nums[randIdx] = nums[randIdx], nums[left]
		lt, i, gt := left, left+1, right
		// nums[left:lt) 小于pivot的元素区间
		// nums[lt:i)等于pivot的元素区间
		// nums[i:gt]待处理的元素区间
		// nums(gt:right]大于pivot的元素区间
		for i <= gt {
			if nums[i] < pivot {
				nums[lt], nums[i] = nums[i], nums[lt]
				lt++
			} else if nums[i] > pivot {
				nums[gt], nums[i] = nums[i], nums[gt]
				gt--
			} else {
				i++
			}
		}
		return lt, gt
	}
	left, right := 0, len(nums)-1
	lt, gt := partition(left, right)
	sortColors(nums[left:lt])
	sortColors(nums[gt+1 : right+1])
}
