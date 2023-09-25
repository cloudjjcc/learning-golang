package leetcode

func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	push := func(n *TreeNode) {
		queue = append(queue, n)
	}
	pop := func() *TreeNode {
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
	push(root)
	ans := make([]int, 0)
	for !empty() {
		s := size()
		for i := 0; i < s; i++ {
			node := pop()
			if node.Left != nil {
				push(node.Left)
			}
			if node.Right != nil {
				push(node.Right)
			}
			if i == s-1 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}
