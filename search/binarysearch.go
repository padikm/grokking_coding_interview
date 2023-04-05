package search

func BinarySearch(a []int, i int) bool {
	low := 0
	high := len(a)
	if low > high {
		return false
	}
	if len(a) == 1 && a[0] != i {
		return false
	}
	mid := low + (high-low)/2
	if i < a[mid] {
		return BinarySearch(a[:mid], i)
	}
	if a[mid] == i {
		return true
	}
	return BinarySearch(a[mid:], i)

}

func BinarySearchIter(a []int, i int) bool {

	l := 0
	h := len(a)
	for l < h {
		mid := l + (h-l)/2
		if a[mid] == i {
			return true
		}

		if i < a[mid] {
			l = 0
			h = mid - 1
		} else {
			l = mid + 1
			h = len(a)
		}
	}
	return false
}
