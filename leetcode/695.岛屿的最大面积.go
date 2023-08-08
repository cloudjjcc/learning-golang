package leetcode

//给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
//找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-area-of-island
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	// init visited matrix
	visited := make([][]bool, 0)
	for i := 0; i < rows; i++ {
		visited = append(visited, make([]bool, cols))
	}
	// maxArea
	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tmp := dfs(grid, i, j, visited)
			if tmp > maxArea {
				maxArea = tmp
			}
		}
	}
	return maxArea
}

var dir = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func dfs(grid [][]int, i int, j int, visited [][]bool) int {
	if i < 0 ||
		j < 0 ||
		len(grid) == 0 ||
		i >= len(grid) ||
		j >= len(grid[0]) ||
		visited[i][j] ||
		grid[i][j] == 0 {
		return 0
	}
	area := 1
	visited[i][j] = true
	for k := 0; k < len(dir); k++ {
		area += dfs(grid, i+dir[k][0], j+dir[k][1], visited)
	}
	return area
}
