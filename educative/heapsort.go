package educative

import "fmt"

func MinHeap(a []int) []int {
	for i := len(a) / 2; i >= 0; i-- {
		minHeapify(a, i)
	}
	return a
}

func minHeapify(a []int, i int) {
	left := 2 * i
	right := (2 * i) + 1
	small := i
	if left < len(a) && a[left] < a[i] {
		small = left
	}
	if right < len(a) && a[right] < a[small] {
		small = right
	}
	if small != i {
		a[small], a[i] = a[i], a[small]
		minHeapify(a, small)
	}
}

func PopMin(a []int) []int {
	if len(a) <= 1 {
		return nil
	}
	fmt.Println(a[1])
	a[1] = a[len(a)-1]
	a = a[:len(a)-1]
	minHeapify(a, 1)
	return a
}

func PushMin(a []int, key int) []int {
	a = append(a, key)
	heapI := (len(a) - 1) / 2
	for i := heapI; i > 0; i-- {
		minHeapify(a, i)
	}
	return a
}
