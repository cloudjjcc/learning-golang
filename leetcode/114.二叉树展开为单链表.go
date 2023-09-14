package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	helper := func(t *TreeNode) *TreeNode {
		stack := make([]*TreeNode, 0)
		push := func(n *TreeNode) {
			stack = append(stack, n)
		}
		pop := func() *TreeNode {
			if len(stack) == 0 {
				return nil
			}
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			return t
		}
		empty := func() bool {
			return len(stack) == 0
		}
		cur := root
		nodes := make([]*TreeNode, 0)
		for !empty() || cur != nil {
			for cur != nil {
				push(cur)
				nodes = append(nodes, cur)
				cur = cur.Left
			}
			node := pop()
			cur = node.Right
		}
		dummy := new(TreeNode)
		cur = dummy
		for _, v := range nodes {
			cur.Right = v
			cur = cur.Right
			cur.Left = nil
			cur.Right = nil
		}
		return nodes[len(nodes)-1]
	}
	helper(root)
	return
}
