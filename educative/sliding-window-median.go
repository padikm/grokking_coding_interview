package educative

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (m *maxHeap) Len() int {
	//TODO implement me
	return len(*m)
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Swap(i, j int) {
	//TODO implement me
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	big := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return big
}

type minHeap []int

func (m *minHeap) Len() int {
	//TODO implement me
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Swap(i, j int) {
	//TODO implement me
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	big := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return big
}

func MedianSlidingWindow(nums []int, k int) []int {
	maxh := maxHeap{}
	minH := minHeap{}
	heap.Init(&minH)
	heap.Init(&maxh)
	fp := 0

	for i := 0; i < k; i++ {
		heap.Push(&maxh, nums[i])
	}
	res := make([]int, 0)
	for len(maxh) > (k/2)+1 {
		heap.Push(&minH, heap.Pop(&maxh))
	}
	i := k
	for true {
		if i > len(nums) {
			break
		}
		res = append(res, maxh[0])
		if maxh[0] == nums[fp] {
			heap.Pop(&maxh)
		} else if minH[0] == nums[fp] {
			heap.Pop(&minH)
		}

		if i < len(nums) && nums[i] < maxh[0] {
			heap.Push(&maxh, nums[i])
		} else if i < len(nums) {
			heap.Push(&minH, nums[i])
		}

		for len(maxh) > (k/2)+1 {
			heap.Push(&minH, heap.Pop(&maxh))
		}

		fmt.Println(maxh, minH)
		fp++
		i++

	}
	return res
}
