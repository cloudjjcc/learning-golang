package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

type vf struct {
	v int
	f int
}

// 小顶堆 O(nlogk) O(n)
func topKFrequent(nums []int, k int) []int {
	// 获取 val=>freq 映射
	fByVal := make(map[int]int)
	for _, v := range nums {
		fByVal[v]++
	}
	// top k
	heap := make([]vf, 0, k+1)
	heapPushFn := func(val vf) {
		heap = append(heap, val)
		minHeapUp1(heap, len(heap)-1)
		fmt.Printf("push:%v=>%v\n", val, heap)
	}
	heapPopFn := func() vf {
		tmp := heap[0]
		heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
		heap = heap[:len(heap)-1]
		minHeapDown1(heap, 0)
		fmt.Printf("pop:%v=>%v\n", heap, tmp)
		return tmp
	}
	// top k
	for v, f := range fByVal {
		heapPushFn(vf{
			v: v,
			f: f,
		})
		if len(heap) > k {
			heapPopFn()
		}
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heapPopFn().v
	}
	return ans
}
func buildMinHeap1(arr []vf) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		minHeapDown1(arr, i)
	}
}

// 上浮
func minHeapUp1(arr []vf, cur int) {
	for {
		root := (cur - 1) / 2
		if root == cur || arr[root].f <= arr[cur].f {
			return
		}
		arr[root], arr[cur] = arr[cur], arr[root]
		cur = root
	}
}

// 下沉
func minHeapDown1(arr []vf, cur int) {
	for {
		leftCh := 2*cur + 1
		if leftCh >= len(arr) {
			return
		}
		minCh := leftCh
		rightCh := leftCh + 1
		if rightCh < len(arr) && arr[rightCh].f < arr[leftCh].f {
			minCh = rightCh
		}
		if arr[cur].f <= arr[minCh].f {
			return
		}
		arr[cur], arr[minCh] = arr[minCh], arr[cur]
		cur = minCh
	}
}

// 快速选择法 O(n)
func topKFrequent2(nums []int, k int) []int {
	// 获取 val=>freq 映射
	fByVal := make(map[int]int)
	for _, v := range nums {
		fByVal[v]++
	}
	// vfs
	vfs := make([]vf, 0, len(fByVal))
	for v, f := range fByVal {
		vfs = append(vfs, vf{
			v: v,
			f: f,
		})
	}
	ret := make([]int, k)
	rand.Seed(time.Now().Unix())
	quickSelect1(vfs, 0, len(vfs)-1, len(vfs)-k, ret, 0)
	return ret
}
func quickSelect1(nums []vf, begin, end, k int, ret []int, retIndex int) {
	if begin > end {
		return
	}
	pivotIndex := partition1(nums, begin, end)
	if pivotIndex >= k {
		for i := pivotIndex; i <= end; i++ {
			ret[retIndex] = nums[i].v
			retIndex++
		}
		if retIndex == len(ret) {
			return
		}
		quickSelect1(nums, begin, pivotIndex-1, k, ret, retIndex)
		return
	}
	quickSelect1(nums, pivotIndex+1, end, k, ret, retIndex)
}
func partition1(nums []vf, begin, end int) int {
	pi := rand.Int()%(end-begin+1) + begin
	nums[pi], nums[end] = nums[end], nums[pi]
	left, right := begin, begin
	for ; right < end; right++ {
		if nums[right].f < nums[end].f {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	nums[left], nums[end] = nums[end], nums[left]
	return left
}
