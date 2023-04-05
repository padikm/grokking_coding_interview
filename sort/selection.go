package sort

func SelectionSort(a []int) {
	//a = []int{5, 6, 7, 2, 1}
	for i := 0; i < len(a); i++ {
		var posSmall = i
		for j := i; j < len(a); j++ {
			if a[posSmall] > a[j] {
				posSmall = j
			}
		}
		tmp := a[i]
		a[i] = a[posSmall]
		a[posSmall] = tmp

	}
}
