package leet

import (
	"container/heap"
)

type minHeapPair [][]int

func (m *minHeapPair) Len() int {
	return len(*m)
}

func (m *minHeapPair) Less(i, j int) bool {
	return ((*m)[i][0] + (*m)[i][1]) < ((*m)[j][0] + (*m)[j][1])
}

func (m *minHeapPair) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapPair) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minHeapPair) Pop() interface{} {
	min := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return min
}

func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	minHeap := minHeapPair{}
	res := make([][]int, 0)
	heap.Init(&minHeap)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if len(minHeap) > 0 && len(minHeap) >= k && nums1[i]+nums2[j] > minHeap[0][0]+minHeap[0][1] {
				break
			}
			if len(minHeap) < k {
				heap.Push(&minHeap, []int{nums1[i], nums2[j]})
			} else if nums1[i]+nums2[j] < minHeap[0][0]+minHeap[0][1] {
				_ = heap.Pop(&minHeap).([]int)
				heap.Push(&minHeap, []int{nums1[i], nums2[j]})
			}
		}
	}
	for len(minHeap) > 0 && k != 0 {
		res = append(res, heap.Pop(&minHeap).([]int))
		k--
	}
	return res
}
