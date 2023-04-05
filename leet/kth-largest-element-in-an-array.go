package leet

import "container/heap"

type maxHeapKth []int

func (m *maxHeapKth) Len() int {
	return len(*m)
}

func (m *maxHeapKth) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeapKth) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]

}

func (m *maxHeapKth) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeapKth) Pop() interface{} {
	max := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return max
}

func FindKthLargest(nums []int, k int) int {
	m := maxHeapKth{}
	m = nums
	max := 0
	heap.Init(&m)
	for i := 0; i < k; i++ {
		max = heap.Pop(&m).(int)
	}
	return max
}
