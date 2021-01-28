package main

import "fmt"

//在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。
//请找出数组中任意一个重复的数字。
//Input:
//{2, 3, 1, 0, 2, 5}
//Output:
//2
func main() {
	testNums := []int{4, 1, 1, 1, 2, 5}
	PrintRepeatNum(testNums)
	PrintRepeatNum2(testNums)
}

// 利用数组中的数只能在0~n-1范围解决问题
// 时间复杂度O(n) 空间复杂度O(1)
func PrintRepeatNum(nums []int) {
	for i := range nums {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				fmt.Println(nums[i])
				return
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

}

// 利用map
// 时间复杂度O(n) 空间复杂度O(n)
func PrintRepeatNum2(nums []int) {
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			fmt.Println(v)
			return
		}
		numsMap[v] = true
	}
}
