package huawei

import "fmt"

func alive(arr []int) int {
	stack := make([]int, 0)
	push := func(v int) {
		stack = append(stack, v)
	}
	pop := func() int {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return t
	}
	peek := func() int {
		return stack[len(stack)-1]
	}
	empty := func() bool {
		return len(stack) == 0
	}
	for i := range arr {
		fmt.Println(arr)
		if arr[i] > 0 {
			push(i)
			continue
		}
		for !empty() {
			pre := peek()
			if arr[pre] < 0 {
				break
			}
			if arr[pre] < (-arr[i]) {
				pop()
				arr[i] += arr[pre]
				push(i)
			} else if arr[pre] == -arr[i] {
				pop()
				break
			} else {
				arr[pre] += arr[i]
				break
			}
		}
	}
	return len(stack)
}
