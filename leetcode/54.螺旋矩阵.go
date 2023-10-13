package leetcode

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	right, down, left, up := n-1, m-1, 0, 0
	ans := make([]int, 0, n*m)
	for {
		// 右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		// 下
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 左
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		// 上
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}
