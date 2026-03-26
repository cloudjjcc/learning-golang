package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	empty := func() bool {
		return len(queue) == 0
	}
	back := func() int {
		return queue[len(queue)-1]
	}
	push := func(val int) {
		for !empty() && val > back() {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, val)
	}

	front := func() int {
		return queue[0]
	}
	pop := func(val int) {
		if !empty() && val == front() {
			queue = queue[1:]
		}
	}
	for i := 0; i < k; i++ {
		push(nums[i])
	}
	ans := make([]int, 0)
	ans = append(ans, front())
	for i := k; i < len(nums); i++ {
		pop(nums[i-k])
		push(nums[i])
		ans = append(ans, front())
	}
	return ans
}
