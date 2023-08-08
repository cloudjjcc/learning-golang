package leetcode

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 寻找值为i的数
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
