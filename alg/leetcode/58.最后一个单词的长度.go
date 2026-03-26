package leetcode

func lengthOfLastWord(s string) int {
	ll := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if ll == 0 {
				continue
			}
			return ll
		}
		ll++
	}
	return ll
}
