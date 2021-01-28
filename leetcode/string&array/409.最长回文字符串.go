package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bananas"))
}

func longestPalindrome(str string) int {
	if len(str) == 0 {
		return 0
	}
	charCount := make(map[byte]int)
	for _, v := range str {
		charCount[(byte(v))]++
	}
	count := 0
	for _, v := range charCount {
		if v > 1 {
			count += (v / 2) * 2
		}
		if v == len(str) {
			return v
		}
	}
	if count == len(str) {
		return count
	}
	return count + 1
}
