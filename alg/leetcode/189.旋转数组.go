package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)
	for i, v := range nums {
		tmp[(i+k)%n] = v
	}
	copy(nums, tmp)
	return
}
