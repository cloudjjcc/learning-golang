package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
