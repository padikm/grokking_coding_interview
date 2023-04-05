package recursion

import (
	"container/heap"
	"fmt"
	"strings"
)

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

func PrintSubSeqOfThreeStr() {
	h := maxHeapKth{}
	heap.Init(&h)
	printSubSeqOfStr("inclu", "socluue", "anyclue", 0, 0, &h)
	fmt.Println(heap.Pop(&h))

}

func printSubSeqOfStr(s, s1, s2 string, j int, i int, h *maxHeapKth) {
	if i == len(s) {
		return
	}

	if j == len(s) {
		printSubSeqOfStr(s, s1, s2, i+1, i+1, h)
		return
	}
	s3 := s[i : j+1]
	if strings.Contains(s1, s3) && strings.Contains(s2, s3) {
		heap.Push(h, len(s3))
	}
	fmt.Println(s[i : j+1])
	printSubSeqOfStr(s, s1, s2, j+1, i, h)
}
