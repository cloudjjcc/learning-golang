package leetcode

//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	if k == 0 {
		return []int{}
	}
	res := arr[:k]
	// build max heap
	for i := len(res)/2 - 1; i >= 0; i-- {
		bubbleDown(res, i)
	}
	// get min k
	for i := k; i < len(arr); i++ {
		if res[0] > arr[i] {
			res[0], arr[i] = arr[i], res[0]
			bubbleDown(res, 0)
		}
	}
	return res
}

func bubbleDown(arr []int, i int) {
	var (
		tmp  = arr[i]
		hole = i
		left = 0
	)
	for {
		left = 2*hole + 1
		if left >= len(arr) {
			break
		}
		max := left
		if left < len(arr)-1 && arr[left+1] > arr[left] {
			max = left + 1
		}
		if tmp < arr[max] {
			arr[hole] = arr[max]
			hole = max
		} else {
			break
		}
	}
	arr[hole] = tmp
}
