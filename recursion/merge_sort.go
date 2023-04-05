package recursion

func MergeSort(a []int, l, h int) {
	if l < h {
		mid := (h + l) / 2
		MergeSort(a, l, mid)
		MergeSort(a, mid+1, h)
		merge(a, l, mid, h)
	}
}

func merge(a []int, l int, mid int, h int) {
	k := 0
	i := l
	j := mid + 1
	res := make([]int, 0)
	for i <= mid && j <= h {
		if a[i] < a[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, a[j])
			j++
		}
		k++
	}

	for i <= mid {
		res = append(res, a[i])
		i++
	}

	for j <= h {
		res = append(res, a[j])
		j++
	}

	for i := l; i <= h; i++ {
		a[i] = res[i-l]
	}
}
