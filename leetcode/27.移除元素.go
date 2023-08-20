package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] != val {
			left++
			continue
		}
		if nums[right] == val {
			right--
			continue
		}
		nums[left] = nums[right]
		left++
		right--
	}
	return left
}
