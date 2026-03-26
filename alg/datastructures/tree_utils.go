package datastructures

import "fmt"

type TestTreeNode struct {
	Left, Right, Parent *TestTreeNode
	Value               int
}

// 根据前序遍历和中序遍历的结构构建二叉树
func BuildTreeFromOrder(preorder, inorder []int) *TestTreeNode {
	if len(preorder) < 1 || len(inorder) < 1 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TestTreeNode{}
	root.Value = preorder[0]
	for i, v := range inorder {
		if v == root.Value {
			root.Left = BuildTreeFromOrder(preorder[1:i+1], inorder[:i])
			root.Right = BuildTreeFromOrder(preorder[i+1:], inorder[i+1:])
			if root.Left != nil {
				root.Left.Parent = root
			}
			if root.Right != nil {
				root.Right.Parent = root
			}
			return root
		}
	}
	return root
}

// 从数组构建二叉树
func BuildTreeFromArray(arr []interface{}) *TestTreeNode {
	if len(arr) == 0 {
		return nil
	}
	tmp := make(map[int]*TestTreeNode)
	for i := len(arr) - 1; i >= 0; i-- {
		if v, ok := arr[i].(int); ok {
			node := &TestTreeNode{Value: v}
			if vv, ok := tmp[2*i+1]; ok {
				node.Left = vv
				vv.Parent = node
			}
			if vv, ok := tmp[2*i+2]; ok {
				node.Right = vv
				vv.Parent = node
			}
			tmp[i] = node
		}
	}
	return tmp[0]
}

type EleFunc func(node *TestTreeNode)
type LevelEleFunc func(int, int)

// 前序遍历
func (t *TestTreeNode) PreOrder(eleFunc EleFunc) {
	if t == nil {
		return
	}
	eleFunc(t)
	t.Left.PreOrder(eleFunc)
	t.Right.PreOrder(eleFunc)
}

// 中序遍历
func (t *TestTreeNode) InOrder(eleFunc EleFunc) {
	if t == nil {
		return
	}
	t.Left.InOrder(eleFunc)
	eleFunc(t)
	t.Right.InOrder(eleFunc)
}

// 中序非递归
func (t *TestTreeNode) InOrderNotRecursive(eleFunc EleFunc) {
	if t == nil {
		return
	}
	stack := Stack{}
	cur := t
	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		cur = stack.Pop().(*TestTreeNode)
		eleFunc(cur)
		cur = cur.Right
	}
}

// 后序遍历
func (t *TestTreeNode) PostOrder(eleFunc EleFunc) {
	if t == nil {
		return
	}
	t.Left.PostOrder(eleFunc)
	t.Right.PostOrder(eleFunc)
	eleFunc(t)
}

// 前序非递归
func (t *TestTreeNode) PreOrderNotRecursive(eleFunc EleFunc) {
	if t == nil {
		return
	}
	stack := Stack{}
	cur := t
	for cur != nil || !stack.Empty() {
		for cur != nil {
			stack.Push(cur)
			eleFunc(cur)
			cur = cur.Left
		}
		cur = stack.Pop().(*TestTreeNode)
		cur = cur.Right
	}
}

// 层序遍历
func (t *TestTreeNode) LevelOrder(eleFunc LevelEleFunc) {
	if t == nil {
		return
	}
	queue := Queue{}
	queue.Enqueue(t)
	var (
		plast = t //
		last  = t
		level = 0 //当前层
	)
	for !queue.Empty() {
		root := queue.Dequeue().(*TestTreeNode)
		eleFunc(root.Value, level)
		// left child
		if root.Left != nil {
			queue.Enqueue(root.Left)
			last = root.Left
		}
		// right child
		if root.Right != nil {
			queue.Enqueue(root.Right)
			last = root.Right
		}
		// level finish
		if root == plast {
			level++
			plast = last
		}
	}
}

func (t *TestTreeNode) String() string {
	buf := make([]byte, 0)
	curLevel := 0
	t.LevelOrder(func(ele int, level int) {
		// new level
		if level != curLevel {
			curLevel++
			buf = append(buf, '\n')
		}
		// add level element
		buf = append(buf, fmt.Sprintf("%d ", ele)...)
	})
	return string(buf)
}
