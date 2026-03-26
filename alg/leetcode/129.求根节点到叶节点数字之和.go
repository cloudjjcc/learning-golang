package leetcode

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type node struct {
		*TreeNode
		p int
	}
	queue := make([]*node, 0)
	push := func(n *node) {
		queue = append(queue, n)
	}
	pop := func() *node {
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
		p:        root.Val,
	})
	sum := 0
	for !empty() {
		s := size()
		for i := 0; i < s; i++ {
			n := pop()
			if n.Left != nil {
				push(&node{
					TreeNode: n.Left,
					p:        n.p*10 + n.Left.Val,
				})
			}
			if n.Right != nil {
				push(&node{
					TreeNode: n.Right,
					p:        n.p*10 + n.Right.Val,
				})
			}
			if n.Left == nil && n.Right == nil {
				sum += n.p
			}
		}
	}
	return sum
}
