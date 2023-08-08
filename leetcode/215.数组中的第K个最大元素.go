package leetcode

import "math/rand"

// 运用小根堆O(nlogk) O(k)
func findKthLargest(nums []int, k int) int {
	buildMinHeap(nums[:k])
	for i := k; i < len(nums); i++ {
		if nums[0] < nums[i] {
			nums[0], nums[i] = nums[i], nums[0]
			minHeapDown(nums[:k], 0)
		}
	}
	return nums[0]
}

func buildMinHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		minHeapDown(arr, i)
	}
}
func minHeapDown(arr []int, i int) {
	root := i
	for {
		leftCh := 2*root + 1
		if leftCh >= len(arr) {
			return
		}
		minCh := leftCh
		rightCh := 2*root + 2
		if rightCh < len(arr) && arr[rightCh] < arr[minCh] {
			minCh = rightCh
		}
		if arr[minCh] >= arr[root] { //符合堆序
			return
		}
		arr[minCh], arr[root] = arr[root], arr[minCh]
		root = minCh
	}
}

// 运用快速选择算法 O(n)
func findKthLargest2(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}
func quickSelect(nums []int, begin, end, k int) int {
	pivot := partition(nums, begin, end)
	if pivot == k {
		return nums[k]
	}
	if pivot < k {
		return quickSelect(nums, pivot+1, end, k)
	}
	return quickSelect(nums, begin, pivot-1, k)
}
func partition(nums []int, begin, end int) int {
	randIdx := rand.Int()%(end-begin+1) + begin
	nums[randIdx], nums[end] = nums[end], nums[randIdx]
	pivotVal := nums[end]
	left, right := begin, begin
	for ; right < end; right++ {
		if nums[right] < pivotVal {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	nums[left], nums[end] = nums[end], nums[left]
	return left
}
