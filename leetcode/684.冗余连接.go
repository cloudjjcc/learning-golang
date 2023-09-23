package leetcode

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	// 并查集
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	var find func(t int) int
	find = func(t int) int {
		if parents[t] != t {
			parents[t] = find(parents[t])
		}
		return parents[t]
	}
	isSame := func(a, b int) bool {
		return find(a) == find(b)
	}
	join := func(a, b int) {
		a = find(a)
		b = find(b)
		if a != b {
			parents[a] = b
		}
	}
	for _, v := range edges {
		if isSame(v[0], v[1]) {
			return v
		}
		join(v[0], v[1])
	}
	return nil
}
