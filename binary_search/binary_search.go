package binarysearch

func binarySearch(a []int, k int) bool {
	l := 0
	h := len(a) - 1
	for l <= h {
		m := (l + h) / 2
		if a[m] == k {
			return true
		}
		if k > a[m] {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return false
}


func bs(a []int, l, h, k int) bool {
	m := (l + h) / 2
	if l > h {
		return false
	}
	if a[m] == k {
		return true
	}
	if k > a[m] {
		return bs(a, m+1, h, k)
	}
	return bs(a, l, m-1, k)
}
