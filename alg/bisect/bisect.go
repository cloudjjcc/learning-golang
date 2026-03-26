package bisect

// LeftBound 二分算法查找左边界，如果target 未找到，返回大于target的最小索引
func LeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else { // nums[mid]=target
			right = mid - 1
		}
	}
	return left
}

// RightBound 二分算法查找右侧边界，如果target未找到，返回小于target的最大索引
func RightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else { // nums[mid]==target
			left = mid + 1
		}
	}
	return right
}
