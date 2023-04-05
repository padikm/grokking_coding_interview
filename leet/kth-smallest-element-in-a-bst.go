package leet

import "container/heap"

/**
 * Definition for a binary tree node.
 * type TreeNode structtest {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type minHeapBst []int

func (m *minHeapBst) Len() int {
	return len(*m)
}

func (m *minHeapBst) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeapBst) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeapBst) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeapBst) Pop() interface{} {
	p := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return p
}

func kthSmallest(root *TreeNode, k int) int {
	minHeap := minHeapBst{}
	heap.Init(&minHeap)
	inorderKth(root, k, &minHeap)
	var res int
	for len(minHeap) > 0 && k != 0 {
		k--
		res = heap.Pop(&minHeap).(int)
	}
	return res
}

func inorderKth(root *TreeNode, k int, m *minHeapBst) {
	if root == nil {
		return
	}
	inorderKth(root.Left, k, m)
	heap.Push(m, root.Val)
	inorderKth(root.Right, k, m)
}
