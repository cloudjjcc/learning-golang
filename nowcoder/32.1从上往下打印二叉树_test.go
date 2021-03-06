package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
	"testing"
)

func Test_printBinaryTree(t *testing.T) {
	tree := datastructures.BuildTreeFromArray([]interface{}{1, 2, 3, 4, 5, 6, 7})
	printBinaryTree(tree)
}
