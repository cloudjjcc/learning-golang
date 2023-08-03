package leetcode

func isValid(s string) bool {
	stack := make([]byte, 0)
	m := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	for i := 0; i < len(s); i++ {
		var ch byte
		if len(stack) > 0 {
			ch = stack[len(stack)-1]
		}
		if v, ok := m[s[i]]; ok { //右括号
			if ch == v { //括号匹配
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else { //左括号入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
