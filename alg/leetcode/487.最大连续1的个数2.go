package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	left, right := 0, 0
	ans := 1
	lastZeroIdx := -1
	for right < len(nums) {
		if nums[right] == 0 {
			if lastZeroIdx != -1 {
				left = lastZeroIdx + 1
			}
			lastZeroIdx = right
		}
		if ll := right - left + 1; ll > ans {
			ans = ll
		}
		right++
	}
	return ans
}
