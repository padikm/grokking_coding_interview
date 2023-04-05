package leet

import (
	"container/heap"
)

type minHeapMedian [][]int

func (m *minHeapMedian) Len() int {
	return len(*m)
}

func (m *minHeapMedian) Less(i, j int) bool {
	return (*m)[i][0] < (*m)[j][0]
}

func (m *minHeapMedian) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapMedian) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minHeapMedian) Pop() interface{} {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}

func FindMedianSortedArrays(a []int, b []int) float64 {

	minHeap := minHeapMedian{}
	heap.Init(&minHeap)
	if len(a) > 0 {
		heap.Push(&minHeap, []int{a[0], 1})
	}
	if len(b) > 0 {
		heap.Push(&minHeap, []int{b[0], 2})
	}
	res := make([]int, 0)
	i, j := 1, 1
	for i < len(a) || j < len(b) {
		e := heap.Pop(&minHeap).([]int)
		res = append(res, e[0])
		if i < len(a) && e[1] == 1 {
			heap.Push(&minHeap, []int{a[i], 1})
			i++
		} else if j < len(b) && e[1] == 2 {
			heap.Push(&minHeap, []int{b[j], 2})
			j++
		}
	}
	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	for j < len(b) {
		res = append(res, a[j])
		j++
	}

	for len(minHeap) > 0 {
		e := heap.Pop(&minHeap).([]int)
		res = append(res, e[0])
	}
	median := 0.000000
	l := len(a) + len(b)
	if l%2 == 0 {
		l = l / 2
		median = (float64(res[l]) + float64(res[l-1])) / 2
	} else {
		l = l / 2
		median = float64(res[l])
	}
	return median
}
