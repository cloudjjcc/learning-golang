package leetcode

import (
	"math/rand"
)

// 快速排序
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	partition := func() int {
		pivotIdx := rand.Int() % len(nums)
		pivotVal := nums[pivotIdx]
		nums[pivotIdx], nums[len(nums)-1] = nums[len(nums)-1], nums[pivotIdx]
		left, right := 0, 0
		for right < len(nums)-1 {
			if nums[right] < pivotVal {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
			right++
		}
		nums[left], nums[len(nums)-1] = nums[len(nums)-1], nums[left]
		return left
	}
	pivot := partition()
	sortArray(nums[:pivot])
	sortArray(nums[pivot+1:])
	return nums
}

// 归并排序
func sortArray2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	sortArray(nums[:mid])
	sortArray(nums[mid:])
	return mergeArr(nums[:mid], nums[mid:])
}
func mergeArr(arr1, arr2 []int) []int {
	ans := make([]int, 0, len(arr1)+len(arr2))
	i1, i2 := 0, 0
	for {
		if i1 == len(arr1) {
			ans = append(ans, arr2[i2:]...)
			break
		}
		if i2 == len(arr2) {
			ans = append(ans, arr1[i1:]...)
			break
		}
		if arr1[i1] < arr2[i2] {
			ans = append(ans, arr1[i1])
			i1++
		} else {
			ans = append(ans, arr2[i2])
			i2++
		}
	}
	return ans
}
