package leetcode

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	right, down, left, up := n-1, m-1, 0, 1
	i, j := 0, 0
	direction := 0
	step := func() bool {
		switch direction {
		case 0:
			if j+1 > right {
				return false
			}
			j++
			if j == right {
				right--
				direction = 1
			}
		case 1:
			if i+1 > down {
				return false
			}
			i++
			if i == down {
				down--
				direction = 2
			}
		case 2:
			if j-1 < left {
				return false
			}
			j--
			if j == left {
				left++
				direction = 3
			}
		case 3:
			if i-1 < up {
				return false
			}
			i--
			if i == up {
				up++
				direction = 0
			}
		}
		return true
	}
	ans := make([]int, 0, m*n)
	ans = append(ans, matrix[0][0])
	if j == right {
		right--
		direction = 1
	}
	for step() {
		ans = append(ans, matrix[i][j])
	}
	return ans
}
