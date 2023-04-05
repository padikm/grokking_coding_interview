package leet

func SearchInsert(a []int, key int) int {
	l := 0
	h := len(a)
	for l < h {
		mid := (l + h) / 2
		if a[mid] == key {
			return mid
		}
		if a[mid] > key {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return h
}
