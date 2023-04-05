package leet

import (
	"container/heap"
	"strconv"
)

type maxHeapReOrgStr [][]string

func (m *maxHeapReOrgStr) Len() int {
	return len(*m)
}

func (m *maxHeapReOrgStr) Less(i, j int) bool {
	v1, _ := strconv.Atoi((*m)[i][1])
	v2, _ := strconv.Atoi((*m)[j][1])
	return v1 >= v2
}

func (m *maxHeapReOrgStr) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapReOrgStr) Push(x any) {
	*m = append(*m, x.([]string))
}

func (m *maxHeapReOrgStr) Pop() any {
	max := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return max
}

func ReorganizeString(s string) string {
	maxHeap := maxHeapReOrgStr{}
	heap.Init(&maxHeap)
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}

	for k, v := range m {
		i := strconv.Itoa(v)
		heap.Push(&maxHeap, []string{k, i})
	}
	var res string
	for len(maxHeap) > 0 {
		if maxHeap[0][1] == "0" {
			heap.Pop(&maxHeap)
			continue
		}
		v := heap.Pop(&maxHeap).([]string)
		res = res + v[0]
		v1, _ := strconv.Atoi(v[1])
		v1--
		v[1] = strconv.Itoa(v1)
		if len(maxHeap) > 0 && maxHeap[0][1] != "0" {
			v := heap.Pop(&maxHeap).([]string)
			res = res + v[0]
			v1, _ := strconv.Atoi(v[1])
			v1--
			v[1] = strconv.Itoa(v1)
			heap.Push(&maxHeap, v)
		}
		heap.Push(&maxHeap, v)
	}
	var prev string
	for _, v := range res {
		str := string(v)
		if str == prev {
			return ""
		}
		prev = str
	}
	return res
}
