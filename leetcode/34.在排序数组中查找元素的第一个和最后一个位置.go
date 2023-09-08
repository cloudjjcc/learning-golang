package leetcode

func searchRange(nums []int, target int) []int {

	binarySearch := func(target int, last bool) int {
		left := 0
		right := len(nums) - 1
		ans := -1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				ans = mid
				if last {
					left = mid + 1
				} else { //first
					right = mid - 1
				}
				continue
			}
			if target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return ans
	}
	return []int{binarySearch(target, false), binarySearch(target, true)}
}
