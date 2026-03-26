package datastructures

import "math"

type avlTreeNode struct {
	key                 int
	balance             int
	left, right, parent *avlTreeNode
}
type AVLTree struct {
	root *avlTreeNode
}

func (t *AVLTree) Add(key int) {
	node := &avlTreeNode{
		key: key,
	}
	if t.root == nil {
		t.root = node
		return
	}
	p := t.root
	for {
		if node.key == p.key {
			return
		}
		if node.key > p.key {
			if p.right == nil {
				p.right = node
				node.parent = p
				break
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = node
				node.parent = p
				break
			}
			p = p.left
		}
	}
	// 检查树的平衡性
	p = node.parent
	for p != nil {
		if node.key > p.key {
			p.balance--
		} else {
			p.balance++
		}
		if p.balance == 0 {
			return
		}
		if math.Abs(float64(p.balance)) == 2 {
			t.fix(p)
			break
		}
		p = p.parent
	}
}

// 修复不平衡树
func (t *AVLTree) fix(p *avlTreeNode) {
	if p.balance == 2 {
		if p.left.balance == 1 { //LL
			t.rightRotate(p)
		} else { //LR
			t.leftRotate(p.left)
			t.rightRotate(p)
		}
	} else {
		if p.right.balance == -1 { //RR
			t.leftRotate(p)
		} else { //RL
			t.rightRotate(p.right)
			t.leftRotate(p)
		}
	}
}
func (t *AVLTree) leftRotate(p *avlTreeNode) *avlTreeNode {
	pivot := p.right
	pivot.parent = p.parent
	if pivot.parent == nil {
		t.root = pivot
	}
	p.right = pivot.left
	pivot.left = p
	p.parent = pivot
	p.balance = 0
	pivot.balance = 0
	return pivot
}
func (t *AVLTree) rightRotate(p *avlTreeNode) *avlTreeNode {
	pivot := p.left
	p.left = pivot.right
	pivot.right = p
	pivot.parent = p.parent
	if pivot.parent == nil {
		t.root = pivot
	}
	p.parent = pivot
	p.balance = 0
	pivot.balance = 0
	return pivot
}
