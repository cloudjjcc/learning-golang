package leetcode

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }

type MedianFinder struct {
	maxHeap hp
	minHeap hp
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	maxHp, minHp := &mf.maxHeap, &mf.minHeap
	if minHp.Len() == 0 || num <= -minHp.IntSlice[0] {
		heap.Push(minHp, -num)
		if maxHp.Len()+1 < minHp.Len() {
			heap.Push(maxHp, -heap.Pop(minHp).(int))
		}
	} else {
		heap.Push(maxHp, num)
		if maxHp.Len() > minHp.Len() {
			heap.Push(minHp, -heap.Pop(maxHp).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	maxHp, minHp := &mf.maxHeap, &mf.minHeap
	if minHp.Len() > maxHp.Len() {
		return float64(-minHp.IntSlice[0])
	}
	return float64(maxHp.IntSlice[0]-minHp.IntSlice[0]) / 2
}
