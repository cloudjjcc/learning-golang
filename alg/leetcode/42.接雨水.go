package leetcode

// 动态规划解法
func trap1(height []int) int {
	leftMaxArr := make([]int, len(height))
	rightMaxArr := make([]int, len(height))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	leftMax, rightMax := 0, 0
	for i := 0; i < len(height); i++ {
		leftMax = max(leftMax, height[i])
		leftMaxArr[i] = leftMax
	}
	for i := len(height) - 1; i >= 0; i-- {
		rightMax = max(rightMax, height[i])
		rightMaxArr[i] = rightMax
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		ans += min(leftMaxArr[i], rightMaxArr[i]) - height[i]
	}
	return ans
}

// 双指针解法
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
