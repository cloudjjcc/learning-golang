package sort

// BubbleSort 冒泡排序 时间复杂度O(n)2 空间复杂度O(1)
func BubbleSort(arr []int) {
	var exchangeFlag bool
	for right := len(arr) - 1; right > 0; right-- {
		exchangeFlag = false
		for left := 0; left < right; left++ {
			if arr[left] > arr[left+1] {
				arr[left], arr[left+1] = arr[left+1], arr[left]
				exchangeFlag = true
			}
		}
		if !exchangeFlag {
			return
		}
	}
}

// SelectionSort 选择排序 时间复杂度O(n)2 空间复杂度O(1)
// 每次从待排序数组中选择最大（小）的数
func SelectionSort(arr []int) {
	min := 0
	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

// InsertionSort 插入排序 时间复杂度O(n)2 空间复杂度O(1)
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

// MergeSort 归并排序 时间复杂度O(nlog n) 空间复杂度O(n)
func MergeSort(arr []int) {
	if len(arr) == 1 {
		return
	}
	mid := len(arr) / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])
	i1, i2 := 0, mid
	tempArr := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if i1 == mid {
			tempArr = append(tempArr, arr[i2:]...)
			break
		}
		if i2 == len(arr) {
			tempArr = append(tempArr, arr[i1:]...)
			break
		}
		if arr[i1] <= arr[i2] {
			tempArr = append(tempArr, arr[i1])
			i1++
		} else {
			tempArr = append(tempArr, arr[i2])
			i2++
		}
	}
	copy(arr, tempArr)
}

// HeapSort 堆排序 时间复杂度O(n log n) 空间复杂度O(1)
func HeapSort(arr []int) {
	buildHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapDown(arr[:i], 0)
	}
}

// O(n)
func buildHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapDown(arr, i)
	}
}

// O(log n)
func heapDown(arr []int, i int) {
	hole := i
	for {
		leftCh := 2*hole + 1
		if leftCh >= len(arr) {
			return
		}
		maxCh := leftCh
		rightCh := 2*hole + 2
		if rightCh < len(arr) && arr[rightCh] > arr[maxCh] {
			maxCh = rightCh
		}
		if arr[hole] < arr[maxCh] {
			arr[hole], arr[maxCh] = arr[maxCh], arr[hole]
			hole = maxCh
		} else {
			return
		}
	}
}

// QuickSort 快速排序O(nlogn) O(logn)
func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	partitionFn := func() int {
		pivotVal := nums[len(nums)-1]
		left, right := 0, 0
		for right < len(nums)-1 {
			if nums[right] < pivotVal {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
			right++
		}
		nums[left], nums[len(nums)-1] = nums[len(nums)-1], nums[left]
		return left
	}
	pivot := partitionFn()
	QuickSort(nums[:pivot])
	QuickSort(nums[pivot+1:])
}
