package leetcode

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	copyPath := func() []int {
		t := make([]int, len(path))
		copy(t, path)
		return t
	}
	var backtracking func(i int)
	backtracking = func(ii int) {
		ans = append(ans, copyPath())
		if ii >= len(nums) {
			return
		}
		for i := ii; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}
