package leet

func Search(a []int, target int) bool {
	if len(a) == 1 {
		return a[0] == target
	}
	p := findPivotRotate(a)
	if p == -1 {
		if binSearchRotate(a, target, 0, len(a)) {
			return true
		} else {
			return false
		}
	}
	if a[len(a)-1] < target {
		return binSearchRotate(a, target, 0, p)
	} else {
		return binSearchRotate(a, target, p, len(a))
	}
	return false
}

func binSearchRotate(a []int, key, low, high int) bool {
	//low := 0
	//high := len(a)
	for low < high {
		mid := (low + high) / 2
		if a[mid] == key {
			return true
		}
		if key < a[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}

func findPivotRotate(a []int) int {
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

		if a[low] > a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1

}
