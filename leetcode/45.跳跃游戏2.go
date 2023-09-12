package leetcode

func jump(nums []int) int {
	n := len(nums)
	rightMost := nums[0]
	end := 0
	steps := 0
	for i := 0; i < n; i++ {
		if t := nums[i] + i; t > rightMost {
			rightMost = t
		}
		if i == end {
			end = rightMost
			steps++
		}
	}
	return steps
}
