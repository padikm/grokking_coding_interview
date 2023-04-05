package recursion

func SelectionSort(a []int) []int {
	return selectionSort(a, 0, 0, 0)
}

func selectionSort(a []int, i, j, small int) []int {
	if i == len(a) {
		return a
	}
	if j == len(a) {
		a[i], a[small] = a[small], a[i]
		return selectionSort(a, i+1, i+1, i+1)
	}
	if a[small] > a[j] {
		small = j
	}
	return selectionSort(a, i, j+1, small)
}
