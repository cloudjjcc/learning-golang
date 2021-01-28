package main

import "fmt"

//在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。
//每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。
//请你返回最终形体的表面积。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/surface-area-of-3d-shapes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(surfaceArea(grid))
}

func surfaceArea(grid [][]int) int {
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	N := len(grid)
	ans := 0

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if grid[r][c] > 0 {
				ans += 2
				for k := 0; k < 4; k++ {
					nr := r + dr[k]
					nc := c + dc[k]
					nv := 0
					if 0 <= nr && nr < N && 0 <= nc && nc < N {
						nv = grid[nr][nc]

					}
					if grid[r][c] > nv {
						ans += grid[r][c] - nv
					}
				}
			}
		}
	}
	return ans
}
