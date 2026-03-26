package leetcode

func generateParenthesis(n int) []string {
	var ans []string
	var generateParenthesisFn func(s string, left, right int)
	generateParenthesisFn = func(s string, left, right int) {
		if left == 0 && right == 0 {
			ans = append(ans, s)
			return
		}
		if left >= right {
			generateParenthesisFn(s+"(", left-1, right)
		} else {
			if left > 0 {
				generateParenthesisFn(s+"(", left-1, right)
			}
			generateParenthesisFn(s+")", left, right-1)
		}
		return
	}
	generateParenthesisFn("", n, n)
	return ans
}
