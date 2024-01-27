package leetcode

func countTarget(scores []int, target int) int {
	if len(scores) == 0 {
		return 0
	}
	// 返回等于target的左边界
	searchLeft := func(target int) int {
		left, right := 0, len(scores)-1
		ans := -1
		for left <= right {
			mid := (left + right) / 2
			if target < scores[mid] {
				right = mid - 1
			} else if target > scores[mid] {
				left = mid + 1
			} else {
				right = mid - 1
				ans = mid
			}
		}
		return ans
	}
	searchRight := func(target int) int {
		left, right := 0, len(scores)-1
		ans := -1
		for left <= right {
			mid := (left + right) / 2
			if target < scores[mid] {
				right = mid - 1
			} else if target > scores[mid] {
				left = mid + 1
			} else {
				left = mid + 1
				ans = mid
			}
		}
		return ans
	}
	left := searchLeft(target)
	right := searchRight(target)
	if left != -1 && right != -1 {
		return right - left + 1
	}
	return 0
}
