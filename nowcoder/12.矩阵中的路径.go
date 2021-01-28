package main

import "fmt"

//题目描述
//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
//路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左，向右，向上，向下移动一个格子。
//如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
//例如 a b c e s f c s a d e e 矩阵中包含一条字符串"bcced"的路径，
//但是矩阵中不包含"abcb"路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入该格子
// 四个方向
var (
	dir = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func main() {
	matrix := [][]rune{{'a', 'b', 't', 'g'}, {'c', 'f', 'c', 's'}, {'j', 'd', 'e', 'h'}}
	testPath := "bfcedh"
	fmt.Println(hasPath(matrix, testPath))
}

// DFS (深度优先搜索)回溯法
func hasPath(matrix [][]rune, path string) bool {
	var visited [][]bool
	for range matrix {
		visited = append(visited, make([]bool, len(matrix[0])))
	}
	for r, row := range matrix {
		for c := range row {
			if backtracking(matrix, visited, path, 0, r, c) {
				return true
			}
		}
	}
	return false
}

func backtracking(matrix [][]rune, visited [][]bool, path string, l int, r int, c int) bool {
	if len(path) == l {
		return true
	}
	if r >= len(matrix) ||
		r < 0 ||
		c >= len(matrix[0]) ||
		c < 0 ||
		l >= len(path) ||
		matrix[r][c] != rune(path[l]) ||
		visited[r][c] {
		return false
	}
	// 设置标记
	visited[r][c] = true
	fmt.Printf("%c\n", matrix[r][c])
	// 四个方向寻找
	for _, v := range dir {
		if backtracking(matrix, visited, path, l+1, r+v[0], c+v[1]) {
			return true
		}
	}
	visited[r][c] = false
	return false
}
