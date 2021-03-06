package nowcoder

import (
	"fmt"
	"testing"
)

func Test_levelOrderTraversal(t *testing.T) {
	testTree := &tnode{Value: 0}
	tmpNode := testTree
	for i := 1; i < 10; i++ {
		if i&1 == 0 {
			if tmpNode.right != nil {
				tmpNode = tmpNode.right
			} else {
				tmpNode.right = &tnode{Value: i}
			}
		} else {
			if tmpNode.left != nil {
				tmpNode = tmpNode.left
			} else {
				tmpNode.left = &tnode{Value: i}
			}
		}
	}
	levelOrderTraversal(testTree, func(t *tnode) {
		fmt.Println(t.Value)
	})
}
