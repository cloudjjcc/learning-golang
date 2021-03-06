package nowcoder

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"testing"
)

func Test_findNextNode(t *testing.T) {
	// 构造二叉树
	tree := datastructures.BuildTreeFromArray([]interface{}{5, 4, nil, 3, nil, 2})
	nextNode := findNextNode(tree)
	fmt.Printf("%v\n", nextNode)
}
