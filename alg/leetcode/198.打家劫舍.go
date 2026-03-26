package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	pre2 := nums[0]
	pre1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		cur := max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}
