package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}

func gcdOfStrings(s1 string, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 || s1+s2 != s2+s1 {
		return ""
	}
	return s1[:gcd(len(s1), len(s2))]
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
