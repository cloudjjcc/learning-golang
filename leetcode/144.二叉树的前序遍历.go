package leetcode

func preorderTraversal(root *TreeNode) []int {
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
			ans = append(ans, cur.Val)
			cur = cur.Left
		} else {
			cur = stackPop()
			cur = cur.Right
		}
	}
	return ans
}
