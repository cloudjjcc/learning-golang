package leetcode

// 时间复杂度 O(nlogn)
func twoSum22(numbers []int, target int) []int {
	findTarget := func(i int, target int) int {
		left, right := i, len(numbers)-1
		for left <= right {
			mid := left + (right-left)/2
			if numbers[mid] == target {
				return mid
			} else if numbers[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}
	for i := 0; i < len(numbers); i++ {
		j := findTarget(i+1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return []int{-1, -1}
}

func twoSum20(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
