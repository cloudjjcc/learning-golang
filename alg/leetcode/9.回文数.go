package leetcode

func isPalindrome2(x int) bool {
	num := x
	rx := 0
	for num > 0 {
		rx = rx*10 + num%10
		num /= 10
	}
	return rx == x
}
