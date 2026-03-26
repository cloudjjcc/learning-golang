package leetcode

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	ans := make([]int, len(nums))
	ansIdx := right
	for left <= right {
		leftSqu := nums[left] * nums[left]
		rightSqu := nums[right] * nums[right]
		if leftSqu > rightSqu {
			ans[ansIdx] = leftSqu
			ansIdx--
			left++
			continue
		}
		ans[ansIdx] = rightSqu
		ansIdx--
		right--
	}
	return ans
}
