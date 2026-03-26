package leetcode

func postorderTraversal(root *TreeNode) []int {
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
	stackPeek := func() *TreeNode {
		return stack[len(stack)-1]
	}
	ans := make([]int, 0)

	cur := root
	last := (*TreeNode)(nil)
	for cur != nil || !stackIsEmpty() {
		if cur != nil {
			stackPush(cur)
			cur = cur.Left
		} else {
			cur = stackPeek()
			if cur.Right == nil || cur.Right == last {
				ans = append(ans, cur.Val)
				stackPop()
				last = cur
				cur = nil
			} else {
				cur = cur.Right
			}
		}
	}
	return ans
}
