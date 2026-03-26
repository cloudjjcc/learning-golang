package leetcode

import "unsafe"

// TODO
func convertN(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrx := make([][]byte, numRows)
	x, y := 0, 0
	maxY := numRows - 1
	yDirection := 1
	for i := 0; i < len(s); i++ {
		matrx[y] = append(matrx[y], s[i])
		x++
		y += yDirection
		if y == maxY || y == 0 {
			yDirection = -yDirection
		}
	}
	ans := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		ans = append(ans, matrx[i]...)
	}
	return *(*string)(unsafe.Pointer(&ans))
}
