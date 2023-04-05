package leet

import (
	"container/heap"
	"math"
)

type minHeapPoints [][]float64

func (m *minHeapPoints) Len() int {
	return len(*m)
}

func (m *minHeapPoints) Less(i, j int) bool {
	return (*m)[i][0] < (*m)[j][0]
}

func (m *minHeapPoints) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapPoints) Push(x interface{}) {
	*m = append(*m, x.([]float64))
}

func (m *minHeapPoints) Pop() interface{} {
	min := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return min
}

func KClosest(p [][]int, k int) [][]int {
	minHeap := minHeapPoints{}
	heap.Init(&minHeap)
	for i := 0; i < len(p); i++ {
		d := distance(p[i][0], p[i][1])
		heap.Push(&minHeap, []float64{d, float64(i)})
	}
	var res [][]int
	for len(minHeap) > 0 && k != 0 {
		r := heap.Pop(&minHeap).([]float64)
		res = append(res, p[int(r[1])])
		k--
	}
	return res
}

func distance(i int, j int) float64 {
	i = i * i
	j = j * j
	return math.Sqrt(float64(i) + float64(j))
}
