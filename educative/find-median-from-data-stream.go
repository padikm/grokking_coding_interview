package educative

import "container/heap"

/*

m := educative.Constructor()
	m.AddNum(2)
	m.AddNum(8)
	m.AddNum(4)
	fmt.Println(m.FindMedian())
	m.AddNum(5)
	fmt.Println(m.FindMedian())
*/

type MedianFinder struct {
	minHeap minHeapMedian
	maxHeap maxHeapMedian
}

func Constructor() MedianFinder {
	m := MedianFinder{}
	heap.Init(&m.maxHeap)
	heap.Init(&m.minHeap)
	return m
}

func (m *MedianFinder) AddNum(num int) {
	if len(m.maxHeap) == 0 || num < m.maxHeap[0] {
		heap.Push(&m.maxHeap, num)
	} else {
		heap.Push(&m.minHeap, num)
	}

	if len(m.maxHeap) > len(m.minHeap)+1 {
		heap.Push(&m.minHeap, heap.Pop(&m.maxHeap))
	} else if len(m.minHeap) > len(m.maxHeap) {
		heap.Push(&m.maxHeap, heap.Pop(&m.minHeap))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if len(m.minHeap) == len(m.maxHeap) {
		return (float64(m.minHeap[0]) + float64(m.maxHeap[0])) / 2.0
	}
	return float64(m.maxHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type minHeapMedian []int

func (m *minHeapMedian) Len() int {
	return len(*m)
}

func (m *minHeapMedian) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeapMedian) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapMedian) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeapMedian) Pop() interface{} {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s

}

type maxHeapMedian []int

func (m *maxHeapMedian) Len() int {
	return len(*m)
}

func (m *maxHeapMedian) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeapMedian) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapMedian) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeapMedian) Pop() interface{} {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}
