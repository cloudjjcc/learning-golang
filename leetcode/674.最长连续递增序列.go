package leetcode

func findLengthOfLCIS(nums []int) int {
	pre := 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			pre = max(pre+1, 1)
		} else {
			pre = 1
		}
		ans = max(ans, pre)
	}
	return ans
}
