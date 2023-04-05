package leet

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode structtest {
 *     Val int
 *     Next *ListNode
 * }
 */
type minHeapList []int

func (m *minHeapList) Len() int {
	return len(*m)
}

func (m *minHeapList) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeapList) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapList) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeapList) Pop() any {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := minHeapList{}
	heap.Init(&minHeap)
	for _, v := range lists {
		for v != nil {
			heap.Push(&minHeap, v.Val)
			v = v.Next
		}
	}
	var res *ListNode
	var head *ListNode
	for len(minHeap) > 0 {
		e := heap.Pop(&minHeap).(int)
		l := ListNode{Val: e}
		if res == nil {
			res = &l
			head = res
		} else {
			res.Next = &l
			res = res.Next
		}
	}
	return head
}
