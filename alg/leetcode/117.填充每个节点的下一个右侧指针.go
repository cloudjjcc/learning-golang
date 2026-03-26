package leetcode

/**
* Definition for a Node.
*
 */
type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root == nil {
		return nil
	}
	queue := make([]*Node1, 0)
	push := func(node *Node1) {
		queue = append(queue, node)
	}
	pop := func() *Node1 {
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
	for !empty() {
		s := size()
		var pre *Node1
		for i := 0; i < s; i++ {
			t := pop()
			if pre != nil {
				pre.Next = t
			}
			pre = t
			if t.Left != nil {
				push(t.Left)
			}
			if t.Right != nil {
				push(t.Right)
			}
		}
	}
	return root
}
