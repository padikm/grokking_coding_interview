package sort

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i - 1
		key := a[i]
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}
