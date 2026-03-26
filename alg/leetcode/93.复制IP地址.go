package leetcode

import "strconv"

func restoreIpAddresses(s string) []string {
	var ans []string
	isValidPart := func(p string) bool {
		if len(p) > 3 || len(p) < 1 {
			return false
		}
		if len(p) > 1 && p[0] == '0' {
			return false
		}
		atoi, err := strconv.Atoi(p)
		if err != nil {
			return false
		}
		if atoi > 255 || atoi < 0 {
			return false
		}
		return true
	}
	var backtrace func(prefix string, part int, left string)
	backtrace = func(prefix string, part int, left string) {
		if part > 4 || len(left) < 1 {
			return
		}
		if part == 4 {
			if isValidPart(left) {
				ans = append(ans, prefix+left)
			}
			return
		}
		ll := 3
		if len(left) < 3 {
			ll = len(left)
		}
		for i := 1; i <= ll; i++ {
			if isValidPart(left[:i]) {
				backtrace(prefix+left[:i]+".", part+1, left[i:])
			}
		}
	}
	backtrace("", 1, s)
	return ans
}
