package leet

import (
	"container/heap"
)

type minHeapKth [][]int

func (m *minHeapKth) Len() int {
	return len(*m)
}

func (m *minHeapKth) Less(i, j int) bool {
	return (*m)[i][0] < (*m)[j][0]
}

func (m *minHeapKth) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapKth) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *minHeapKth) Pop() any {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}

func KthSmallest(matrix [][]int, k int) int {
	minHeap := minHeapKth{}
	heap.Init(&minHeap)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			heap.Push(&minHeap, []int{matrix[j][i], j, i})
		}
	}
	res := 0
	for len(minHeap) != 0 && k != 0 {
		resArray := heap.Pop(&minHeap).([]int)
		res = resArray[0]
		k--
	}
	return res
}
