package leetcode

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//中序遍历
	stack := make([]*TreeNode, 0)
	push := func(n *TreeNode) {
		stack = append(stack, n)
	}
	pop := func() *TreeNode {
		if len(stack) == 0 {
			return nil
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}
	empty := func() bool {
		return len(stack) == 0
	}
	cur := root
	dummyHead := new(TreeNode)
	pre := dummyHead
	for !empty() || cur != nil {
		for cur != nil {
			push(cur)
			cur = cur.Left
		}
		cur = pop()
		pre.Right = cur
		cur.Left = pre
		pre = cur
		cur = cur.Right
	}
	//
	head := dummyHead.Right
	head.Left = pre
	pre.Right = head
	return head
}
