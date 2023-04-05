package educative

func Search(a []int, key int) int {
	p := findPivot(a)
	if p < 0 {
		return binSearch(a, key, 0, len(a))
	}

	if a[len(a)-1] < key {
		return binSearch(a, key, 0, p)
	} else {
		return binSearch(a, key, p, len(a))
	}
	return -1
}

func binSearch(a []int, key, low, high int) int {
	//low := 0
	//high := len(a)
	for low < high {
		mid := (low + high) / 2
		if a[mid] == key {
			return mid
		}
		if key < a[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func findPivot(a []int) int {
	low := 0
	high := len(a) - 1

	for low < high {
		mid := (low + high) / 2

		if mid < high && a[mid] > a[mid+1] {
			return mid + 1
		}
		if mid > low && a[mid] < a[mid-1] {
			return mid
		}

		if a[low] >= a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1

}
