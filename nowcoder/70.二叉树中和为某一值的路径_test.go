package nowcoder

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"testing"
)

func Test_findPath(t *testing.T) {
	preorder := []int{0, 1, 3, 5, 2, 4, 6}
	inorder := []int{3, 1, 5, 0, 4, 2, 6}
	testTree := datastructures.BuildTreeFromOrder(preorder, inorder)
	testTree.LevelOrder(func(i, j int) {
		fmt.Println(i)
	})
	fmt.Println(findPath(testTree, 6))
}
