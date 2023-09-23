package leetcode

// 并查集
func validPath(n int, edges [][]int, source int, destination int) bool {
	// 并查集
	dj := make([]int, n)
	for i := 0; i < n; i++ {
		dj[i] = i
	}
	// 查找集合号
	var find func(v int) int
	find = func(v int) int {
		if dj[v] != v {
			dj[v] = find(dj[v])
		}
		return dj[v]
	}
	join := func(a, b int) {
		a = find(a)
		b = find(b)
		if a != b {
			dj[a] = b
		}
	}
	isSame := func(a, b int) bool {
		return find(a) == find(b)
	}
	for _, v := range edges {
		join(v[0], v[1])
	}
	return isSame(source, destination)
}

// dfs
func validPath1(n int, edges [][]int, source int, destination int) bool {
	vertexs := make([][]int, n)
	visited := make([]bool, n)
	for _, v := range edges {
		vertexs[v[0]] = append(vertexs[v[0]], v[1])
		vertexs[v[1]] = append(vertexs[v[1]], v[0])
	}
	var dfs func(source, destination int) bool
	dfs = func(source, destination int) bool {
		if source == destination {
			return true
		}
		visited[source] = true
		for _, v := range vertexs[source] {
			if visited[v] {
				continue
			}
			if dfs(v, destination) {
				return true
			}
		}
		return false
	}
	return dfs(source, destination)
}
