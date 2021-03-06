package nowcoder

type sortArr []int

func (s sortArr) Len() int {
	return len(s)
}

func (s sortArr) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 冒泡排序O(n*n)
func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := len(arr) - 1; i > 0; i-- {
		swapFlag := false
		for j := 1; j <= i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swapFlag = true
			}
		}
		if !swapFlag {
			return
		}
	}
}

// 插入排序
func InsertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	var (
		j, tmp int
	)
	for i := 1; i < len(arr); i++ {
		tmp = arr[i]
		for j = i; j >= 1 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

// 希尔排序（O(n**k)）
func ShellSort(arr []int) {
	N := len(arr)
	if N <= 1 {
		return
	}
	var (
		j, tmp int
	)
	// 增量序列
	for k := N / 2; k > 0; k /= 2 {
		// 插入排序
		for i := k; i < N; i++ {
			tmp = arr[i]
			for j = i; j >= k && tmp < arr[j-k]; j -= k {
				arr[j] = arr[j-k]
			}
			arr[j] = tmp
		}
	}
}

// 选择排序(O(n**n))
func SelectionSort(arr []int) {
	N := len(arr)
	if N <= 1 {
		return
	}
	var (
		min int
	)
	for i := 0; i < N; i++ {
		// 选择最小
		min = i
		for j := i + 1; j < N; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

// 堆排序（O(n*log n)）
func HeapSort(arr []int) {
	N := len(arr)
	if N <= 1 {
		return
	}
	// build max heap
	// 从最后一个非叶子节点开始
	for i := N/2 - 1; i >= 0; i-- {
		percDown(arr, i, N)
	}
	for i := N - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		percDown(arr, 0, i)
	}
}

// 下滤
// 使arr[i]为根的二叉堆保持堆序性
// N为元素个数
func percDown(arr []int, i int, N int) {
	if N <= 1 {
		return
	}
	var (
		lChildIdx int      //左孩子索引
		tmp       = arr[i] //待下浮元素
		holeIdx   = i      //空穴索引
		maxIdx    = 0      //左右儿子中较大者索引
	)
	for {
		// 左孩子索引
		lChildIdx = 2*holeIdx + 1
		if lChildIdx >= N { // 不存在左孩子下滤结束
			break
		}
		// 找到最大孩子索引
		maxIdx = lChildIdx
		if lChildIdx != N-1 && arr[lChildIdx] < arr[lChildIdx+1] { //存在右孩子
			maxIdx = lChildIdx + 1
		}
		// 满足堆序性
		if tmp > arr[maxIdx] {
			break
		}
		// 空穴下移
		arr[holeIdx] = arr[maxIdx]
		holeIdx = maxIdx
	}
	arr[holeIdx] = tmp
}

// 归并排序（O(n*log n)
func MergeSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	mergeSort(arr, make([]int, len(arr)))
}

func mergeSort(arr, tempArr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	mergeSort(arr[:mid], tempArr[:mid])
	mergeSort(arr[mid:], tempArr[mid:])
	merge(arr, tempArr, mid)
}

func merge(arr, tempArr []int, mid int) {
	var (
		p1 = 0
		p2 = mid
		i  = 0
	)
	for {
		if p1 >= mid || p2 >= len(arr) {
			break
		}
		if arr[p1] < arr[p2] {
			tempArr[i] = arr[p1]
			p1++
		} else {
			tempArr[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 < mid {
		tempArr[i] = arr[p1]
		p1++
		i++
	}
	for p2 < len(arr) {
		tempArr[i] = arr[p2]
		p2++
		i++
	}
	// copy temp array to arr
	for j := 0; j < len(arr); j++ {
		arr[j] = tempArr[j]
	}
}

// 快速排序(O(n* log n))
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	if len(arr) < 10 {
		InsertionSort(arr)
		return
	}
	// pivot
	i, j := doPivot(arr)
	QuickSort(arr[:i])
	QuickSort(arr[j:])
}

func doPivot(arr []int) (int, int) {
	pivot := median3(arr)
	left, right := 1, len(arr)-3
	for {
		for left <= right && arr[left] < pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		} else {
			break
		}
	}
	return right, left
}

func median3(arr []int) int {
	if len(arr) < 3 {
		panic("三数中值法需保证数组中至少有三个数")
	}
	// arr[0]<=arr[len(arr)/2]<=arr[len(arr)-1]
	a, b, c := 0, len(arr)/2, len(arr)-1
	if arr[a] > arr[b] {
		arr[a], arr[b] = arr[b], arr[a]
	}
	if arr[a] > arr[c] {
		arr[a], arr[c] = arr[c], arr[a]
		if arr[b] > arr[c] {
			arr[b], arr[c] = arr[c], arr[b]
		}
	}
	// swap arr[len(arr)/2],arr[len(arr)-2]
	arr[b], arr[len(arr)-2] = arr[len(arr)-2], arr[b]
	return arr[len(arr)-2]
}
