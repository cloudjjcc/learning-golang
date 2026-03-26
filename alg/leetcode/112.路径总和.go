package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	type node struct {
		*TreeNode
		pathSum int
	}
	queue := make([]*node, 0)
	push := func(n *node) {
		queue = append(queue, n)
	}
	pop := func() *node {
		if len(queue) == 0 {
			return nil
		}
		top := queue[0]
		queue = queue[1:]
		return top
	}
	empty := func() bool {
		return len(queue) == 0
	}
	size := func() int {
		return len(queue)
	}
	push(&node{
		TreeNode: root,
		pathSum:  root.Val,
	})
	for !empty() {
		s := size()
		for i := 0; i < s; i++ {
			n := pop()
			if n.Left != nil {
				push(&node{
					TreeNode: n.Left,
					pathSum:  n.Left.Val + n.pathSum,
				})
			}
			if n.Right != nil {
				push(&node{
					TreeNode: n.Right,
					pathSum:  n.Right.Val + n.pathSum,
				})
			}
			if n.Left == nil && n.Right == nil {
				if n.pathSum == targetSum {
					return true
				}
			}
		}
	}
	return false
}
