package leetcode

func isValidSudoku(board [][]byte) bool {
	var (
		rows, cols, boxes [9]uint16
	)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			digit := board[i][j] - '0'
			mask := uint16(1 << digit)
			if rows[i]&mask != 0 {
				return false
			}
			rows[i] |= mask
			if cols[j]&mask != 0 {
				return false
			}
			cols[j] |= mask
			if boxes[i/3*3+j/3]&mask != 0 {
				return false
			}
			boxes[i/3*3+j/3] |= mask
		}
	}
	return true
}
