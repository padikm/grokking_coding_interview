package leet

import (
	"container/heap"
)

type minHeap [][]int

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i][1] < (*m)[j][1] || ((*m)[i][1] == (*m)[j][1] && (*m)[i][0] < (*m)[j][0])
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minHeap) Pop() interface{} {
	min := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return min
}

func KWeakestRows(mat [][]int, k int) []int {
	m := minHeap{}
	res := make([]int, 0)
	heap.Init(&m)
	for i := 0; i < len(mat); i++ {
		countOnes := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				countOnes++
			}
		}
		heap.Push(&m, []int{i, countOnes})
	}
	for len(m) > 0 && k != 0 {
		min := heap.Pop(&m).([]int)
		res = append(res, min[0])
		k--
	}

	return res

}
