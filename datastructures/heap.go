package datastructures

type Heap struct {
	data []int
}

// 构建优先级队列
func NewHeap(eles []int) *Heap {
	pq := &Heap{
		data: eles,
	}
	for i := pq.Size()/2 - 1; i >= 0; i-- {
		pq.down(i)
	}
	return pq
}

// 入队
func (pq *Heap) Push(ele int) {
	pq.data = append(pq.data, ele)
	pq.up(pq.Size() - 1)
}

// 出队
func (pq *Heap) Pop() int {
	if pq.Size() == 0 {
		return 0
	}
	root := pq.data[0]
	pq.data[0] = pq.data[pq.Size()-1]
	pq.down(0)
	return root
}

// 元素数量
func (pq *Heap) Size() int {
	return len(pq.data)
}

// 下滤
func (pq *Heap) down(i int) {
	var (
		hole   = i
		leftCh = 0
		maxCh  = 0
		tmp    = pq.data[i]
	)
	for {
		leftCh = hole<<1 + 1
		if leftCh >= pq.Size() { // 不存在左孩子，下滤结束
			break
		}
		// 获得最大孩子节点
		maxCh = leftCh
		if leftCh != pq.Size()-1 && pq.data[leftCh] < pq.data[leftCh+1] {
			maxCh = leftCh + 1
		}
		// 比较
		if tmp < pq.data[maxCh] {
			pq.data[hole] = pq.data[maxCh]
			hole = maxCh
		} else {
			break
		}
	}
	pq.data[hole] = tmp
}

// 上滤
func (pq *Heap) up(i int) {
	var (
		hole   = i
		parent = 0
		tmp    = pq.data[i]
	)
	for {
		// 获得父元素位置
		parent = (hole - 1) / 2
		if parent == hole || pq.data[parent] > tmp {
			break
		}
		pq.data[hole] = pq.data[parent]
		hole = parent
	}
	pq.data[hole] = tmp
}
