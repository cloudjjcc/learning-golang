package leetcode

type KthLargest struct {
	h *Heap
	k int
}

func KthLargestConstructor(k int, nums []int) KthLargest {
	data := make([]int, 0, k)
	h := &Heap{data: data}
	kl := KthLargest{
		h: h,
		k: k,
	}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.h.data) < this.k {
		this.h.Push(val)
	} else {
		if min := this.h.Peek(); min < val {
			this.h.Pop()
			this.h.Push(val)
		}
	}
	return this.h.Peek()
}

// Heap 二叉堆
type Heap struct {
	data []int
}

// Push 压入元素
func (h *Heap) Push(val int) {
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

// Pop 弹出堆顶元素
func (h *Heap) Pop() int {
	ele := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return ele
}

// Peek 查看堆顶元素
func (h *Heap) Peek() int {
	ele := h.data[0]
	return ele
}

// 元素上浮
func (h *Heap) up(i int) {
	hole := i
	val := h.data[i]
	for {
		p := (hole - 1) / 2
		if p < 0 || p == hole {
			break
		}
		if h.data[p] > val {
			h.data[hole] = h.data[p]
			hole = p
		} else {
			break
		}
	}
	h.data[hole] = val
}

// 元素下沉
func (h *Heap) down(i int) {
	hole := i
	val := h.data[i]
	for {
		left := 2*hole + 1
		if left >= len(h.data) {
			break
		}
		min := left
		if right := left + 1; right < len(h.data) && h.data[left] > h.data[right] {
			min = right
		}
		if val > h.data[min] {
			h.data[hole] = h.data[min]
			hole = min
		} else {
			break
		}
	}
	h.data[hole] = val
}

// NewHeap 构建堆
func NewHeap(data []int) *Heap {
	h := &Heap{data: data}
	for i := len(data) / 2; i >= 0; i-- {
		h.down(i)
	}
	return h
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
