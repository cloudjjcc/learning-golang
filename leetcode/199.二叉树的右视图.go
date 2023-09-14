package leetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	push := func(n *TreeNode) {
		queue = append(queue, n)
	}
	pop := func() *TreeNode {
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
			n := pop()
			if i == s-1 {
				ans = append(ans, n.Val)
			}
			if n.Left != nil {
				push(n.Left)
			}
			if n.Right != nil {
				push(n.Right)
			}
		}
	}
	return ans
}
