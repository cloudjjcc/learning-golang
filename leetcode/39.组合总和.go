package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var dfs func(start int)
	var path []int
	pathSum := 0
	ans := make([][]int, 0)
	dfs = func(start int) {
		if pathSum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if pathSum+candidates[i] > target {
				continue
			}
			pathSum += candidates[i]
			path = append(path, candidates[i])
			dfs(i)
			pathSum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
