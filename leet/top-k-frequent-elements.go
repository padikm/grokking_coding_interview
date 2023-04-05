package leet

import (
	"container/heap"
)

type maxHeapFreq [][]int

func (m *maxHeapFreq) Len() int {
	return len(*m)
}

func (m *maxHeapFreq) Less(i, j int) bool {
	return (*m)[i][1] > (*m)[j][1]
}

func (m *maxHeapFreq) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapFreq) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *maxHeapFreq) Pop() interface{} {
	p := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return p
}

func TopKFrequent(nums []int, k int) []int {
	mf := maxHeapFreq{}
	res := make([]int, 0)
	heap.Init(&mf)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		heap.Push(&mf, []int{k, v})
	}
	for len(mf) > 0 && k != 0 {
		max := heap.Pop(&mf).([]int)
		res = append(res, max[0])
		k--
	}
	return res
}
