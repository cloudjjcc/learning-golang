package leetcode

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
