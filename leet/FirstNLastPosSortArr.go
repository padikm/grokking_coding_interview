package leet

func SearchRange(a []int, key int) []int {
	found := binSearch(a, 0, len(a), key)
	if found == -1 {
		return []int{-1, -1}
	}
	var res []int
	left := binSearchFound(a, 0, found, key, false)
	if left == -1 {
		res = append(res, found)
	} else {
		res = append(res, left)
	}
	res = append(res, binSearchFound(a, found, len(a), key, true))
	return res
}

func binSearch(a []int, low, high, key int) int {
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

func binSearchFound(a []int, low, high, key int, isRight bool) int {
	var found = -1
	for low < high {
		mid := (low + high) / 2
		if a[mid] == key && !isRight {
			high = mid
			found = mid
		} else if a[mid] == key && isRight {
			low = mid + 1
			found = mid
		} else if a[mid] > key {
			high = mid
		} else if a[mid] < key {
			low = mid + 1
		}
	}
	return found
}
