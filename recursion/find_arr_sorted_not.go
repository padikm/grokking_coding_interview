package recursion

func IsArrSorted(a []int) bool {
	return isArrSortedNeat(a, 0)
}

func isArrSorted(a []int, i, j int) bool {
	if j == len(a) {
		return true
	}
	if len(a) == 0 {
		return false
	}
	if len(a) == 1 {
		return true
	}

	if a[i] < a[j] {
		return isArrSorted(a, i+1, j+1)
	}
	return false
}

func isArrSortedNeat(a []int, i int) bool {
	if i == len(a)-1 {
		return true
	}
	return a[i] <= a[i+1] && isArrSortedNeat(a, i+1)
}
