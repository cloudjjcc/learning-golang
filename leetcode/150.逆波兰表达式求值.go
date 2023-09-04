package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))
	stackPush := func(val string) {
		stack = append(stack, val)
	}
	stackPop := func() string {
		if len(stack) == 0 {
			return ""
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return tmp
	}
	exeOp := func(exe func(a, b int) int) {
		b, _ := strconv.Atoi(stackPop())
		a, _ := strconv.Atoi(stackPop())
		stackPush(strconv.Itoa(exe(a, b)))
		return
	}
	add := func(a, b int) int { return a + b }
	sub := func(a, b int) int { return a - b }
	mul := func(a, b int) int { return a * b }
	div := func(a, b int) int { return a / b }

	for _, v := range tokens {
		switch v {
		case "+":
			exeOp(add)
		case "-":
			exeOp(sub)
		case "*":
			exeOp(mul)
		case "/":
			exeOp(div)
		default:
			stackPush(v)
		}
	}
	ans := stackPop()
	if ans == "" {
		return 0
	}
	ansv, _ := strconv.Atoi(ans)
	return ansv
}
