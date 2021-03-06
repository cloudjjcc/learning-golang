package nowcoder

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
