package educative

import (
	"container/heap"
)

type minHeapTasks [][]int

func (m *minHeapTasks) Len() int {
	return len(*m)
}

func (m *minHeapTasks) Less(i, j int) bool {
	return (*m)[i][0] < (*m)[j][0]
}

func (m *minHeapTasks) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapTasks) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *minHeapTasks) Pop() any {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}

type minHeapEndTime [][]int

func (m *minHeapEndTime) Len() int {
	return len(*m)
}

func (m *minHeapEndTime) Less(i, j int) bool {
	return (*m)[i][1] < (*m)[j][1]
}

func (m *minHeapEndTime) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapEndTime) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *minHeapEndTime) Pop() any {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}

func Tasks(T [][]int) int {
	minHeap := minHeapTasks(T)
	minHeapMachine := minHeapEndTime{}
	heap.Init(&minHeap)
	count := 0
	for len(minHeap) > 0 {
		m := heap.Pop(&minHeap).([]int)
		if len(minHeapMachine) == 0 {
			heap.Push(&minHeapMachine, m)
			count++
		} else {
			mS := heap.Pop(&minHeapMachine).([]int)
			if m[0] < mS[1] {
				count++
				heap.Push(&minHeapMachine, m)
				heap.Push(&minHeapMachine, mS)
			} else {
				heap.Push(&minHeapMachine, m)
			}
		}
	}
	return count
}
