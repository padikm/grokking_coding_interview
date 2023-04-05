package educative

func twoDimHeap(a [][]int) {
	for i := len(a) / 2; i >= 0; i-- {
		twoDHeapify(a, i)
	}
}

func twoDHeapify(a [][]int, i int) {
	left := 2 * i
	right := (2 * i) + 1
	small := i
	if left < len(a) && a[left][0] < a[small][0] {
		small = left
	}

	if right < len(a) && a[right][0] < a[small][0] {
		small = right
	}

	if small != i {
		a[small], a[i] = a[i], a[small]
		twoDHeapify(a, small)
	}
}

func popMinHeap(a [][]int) ([]int, [][]int) {
	small := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	twoDHeapify(a, 0)
	return small, a
}

func pushMinHeap(a [][]int, i, j int) [][]int {
	a = append(a, []int{i, j})
	twoDHeapify(a, len(a)/2)
	return a
}

func twoDMaxHeap(a [][]int) {
	for i := len(a) / 2; i >= 0; i-- {
		twoDMaxHeapify(a, i)
	}
}

func twoDMaxHeapify(a [][]int, i int) {
	left := 2 * i
	right := (2 * i) + 1
	small := i
	if left < len(a) && a[left][1] > a[small][1] {
		small = left
	}

	if right < len(a) && a[right][1] > a[small][1] {
		small = right
	}
	if small != i {
		a[small], a[i] = a[i], a[small]
		twoDMaxHeapify(a, small)
	}
}

func popMaxHeap(a [][]int) ([]int, [][]int) {
	small := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	twoDMaxHeapify(a, 0)
	return small, a
}

func pushMaxHeap(a [][]int, i, j int) [][]int {
	a = append(a, []int{i, j})
	//twoDMaxHeapify(a, 0)
	twoDMaxHeap(a)
	return a
}
