package leetcode

func dailyTemperatures(temperatures []int) []int {
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
	ans := make([]int, len(temperatures))
	for i, v := range temperatures {
		for !empty() && v > temperatures[peek()] {
			t := pop()
			ans[t] = i - t
		}
		push(i)
	}
	return ans
}
