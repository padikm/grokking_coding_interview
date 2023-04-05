package leet

import (
	"container/heap"
)

type maxHeapSetSize [][]int

func (m *maxHeapSetSize) Len() int {
	return len(*m)
}

func (m *maxHeapSetSize) Less(i, j int) bool {
	return (*m)[i][1] > (*m)[j][1]
}

func (m *maxHeapSetSize) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapSetSize) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *maxHeapSetSize) Pop() interface{} {
	max := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return max
}

func MinSetSize(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	maxHeap := maxHeapSetSize{}
	heap.Init(&maxHeap)
	for k, v := range m {
		heap.Push(&maxHeap, []int{k, v})
	}
	sum := 0
	count := 0
	for len(maxHeap) > 0 && sum < len(arr)/2 {
		e := heap.Pop(&maxHeap).([]int)
		sum = sum + e[1]
		count++
	}
	return count
}
