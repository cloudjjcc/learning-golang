package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	pop := func() int {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return top
	}
	peek := func() int {
		return stack[len(stack)-1]
	}
	empty := func() bool {
		return len(stack) == 0
	}
	push := func(v int) {
		stack = append(stack, v)
	}
	ans := make([]int, len(nums1))
	m := make(map[int]int)
	for i, v := range nums1 {
		ans[i] = -1
		m[v] = i
	}

	for i, v := range nums2 {
		for !empty() && v > nums2[peek()] {
			if idx, ok := m[nums2[pop()]]; ok {
				ans[idx] = v
			}
		}
		push(i)
	}
	return ans
}
