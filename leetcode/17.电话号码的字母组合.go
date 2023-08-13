package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var ans []string
	charsByDigit := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfsFn func(i int, p string)
	dfsFn = func(i int, p string) {
		for _, ch := range charsByDigit[digits[i]] {
			p += string(ch)
			if i == len(digits)-1 {
				ans = append(ans, p)
			} else {
				dfsFn(i+1, p)
			}
			p = p[:len(p)-1]
		}
	}
	dfsFn(0, "")
	return ans
}
