package datastructures

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	Value  interface{}
	tree   *BinaryTree
}
type ITree interface {
	Root() *TreeNode
	Find(value interface{}) *TreeNode
	Delete(node *TreeNode) bool
	Insert(value interface{}) *TreeNode
}
type BinaryTree struct {
	root *TreeNode
	Comparator
}

func NewWith(comparator Comparator) *BinaryTree {
	return &BinaryTree{
		root:       nil,
		Comparator: comparator,
	}
}
func (b *BinaryTree) Root() *TreeNode {
	return b.root
}

func (b *BinaryTree) Find(value interface{}) *TreeNode {
	slot := b.root
	for slot != nil {
		comp := b.Comparator(slot.Value, value)
		if comp == 0 {
			return slot
		}
		if comp < 0 {
			slot = slot.right
		} else {
			slot = slot.left
		}
	}
	return nil
}
func (b *BinaryTree) Delete(node *TreeNode) bool {
	if node == nil || node.tree != b {
		return false
	}
	var child *TreeNode
	if node.left == nil && node.right == nil {
		child = nil
	} else if node.left == nil {
		child = node.left
	} else if node.right == nil {
		child = node.right
	} else { // if node has two child ,find max node
		child = findMax(node)
		if !b.Delete(child) {
			return false
		}
	}
	if node.parent.left == node {
		node.parent.left = child
	} else {
		node.parent.right = child
	}
	return true
}
func findMin(tree *TreeNode) *TreeNode {
	for tree != nil {
		if tree.left != nil {
			tree = tree.left
		} else {
			return tree
		}
	}
	return nil
}
func findMax(tree *TreeNode) *TreeNode {
	for tree != nil {
		if tree.right != nil {
			tree = tree.right
		} else {
			return tree
		}
	}
	return nil
}
func (b *BinaryTree) Insert(value interface{}) *TreeNode {
	node := &TreeNode{Value: value, tree: b}
	slot := b.root
	for slot != nil {
		if b.Comparator(slot.Value, value) > 0 {
			if slot.left == nil {
				slot.left = node
				node.parent = slot.left
				break
			}
			slot = slot.left
		} else {
			if slot.right == nil {
				slot.right = node
				node.parent = slot.right
				break
			}
			slot = slot.right
		}
	}
	return node
}
