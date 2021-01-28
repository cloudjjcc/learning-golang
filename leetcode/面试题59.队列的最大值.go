package main

func main() {
	queue := Constructor()
	println(queue.Max_value())
	queue.Push_back(1)
	queue.Push_back(0)
	queue.Push_back(1)
	queue.Push_back(2)
	queue.Push_back(2)
	queue.Push_back(3)
	println(queue.Pop_front())
	queue.Push_back(3)
	println(queue.Max_value())
	println(queue.Pop_front())
	println(queue.Pop_front())
	println(queue.Pop_front())
	println(queue.Pop_front())
	println(queue.Max_value())
}

type MaxQueue struct {
	data  []int
	mdata []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (q *MaxQueue) Max_value() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.mdata[0]
}

func (q *MaxQueue) Push_back(value int) {
	for len(q.mdata) > 0 && q.mdata[len(q.mdata)-1] < value {
		q.mdata = q.mdata[:len(q.mdata)-1]
	}
	q.mdata = append(q.mdata, value)
	q.data = append(q.data, value)
}

func (q *MaxQueue) Pop_front() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	if q.Max_value() == val {
		q.mdata = q.mdata[1:]
	}
	return val
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
