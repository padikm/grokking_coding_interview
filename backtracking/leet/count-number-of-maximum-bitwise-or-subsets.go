package leet

import (
	"container/heap"
)

type mHeap []int

func (h *mHeap) Len() int {
	return len(*h)
}

func (h *mHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *mHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *mHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *mHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func CountMaxOrSubsets(nums []int) int {
	var maxHeap mHeap
	heap.Init(&maxHeap)
	countMaxOrSubsets1(nums, []int{}, &maxHeap, 0)
	m := heap.Pop(&maxHeap).(int)
	count := 1
	for len(maxHeap) != 0 {
		v := heap.Pop(&maxHeap).(int)
		if m != v {
			break
		}
		count++
	}
	return count
}

func countMaxOrSubsets1(nums []int, res []int, maxHeap *mHeap, j int) {
	tmp := 0
	for k := 0; k < len(res); k++ {
		tmp |= res[k]
	}
	heap.Push(maxHeap, tmp)

	for i := j; i < len(nums); i++ {
		res = append(res, nums[i])
		countMaxOrSubsets1(nums, res, maxHeap, i+1)
		res = res[:len(res)-1]
	}
}
