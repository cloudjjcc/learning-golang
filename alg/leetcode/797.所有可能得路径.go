package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)
	dfs = func(i int) {
		path = append(path, i)
		defer func() {
			path = path[:len(path)-1]
		}()
		if i == len(graph)-1 {
			t := make([]int, len(path))
			copy(t, path)
			ans = append(ans, t)
			return
		}
		for _, v := range graph[i] {
			dfs(v)
		}
	}
	dfs(0)
	return ans
}
