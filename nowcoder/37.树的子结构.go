package nowcoder

import (
	"github.com/cloudjjcc/go-exercises/datastructures"
)

func hasSubTree(tree1 *datastructures.TestTreeNode, tree2 *datastructures.TestTreeNode) bool {
	if tree1 == nil || tree2 == nil {
		return false
	}
	if tree1.Value == tree2.Value {
		if judge(tree1, tree2) {
			return true
		}
	}
	return hasSubTree(tree1.Left, tree2) || hasSubTree(tree1.Right, tree2)
}

// 判断tree2是否是tree1的子结构
func judge(tree1 *datastructures.TestTreeNode, tree2 *datastructures.TestTreeNode) bool {
	if tree2 == nil {
		return true
	}
	if tree1 == nil {
		return false
	}
	if tree1.Value != tree2.Value {
		return false
	}
	return judge(tree1.Left, tree2.Left) && judge(tree1.Right, tree2.Right)
}
