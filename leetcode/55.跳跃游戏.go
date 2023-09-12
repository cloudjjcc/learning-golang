package leetcode

func canJump(nums []int) bool {
	jump := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		jump[i] = i + nums[i]
	}
	maxJump := jump[0]
	i := 0
	for ; i < len(nums) && i <= maxJump; i++ {
		if jump[i] > maxJump {
			maxJump = jump[i]
		}
	}
	if i == len(nums)-1 {
		return true
	}
	return false
}
