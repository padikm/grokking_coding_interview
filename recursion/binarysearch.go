package recursion

func BinarySearch(a []int, l, h, k int) int {
	if l >= h {
		return -1
	}
	mid := (h + l) / 2

	if a[mid] == k {
		return mid
	}

	if a[mid] > k {
		return BinarySearch(a, l, mid-1, k)
	}
	return BinarySearch(a, mid+1, h, k)
}
