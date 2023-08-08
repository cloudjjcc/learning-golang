package leetcode

// 递归
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var dsfFn func(node *TreeNode)
	dsfFn = func(node *TreeNode) {
		if node == nil {
			return
		}
		dsfFn(node.Left)
		ans = append(ans, node.Val)
		dsfFn(node.Right)
	}
	dsfFn(root)
	return ans
}

// 用栈
func inorderTraversal1(root *TreeNode) []int {
	// stack
	stack := make([]*TreeNode, 0)
	stackPush := func(val *TreeNode) {
		stack = append(stack, val)
	}
	stackPop := func() *TreeNode {
		if len(stack) == 0 {
			return nil
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}
	stackIsEmpty := func() bool {
		return len(stack) == 0
	}

	ans := make([]int, 0)

	cur := root
	for cur != nil || !stackIsEmpty() {
		if cur != nil {
			stackPush(cur)
			cur = cur.Left
		} else {
			cur = stackPop()
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return ans
}
