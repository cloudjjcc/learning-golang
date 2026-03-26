package nowcoder

import (
	"fmt"
	"testing"
)

func Test_rebuildBinaryTree(t *testing.T) {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	in := []int{4, 7, 2, 1, 5, 3, 8, 6}
	tree := rebuildBinaryTree(pre, in)
	fmt.Println(tree)
}
