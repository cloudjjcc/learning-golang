package nowcoder

//题目描述
//统计一个数字在排序数组中出现的次数。

func numOfK(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	getFirst := func(arr []int, k int) int {
		var (
			left  = 0
			right = len(arr) - 1
		)
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == k {
				if mid == 0 || arr[mid-1] != k {
					return mid
				}
				right = mid - 1
			} else if arr[mid] < k {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}
	getLast := func(arr []int, k int) int {
		var (
			left  = 0
			right = len(arr) - 1
		)
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == k {
				if mid == right || arr[mid+1] != k {
					return mid
				}
				left = mid + 1
			} else if arr[mid] < k {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}
	first := getFirst(arr, k)
	last := getLast(arr, k)
	return last - first + 1
}
