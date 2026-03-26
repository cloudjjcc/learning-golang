package pkgdemo

import "sort"

func Search(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

func Sort(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}
