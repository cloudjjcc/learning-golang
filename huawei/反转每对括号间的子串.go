package huawei

import (
	"fmt"
	"unsafe"
)

func mainReverseParentheses() {
	var s string
	if _, err := fmt.Scanln(&s); err != nil {
		return
	}
	s = reverseParentheses(s)
	fmt.Println(s)
}
func reverseParentheses(s string) string {
	if s == "" {
		return s
	}
	stack := make([]int, 0)
	pair := make(map[int]int) //括号匹配位置
	push := func(v int) {
		stack = append(stack, v)
	}
	pop := func() int {
		if len(stack) == 0 {
			return -1
		}
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return t
	}
	ansLen := 0
	for i, v := range s {
		switch v {
		case '(':
			push(i)
		case ')':
			start := pop()
			pair[start], pair[i] = i, start
		default:
			ansLen++
		}
	}
	ans := make([]byte, 0, ansLen)
	for i, step := 0, 1; i < len(s); i += step {
		switch v := s[i]; v {
		case '(', ')':
			i = pair[i]
			step = -step
		default:
			ans = append(ans, v)
		}
	}
	return *(*string)(unsafe.Pointer(&ans))
}
