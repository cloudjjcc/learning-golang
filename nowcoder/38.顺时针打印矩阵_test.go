package nowcoder

import (
	"fmt"
	"testing"
)

func Test_printMatrix(t *testing.T) {
	testMatrix := make([][]int, 4)
	testMatrix[0] = []int{1, 2, 3, 4}
	testMatrix[1] = []int{5, 6, 7, 8}
	testMatrix[2] = []int{9, 10, 11, 12}
	testMatrix[3] = []int{13, 14, 15, 16}
	fmt.Println(printMatrix(testMatrix))
}
