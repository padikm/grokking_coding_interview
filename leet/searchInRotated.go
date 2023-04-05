package leet

func SearchInRotated(a []int, key int) int {
	pivot := findPivot(a, 0, len(a)-1)
	low := 0
	high := 0
	if pivot == -1 {
		return binarySearch(a, low, len(a), key)
	}
	if key == a[pivot] {
		return pivot
	}
	if key >= a[0] {
		low = 0
		high = pivot
	} else {
		low = pivot + 1
		high = len(a)
	}
	return binarySearch(a, low, high, key)
}

func binarySearch(a []int, low, high, key int) int {
	for low < high {
		mid := (low + high) / 2
		if a[mid] == key {
			return mid
		}

		if a[mid] > key {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func findPivot(a []int, l, h int) int {

	if l > h {
		return -1
	}

	if h == l {
		return l
	}
	mid := (l + h) / 2

	if l < mid && a[mid-1] > a[mid] {
		return mid - 1
	}
	if h > mid && a[mid+1] < a[mid] {
		return mid + 1
	}
	if a[l] >= a[mid] {
		return findPivot(a, l, mid-1)
	}
	return findPivot(a, mid+1, h)
}
