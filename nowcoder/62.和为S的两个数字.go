package main

import "fmt"

func main() {
	fmt.Println(findNumsWithSum([]int{1, 2, 3, 4, 5, 6, 7}, 10))
}

func findNumsWithSum(arr []int, sum int) []int {
	if len(arr) <= 1 {
		return []int{}
	}
	var (
		left  = 0
		right = len(arr) - 1
	)
	for left < right {
		tmp := arr[left] + arr[right]
		if tmp < sum {
			left++
		} else if tmp > sum {
			right--
		} else {
			return []int{arr[left], arr[right]}
		}
	}
	return []int{}
}
