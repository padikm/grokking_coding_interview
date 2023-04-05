package geekForGeek

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	e := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return e
}

func MinimalMovesToFormString(s string) {
	mHeap := minHeap{}
	heap.Init(&mHeap)
	minimalMovesToFormString(s, "", s, 0, 0, &mHeap)
	fmt.Println(heap.Pop(&mHeap))
}

func minimalMovesToFormString(s string, res string, s1 string, count, j int, minHeap *minHeap) {

	if s1 == res {
		heap.Push(minHeap, count)
		return
		//return
	}
	if len(res) >= len(s1) {
		return
		//return
	}

	if j == len(s1) {
		return
		//return
	}

	for i := j; i < len(s); i++ {
		//return min(minimalMovesToFormString(s, res+string(s[i]), s1, count+1, i+1), minimalMovesToFormString(s, res+res, s1, count+1, i+1))
		minimalMovesToFormString(s, res+string(s[i]), s1, count+1, i+1, minHeap)
		minimalMovesToFormString(s, res+res, s1, count+1, i+1, minHeap)

	}
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
