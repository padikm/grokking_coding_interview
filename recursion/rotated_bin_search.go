package recursion

func RotatedBinSearch(a []int, k int) int {
	return rotatedBinSearch(a, 0, len(a)-1, k)
}

func rotatedBinSearch(a []int, l, h, k int) int {
	if l > h {
		return -1
	}
	mid := (l + h) / 2
	if a[mid] == k {
		return mid
	}

	if a[l] > a[mid] {
		if k < a[l] && k > a[mid] {
			return rotatedBinSearch(a, mid+1, h, k)
		} else {
			return rotatedBinSearch(a, l, mid-1, k)
		}
	} else {
		if k >= a[l] && k <= a[mid] {
			return rotatedBinSearch(a, l, mid-1, k)
		} else {
			return rotatedBinSearch(a, mid+1, h, k)
		}
	}
	//if (a[l] > a[mid]) && (a[l] >= k && a[mid] > k) {
	//	return rotatedBinSearch(a, l, mid-1, k)
	//} else if a[l] > k && a[mid] < k {
	//	return rotatedBinSearch(a, mid+1, h, k)
	//} else if a[l] <= k && a[mid] > k {
	//	return rotatedBinSearch(a, l, mid-1, k)
	//} else {
	//	return rotatedBinSearch(a, mid+1, h, k)
	//}
}
