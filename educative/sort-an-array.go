package educative

func SortArray(nums []int) []int {
	return minHeapSort(nums)
}

func minHeapSort(a []int) []int {
	for i := len(a) / 2; i >= 0; i-- {
		minfyHeapSort(a, i)
	}
	res := make([]int, 0)
	item := 0
	for len(a) > 0 {
		item, a = popMinHeapToSort(a)
		res = append(res, item)
	}
	return res
}

func popMinHeapToSort(a []int) (int, []int) {
	item := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	minfyHeapSort(a, 0)
	return item, a
}

func minfyHeapSort(a []int, i int) {
	left := 2 * i
	right := (2 * i) + 1
	small := i
	if left < len(a) && a[left] < a[small] {
		small = left
	}

	if right < len(a) && a[right] < a[small] {
		small = right
	}

	if small != i {
		a[small], a[i] = a[i], a[small]
		minfyHeapSort(a, small)
	}
}
