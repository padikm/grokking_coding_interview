package leet

import (
	"container/heap"
	"strconv"
)

type maxHeapFreqString [][]string

func (m *maxHeapFreqString) Len() int {
	return len(*m)
}

func (m *maxHeapFreqString) Less(i, j int) bool {
	v1, _ := strconv.Atoi((*m)[i][1])
	v2, _ := strconv.Atoi((*m)[j][1])
	return v1 > v2
}

func (m *maxHeapFreqString) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapFreqString) Push(x interface{}) {
	*m = append(*m, x.([]string))
}

func (m *maxHeapFreqString) Pop() interface{} {
	max := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return max
}

func FrequencySort(s string) string {
	maxHeap := maxHeapFreqString{}
	heap.Init(&maxHeap)
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}

	for k, v := range m {
		heap.Push(&maxHeap, []string{k, strconv.Itoa(v)})
	}
	var res string
	for len(maxHeap) > 0 {
		s := heap.Pop(&maxHeap).([]string)
		v, _ := strconv.Atoi(s[1])
		for v != 0 {
			res = res + s[0]
			v--
		}
	}

	return res
}
