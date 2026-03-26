package leetcode

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type nodeWithNumber struct {
		node *TreeNode
		num  int
	}
	queue := make([]*nodeWithNumber, 0)
	push := func(n *nodeWithNumber) {
		queue = append(queue, n)
	}
	pop := func() *nodeWithNumber {
		if len(queue) == 0 {
			return nil
		}
		top := queue[0]
		queue = queue[1:]
		return top
	}
	size := func() int {
		return len(queue)
	}
	push(&nodeWithNumber{
		node: root,
		num:  0,
	})
	max := 1
	for size() > 0 {
		s := size()
		tmax := queue[len(queue)-1].num - queue[0].num + 1
		if tmax > max {
			max = tmax
		}
		for i := 0; i < s; i++ {
			node := pop()
			if node.node.Left != nil {
				push(&nodeWithNumber{
					node: node.node.Left,
					num:  2*node.num + 1,
				})
			}
			if node.node.Right != nil {
				push(&nodeWithNumber{
					node: node.node.Right,
					num:  2*node.num + 2,
				})
			}
		}
	}
	return max
}
