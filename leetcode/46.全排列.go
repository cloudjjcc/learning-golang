package leetcode

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, len(nums))
	var backtracking func()
	isValid := func(num int) bool {
		for _, v := range path {
			if v == num {
				return false
			}
		}
		return true
	}
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !isValid(nums[i]) {
				continue
			}
			path = append(path, nums[i])
			backtracking()
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return ans
}
