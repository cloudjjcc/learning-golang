package leetcode

// 暴力 时间复杂度O(n)3 空间复杂度O(1)
func longestPalindrome(s string) string {
	max := 1
	maxi, maxj := 0, 0
	isValidFn := func(a, b int) bool {
		for a < b {
			if s[a] != s[b] {
				return false
			}
			a++
			b--
		}
		return true
	}
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if (j-i+1) > max && isValidFn(i, j) {
				max = j - i + 1
				maxi = i
				maxj = j
			}
		}
	}
	return s[maxi : maxj+1]
}

// 动态规划 时间复杂度O(n)2 空间复杂度O(n)2
func longestPalindrome2(s string) string {
	d := make([][]bool, len(s))
	for i := range d {
		d[i] = make([]bool, len(s))
	}
	maxLen := 1
	maxi := 0
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := l + i - 1
			if j > len(s)-1 {
				break
			}
			if s[i] == s[j] {
				if l <= 3 {
					d[i][j] = true
				} else {
					d[i][j] = d[i+1][j-1]
				}
			} else {
				d[i][j] = false
			}
			if d[i][j] && l > maxLen {
				maxLen = l
				maxi = i
			}
		}
	}
	return s[maxi : maxi+maxLen]
}
