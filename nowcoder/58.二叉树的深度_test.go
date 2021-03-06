package nowcoder

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"testing"
)

func Test_getTreeDepth(t *testing.T) {
	tree := datastructures.BuildTreeFromArray([]interface{}{1, 2, 3, 4, 5, 6})
	fmt.Println(getTreeDepth(tree))
}
