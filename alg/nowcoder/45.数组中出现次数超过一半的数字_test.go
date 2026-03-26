package nowcoder

import (
	"fmt"
	"testing"
)

func Test_moreThanHalfNum(t *testing.T) {
	testArr := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(moreThanHalfNum(testArr))
}
