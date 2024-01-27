package leetcode

// 1. 有⼀个已经排好序的整数序列（升序，⽆重复项），序列中可能有正整数、负整数或者0，
// 请 ⽤你认为最优的⽅法求序列中绝对值最⼩的数。
// **_要求不能使⽤顺序⽐较的⽅法（时间复杂 度需要⼩于O(n)），不能使⽤内置查找函数，_**可以⽤任何语⾔实现。
// 输⼊：⼀个有序的
// 整数序列。 输出：绝对值最⼩的数。
func findMinAbs(arr []int) int {
	n := len(arr)
	// 所有数为非负数
	if arr[0] >= 0 {
		return arr[0]
	}
	// 所有数为非正数
	if arr[n-1] <= 0 {
		return arr[n-1]
	}
	// 正负皆有,二分查找
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < 0 {
			left = mid + 1
		} else if arr[mid] > 0 {
			if arr[mid-1]*arr[mid] < 0 { //正负交界
				if -arr[mid-1] > arr[mid] {
					return arr[mid]
				} else {
					return arr[mid-1]
				}
			}
			right = mid - 1
		} else {
			return arr[mid]
		}
	}
	return left
}
