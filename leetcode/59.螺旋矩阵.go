package leetcode

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	topLimit, rightLimit, bottomLimit, leftLimit := 0, n-1, n-1, 0
	v := 1
	for v <= n*n {
		// 右
		for i := leftLimit; i <= rightLimit; i++ {
			matrix[topLimit][i] = v
			v++
		}
		topLimit++
		// 下
		for i := topLimit; i <= bottomLimit; i++ {
			matrix[i][rightLimit] = v
			v++
		}
		rightLimit--
		// 左
		for i := rightLimit; i >= leftLimit; i-- {
			matrix[bottomLimit][i] = v
			v++
		}
		bottomLimit--
		// 上
		for i := bottomLimit; i >= topLimit; i-- {
			matrix[i][leftLimit] = v
			v++
		}
		leftLimit++
	}
	return matrix
}
