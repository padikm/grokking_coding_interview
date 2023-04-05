package recursion

func BubbleSort(a []int) []int {
	return bubbleSort(a, 0, 0)
}

func bubbleSort(a []int, i, j int) []int {
	if i == len(a) {
		return a
	}
	if j == len(a)-i-1 {
		return bubbleSort(a, i+1, 0)
	}
	if a[j] > a[j+1] {
		a[j], a[j+1] = a[j+1], a[j]
	}
	return bubbleSort(a, i, j+1)
}
