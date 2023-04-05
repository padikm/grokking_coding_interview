package leet

func SingleNonDuplicate(a []int) int {
	low := 0
	high := len(a) - 1
	if low < len(a)-1 && a[low] != a[low+1] {
		return a[low]
	} else if high > 0 && a[high] != a[high-1] {
		return a[high]
	}
	return binFind(a, low, high)

}

func binFind(a []int, low int, high int) int {
	for low < high {
		mid := (low + high) / 2
		if mid%2 != 0 {
			mid = mid - 1
		}
		if mid > 0 && mid < len(a) {
			if a[mid] != a[mid+1] && a[mid] != a[mid-1] {
				return a[mid]
			}
		}
		if a[mid] != a[mid+1] {
			high = mid
		} else {
			low = mid + 2
		}
	}

	return a[low]
}
