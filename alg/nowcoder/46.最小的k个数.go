package nowcoder

func getLeastNums(arr []int, k int) []int {
	if len(arr) < k {
		return make([]int, 0)
	}
	res := arr[:k]
	// build max heap
	for i := k/2 - 1; i >= 0; i-- {
		down(res, i)
	}
	// adjust heap
	for i := k; i < len(arr); i++ {
		if res[0] > arr[i] {
			res[0] = arr[i]
			down(res, 0)
		}
	}
	return res
}

// 下滤
func down(arr []int, i int) {
	var (
		cur, left = i, 0
		tmp       = arr[i]
	)
	for {
		left = 2*cur + 1
		if left >= len(arr) {
			break
		}
		// get max child index
		max := left
		if left < len(arr)-1 && arr[left+1] > arr[left] {
			max = left + 1
		}
		// swap
		if tmp < arr[max] {
			arr[cur] = arr[max]
			cur = max
		} else {
			break
		}
	}
	// final position
	arr[cur] = tmp
}
