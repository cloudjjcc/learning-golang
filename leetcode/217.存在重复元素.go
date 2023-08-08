package leetcode

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}
