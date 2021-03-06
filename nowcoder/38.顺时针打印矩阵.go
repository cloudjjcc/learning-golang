package nowcoder

//题目描述
//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，
//例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
//则依次打印出数字1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.
func printMatrix(matrix [][]int) []int {
	buf := make([]int, 0)
	if len(matrix) == 0 {
		return buf
	}
	var (
		rows                  = len(matrix)
		cols                  = len(matrix[0])
		up, down, left, right = 0, rows - 1, 0, cols - 1
		r, c                  = 0, 0
	)
	for {
		for c = left; c <= right; c++ {
			buf = append(buf, matrix[up][c])
		}
		// 向下逼近
		up++
		if up > down {
			break
		}
		for r = up; r <= down; r++ {
			buf = append(buf, matrix[r][right])
		}
		// 向左逼近
		right--
		if left > right {
			break
		}
		for c = right; c >= left; c-- {
			buf = append(buf, matrix[down][c])
		}
		// 向上逼近
		down--
		if up > down {
			break
		}
		for r = down; r >= up; r-- {
			buf = append(buf, matrix[r][left])
		}
		// 向右逼近
		left++
		if left > right {
			break
		}
	}
	return buf
}
