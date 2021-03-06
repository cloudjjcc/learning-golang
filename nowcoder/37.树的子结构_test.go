package nowcoder

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"testing"
)

func Test_hasSubTree(t *testing.T) {
	tree1 := datastructures.BuildTreeFromArray([]interface{}{0, 1, 2, 3, 4, 5, 6})
	tree2 := datastructures.BuildTreeFromArray([]interface{}{2, 5, 6})
	fmt.Println(hasSubTree(tree1, tree2))
}
