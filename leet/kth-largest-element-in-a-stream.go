package leet

import "container/heap"

/*
kLarge := leet.Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kLarge.Add(3)) // return 4
	fmt.Println(kLarge.Add(5))
	fmt.Println(kLarge.Add(10))
	fmt.Println(kLarge.Add(9))
	fmt.Println(kLarge.Add(4))
*/
type maxHeapKthLargest []int

func (m *maxHeapKthLargest) Len() int {
	return len(*m)
}

func (m *maxHeapKthLargest) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *maxHeapKthLargest) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeapKthLargest) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeapKthLargest) Pop() any {
	s := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return s
}

type KthLargest struct {
	k int
	m maxHeapKthLargest
}

func Constructor(k int, nums []int) KthLargest {
	kthLarge := KthLargest{}
	kthLarge.k = k
	heap.Init(&kthLarge.m)
	for _, v := range nums {
		heap.Push(&kthLarge.m, v)
	}
	return kthLarge
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.m, val)
	for len(this.m) > this.k {
		heap.Pop(&this.m)
	}
	return this.m[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
